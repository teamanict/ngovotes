package server

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/teamanict/ngovotes/config"
	"github.com/teamanict/ngovotes/utils"
)


type WebServer struct {
	App    *fiber.App
	Store  *session.Store
	Config *config.Config
}

func SetupWebServer(config *config.Config) (*WebServer, error) {
	// Views engine
	engine := html.New(config.Views.Folder, config.Views.Extension)

	// Pass locals to views

	// Set up webserver
	ws := &WebServer{
		App: fiber.New(fiber.Config{
			ServerHeader: config.Webserver.Header,
			AppName:      config.Webserver.AppName,
			Prefork:      config.Webserver.Prefork,
			Views: engine,
			PassLocalsToViews: true,
		}),
		Store: session.New(session.Config{
			Expiration: time.Duration(config.Session.ExpHrs) * time.Hour,
		}),
		Config: config,
	}

	
	// Add Extra Middlewares
	ws.App.Use(logger.New(logger.Config{
		Next:       utils.IsEnabled(config.Logger.Enabled),
		TimeFormat: config.Logger.Timeformat,
		TimeZone:   config.Logger.Timezone,
		Format:     config.Logger.Format,
	}))

	ws.App.Use(limiter.New(limiter.Config{
		Next:       utils.IsEnabled(config.Limiter.Enabled),
		Max:        config.Limiter.Max,
		Expiration: time.Duration(config.Session.ExpHrs) * time.Hour,
	}))

	ws.App.Use(compress.New(compress.Config{
		Next:  utils.IsEnabled(config.Compress.Enabled),
		Level: config.Compress.Level,
	}))

	return ws, nil
}

func (ws *WebServer) ListenWebServer() error {
	port := os.Getenv("PORT")
  	if os.Getenv("PORT") == "" {
		port = ws.Config.Webserver.Port
	}
	err := ws.App.Listen(":" + port)
	if err != nil {
		return err
	}

	return nil
}