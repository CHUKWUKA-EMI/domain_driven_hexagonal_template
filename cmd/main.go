package main

import (
	"backend_api_template/internal/adapter/http/middleware"
	"backend_api_template/internal/adapter/http/route"
	"backend_api_template/internal/infrastructure/config"
	"log"
	"net/http"
)

func main() {

	appConfig := config.InitializeApp()

	mux := http.NewServeMux()

	route.RegisterRoutes(appConfig, mux)

	// listener, err := net.Listen("tcp", ":5001")
	// if err != nil {
	// 	panic(err)
	// }
	// defer listener.Close()

	middleware := middleware.New(appConfig, mux)

	log.Println("Server running on port 5001")
	log.Fatal(http.ListenAndServe(":8080", middleware.LogIP()))
	// log.Fatal(http.Serve(listener, middleware.LogIP()))
}
