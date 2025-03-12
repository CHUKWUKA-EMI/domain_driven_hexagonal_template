package main

import (
	"backend_api_template/internal/adapter/http/middleware"
	"backend_api_template/internal/adapter/http/route"
	"backend_api_template/internal/infrastructure/config"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	appConfig := config.InitializeApp()

	mux := http.NewServeMux()

	route.RegisterRoutes(appConfig, mux)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", appConfig.Port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	mdw := middleware.New(appConfig)
	mdw.Use(mdw.LogIP)
	mdw.Use(mdw.SetCors)

	// middleware.Use(middleware.LogIP)
	log.Println("Server running on port 5001")
	log.Fatal(http.Serve(listener, mdw.Apply(mux)))
}
