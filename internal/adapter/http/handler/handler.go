package handler

import (
	"backend_api_template/internal/application"
	"backend_api_template/internal/domain"
	"backend_api_template/internal/infrastructure/config"
	"encoding/json"
	"net/http"
	"strings"
)

// UserHandler ...
type UserHandler struct {
	service   domain.Service
	appConfig *config.AppConfig
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(service domain.Service, appConfig *config.AppConfig) *UserHandler {
	return &UserHandler{service: service, appConfig: appConfig}
}

// FindAll returns all users
func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.FindAllUsers()
	if err != nil {
		HandleResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	HandleResponse(w, http.StatusOK, map[string]interface{}{"users": users})
}

// FindByID returns a user by ID
func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" {
		HandleResponse(w, http.StatusBadRequest, map[string]string{"error": "id is required"})
		return
	}

	user, err := h.service.FindUserByID(id)
	if err != nil {
		HandleResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	HandleResponse(w, http.StatusOK, map[string]interface{}{"user": user})
}

// Save creates a new user
func (h *UserHandler) Save(w http.ResponseWriter, r *http.Request) {
	var userDto application.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		HandleResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := userDto.Validate(h.appConfig.Validator); err != nil {
		HandleResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	user := userDto.ToDomain()
	newUser, err := h.service.CreateUser(user.FirstName(), user.LastName(), user.Email(), user.Address().Street(), user.Address().City(), user.Address().State(), user.Address().ZipCode())
	if err != nil {
		HandleResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	HandleResponse(w, http.StatusCreated, map[string]interface{}{"user": newUser})
}

// Update updates a user
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" {
		HandleResponse(w, http.StatusBadRequest, map[string]string{"error": "id is required"})
		return
	}

	var userDto application.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		HandleResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := userDto.Validate(h.appConfig.Validator); err != nil {
		HandleResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	user := userDto.ToDomain()
	updatedUser, err := h.service.UpdateUser(id, user.FirstName(), user.LastName(), user.Email(), user.Address().Street(), user.Address().City(), user.Address().State(), user.Address().ZipCode())
	if err != nil {
		HandleResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	HandleResponse(w, http.StatusOK, map[string]interface{}{"user": updatedUser})
}

// Delete deletes a user
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" {
		HandleResponse(w, http.StatusBadRequest, map[string]string{"error": "id is required"})
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		HandleResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	HandleResponse(w, http.StatusNoContent, nil)
}

func HandleResponse(w http.ResponseWriter, statusCode int, responseBody interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response, _ := json.Marshal(responseBody)
	w.Write(response)
}
