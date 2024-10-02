package main

import (
	"fmt"
	"net/http"
	"sip/config"
	"sip/internals/handlers"
)

const port = ":8080"

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	handlers.CreateTemple(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/help", handlers.Help)
	fmt.Println("http://localhost" + appConfig.Port + " - server started")
	http.ListenAndServe(appConfig.Port, nil)
}
