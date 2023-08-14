package main

import (
	"github.com/timeforaninja/shortpaste/internal/webapp"
	"os"
)

func main() {
	var bind, storagePath, username, password string
	var link307Redirect, ok bool

	if bind, ok = os.LookupEnv("SP_BIND_ADDR"); !ok {
		bind = ":8080"
	}

	// Support for the common PORT environment variable
	if port, ok := os.LookupEnv("PORT"); ok {
		bind = ":" + port
	}

	if storagePath, ok = os.LookupEnv("SP_STORAGE_PATH"); !ok {
		storagePath = "~/.shortpaste"
	}

	_, link307Redirect = os.LookupEnv("SP_307_REDIRECT")

	if username, ok = os.LookupEnv("SP_USERNAME"); !ok {
		username = "admin"
	}

	if password, ok = os.LookupEnv("SP_PASSWORD"); !ok {
		password = "admin"
	}

	app := api.NewApp(bind, storagePath, username, password, link307Redirect)
	app.Run()
}
