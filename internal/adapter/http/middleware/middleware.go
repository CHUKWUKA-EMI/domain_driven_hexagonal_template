package middleware

import (
	"backend_api_template/internal/adapter/http/handler"
	"backend_api_template/internal/infrastructure/config"
	"context"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"firebase.google.com/go/v4/auth"
)

type ContextKey string

const (
	// UserContextKey is the key for the user in the context
	UserContextKey ContextKey = "user"
)

type Middleware struct {
	appConfig *config.AppConfig
	handler   http.Handler
}

func New(appConfig *config.AppConfig, handler http.Handler) *Middleware {
	return &Middleware{appConfig: appConfig, handler: handler}
}

func (m *Middleware) LogIP() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.Header.Get("X-Forwarded-For")
		if clientIP == "" {
			clientIP = r.Header.Get("X-Real-IP")
		}
		if clientIP == "" {
			clientIP = r.RemoteAddr
		}
		fmt.Println("Client IP: ", clientIP)

		m.handler.ServeHTTP(w, r)
	})
}

func ApplyMiddlewares(appConfig *config.AppConfig, h http.Handler) http.Handler {
	return ChainMiddleware(
		SetCors(appConfig),
		Authenticate(appConfig),
	)(h)
}

func ChainMiddleware(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}
		return h
	}
}

// Authenticate is a middleware for authenticating requests
func Authenticate(appConfig *config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api" {
				next.ServeHTTP(w, r)
				return
			}
			authHeader := r.Header.Get("authorization")

			if authHeader == "" {
				handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Missing or Invalid Authorization Header"})
				return
			}

			authPayload := strings.Split(authHeader, " ")
			if len(authPayload) != 2 {
				handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid authentication scheme"})
				return
			}

			authToken := authPayload[1]

			uid := r.Header.Get("UID")
			if uid == "" {
				handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Missing UID request header. Please include it."})
				return
			}

			uid, err := appConfig.Auth.VerifyAuthToken(authToken, uid)

			if err != nil {
				handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
				return
			}
			user, err := appConfig.Auth.GetPrincipal(uid)
			if err != nil {
				handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, user)
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}

func (m *Middleware) Authenticate2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" {
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("authorization")

		if authHeader == "" {
			handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Missing or Invalid Authorization Header"})
			return
		}

		authPayload := strings.Split(authHeader, " ")
		if len(authPayload) != 2 {
			handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid authentication scheme"})
			return
		}

		authToken := authPayload[1]

		uid := r.Header.Get("UID")
		if uid == "" {
			handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Missing UID request header. Please include it."})
			return
		}

		uid, err := m.appConfig.Auth.VerifyAuthToken(authToken, uid)

		if err != nil {
			handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}
		user, err := m.appConfig.Auth.GetPrincipal(uid)
		if err != nil {
			handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, user)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

// HasRoles is a middleware for checking if a user has the required roles
func HasRoles(roles []string, appConfig *config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentUser := r.Context().Value(UserContextKey)
			authUserRecord, ok := currentUser.(*auth.UserRecord)

			if !ok {
				handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Error retrieving auth user record"})
				return
			}

			userRole := authUserRecord.CustomClaims["role"].(string)
			if !slices.Contains(roles, userRole) {
				handler.HandleResponse(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized access"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func SetCors(appConfig *config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := appConfig.AllowedOrigins
			origin := r.Header.Get("Origin")

			if slices.Contains(allowedOrigins, origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// Handle preflight request
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
