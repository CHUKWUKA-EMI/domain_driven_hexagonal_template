package route

import (
	"backend_api_template/internal/adapter/http/handler"
	"backend_api_template/internal/adapter/persistence"
	"backend_api_template/internal/application"
	"backend_api_template/internal/infrastructure/config"
	"backend_api_template/internal/infrastructure/constants"
	"net/http"
)

func RegisterRoutes(app *config.AppConfig, mux *http.ServeMux) {

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	})

	userRepository := persistence.NewMongoDBUserRepository(app.DBClient.Database(constants.DatabaseName))
	userService := application.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, app)

	mux.HandleFunc("GET /users", userHandler.FindAll)
	mux.HandleFunc("GET /users/{id}", userHandler.FindByID)
	mux.HandleFunc("POST /users", userHandler.Save)
	mux.HandleFunc("PUT /users/{id}", userHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", userHandler.Delete)
}
