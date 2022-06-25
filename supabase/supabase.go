package supabase

import (
	"fmt"

	supabase "github.com/nedpals/supabase-go"
)

func sbase() {
  supabaseUrl := "https://pmfatdsurouterlgrsvl.supabase.co"
  supabaseKey := "<SUPABASE-KEY>"
  SupabaseClient := supabase.CreateClient(supabaseUrl, supabaseKey)

  // Auth
  _, err := SupabaseClient.Auth.SignIn(nil, supabase.UserCredentials{
    Email: "teamanict@gmail.com",
    Password: "frkuipers2022",
  })
  if err != nil {
    panic(err)
  }

  // DB
  var results map[string]interface{}
  err = SupabaseClient.DB.From("something").Select("*").Single().Execute(&results)
  if err != nil {
    panic(err)
  }

  fmt.Println(results)
}