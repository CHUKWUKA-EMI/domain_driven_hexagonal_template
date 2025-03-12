package config

import (
	"backend_api_template/internal/infrastructure/auth"
	"backend_api_template/internal/infrastructure/db"
	httputils "backend_api_template/internal/infrastructure/http"
	"backend_api_template/internal/infrastructure/logger"
	secretmanager "backend_api_template/internal/infrastructure/secret_manager"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bool64/cache"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ENV_VARS_SECRET_ID        = ""
	SERVICE_ACCOUNT_SECRET_ID = ""
)

// AppConfig holds the configuration for the application
type AppConfig struct {
	Name                string
	Env                 string
	Port                string
	AllowedOrigins      []string
	ProjectID           string
	Context             context.Context
	DBClient            *mongo.Client
	StorageBucket       string
	MessagingAPIBaseURL string
	ServiceAuthToken    string
	UserCache           *cache.Failover
	ready               bool
	HTTPClient          *httputils.HTTPClient `json:"-"`
	Secretmanager       *secretmanager.SecretManager

	Logger *logrus.Entry `json:"-"`

	Validator *validator.Validate `json:"-"`

	Auth *auth.Auth `json:"-"`
}

func (app *AppConfig) setup() {

	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	app.Name = "backend_api_template"

	app.Secretmanager = secretmanager.NewSecretManager(
		os.Getenv("SECRETS_API_BASE_URL"),
		os.Getenv("SECRETS_API_KEY"),
		app.HTTPClient,
	)

	// Load environment variables from secret/config manager
	log.Println("Loading config...")
	loadEnvVariables(app, ENV_VARS_SECRET_ID)
	log.Println("Config loaded and set as environment variables")

	app.Env = os.Getenv("ENV")
	app.Port = os.Getenv("PORT")
	if app.Port == "" {
		app.Port = "8080"
	}
	app.AllowedOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	app.ProjectID = os.Getenv("PROJECT_ID")

	app.Logger = logger.New().WithFields(logrus.Fields{
		"app": app.Name,
		"env": app.Env,
	})

	app.HTTPClient = httputils.NewHTTPClient()

	app.Validator = validator.New()

	app.Context = context.Background()

	app.MessagingAPIBaseURL = os.Getenv("MESSAGING_API_BASE_URL")

	app.StorageBucket = os.Getenv("STORAGE_BUCKET")

	// Initialize auth service
	authServiceCredentials := getAuthServiceCredentials(app, SERVICE_ACCOUNT_SECRET_ID)
	auth := auth.NewAuth(app.Context, app.Logger, app.HTTPClient, app.ProjectID, authServiceCredentials)
	app.Auth = auth

	// Get service auth token for authenticating with other services
	authToken, err := app.Auth.GetAuthToken(os.Getenv("AUTH_EMAIL"), os.Getenv("AUTH_PASSWORD"), os.Getenv("AUTH_API_KEY"))
	if err != nil {
		log.Fatalf("Error getting auth token: %s", err.Error())
	}
	app.ServiceAuthToken = authToken

	// Connect to database
	dbClient, err := db.ConnectDB(os.Getenv("DB_URL"), app.Logger)
	if err != nil {
		app.Logger.Fatalf("Error connecting to database: %s", err.Error())
	}
	app.DBClient = dbClient

	// Initialize user cache
	app.UserCache = cache.NewFailover(func(cfg *cache.FailoverConfig) {
		cfg.BackendConfig.TimeToLive = 5 * time.Minute
		cfg.MaxStaleness = 5 * time.Second
	})

	app.ready = true

	app.Logger.Info("App config setup complete")
}

// InitializeApp initializes the application and returns the its configuration
func InitializeApp() *AppConfig {
	app := &AppConfig{}
	app.setup()
	return app
}

// Ready returns the status of the application
func (app *AppConfig) Ready() bool {
	return app.ready
}

func loadEnvVariables(app *AppConfig, configKey string) {
	envVars, err := app.Secretmanager.GetSecret(configKey, nil)
	if err != nil {
		log.Fatalf("Error fetching app config: %s", err.Error())
	}

	envVarsByte := []byte(envVars)

	env := make(map[string]string)
	json.Unmarshal(envVarsByte, &env)
	// set environment variables
	for key, value := range env {
		os.Setenv(key, value)
	}
}

func getAuthServiceCredentials(app *AppConfig, secretID string) []byte {
	serviceAccount, err := app.Secretmanager.GetSecret(secretID, nil)
	if err != nil {
		app.Logger.Fatalf("Error getting firebase service account key: %s", err.Error())
	}
	return []byte(serviceAccount)
}
