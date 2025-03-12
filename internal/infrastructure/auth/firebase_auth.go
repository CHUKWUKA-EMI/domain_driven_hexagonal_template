package auth

import (
	httputils "backend_api_template/internal/infrastructure/http"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type Auth struct {
	context     context.Context
	projectID   string
	credentials []byte
	authClient  *auth.Client
	logger      *logrus.Entry
	httpClient  *httputils.HTTPClient
}

func NewAuth(ctx context.Context, logger *logrus.Entry, httpClient *httputils.HTTPClient, projectID string, credentials []byte) *Auth {
	app, err := createApp(ctx, projectID, credentials)
	if err != nil {
		logger.Fatalf("Error initializing firebase app: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		logger.Fatalf("Error initializing firebase auth client: %v", err)
	}

	return &Auth{
		context:     ctx,
		projectID:   projectID,
		credentials: credentials,
		authClient:  authClient,
		logger:      logger,
		httpClient:  httpClient,
	}
}

// GetPrincipalByEmail retrieves an auth user by email
func (a *Auth) GetPrincipalByEmail(email string) (*auth.UserRecord, error) {
	user, err := a.authClient.GetUserByEmail(a.context, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// SetCustomAuthClaims sets auth claims/roles on an auth user account
func (a *Auth) SetCustomAuthClaims(uid string, userClaims map[string]interface{}) error {
	err := a.authClient.SetCustomUserClaims(a.context, uid, userClaims)

	if err != nil {
		return err
	}

	return nil
}

// VerifyAuthToken verifies the authenticity of a user's token
func (a *Auth) VerifyAuthToken(idToken string, uid string) (string, error) {

	authToken, err := a.authClient.VerifyIDToken(a.context, idToken)

	if err != nil {
		return "", err
	}

	if authToken.UID != uid {
		return "", errors.New("Unathorized request by user " + uid)
	}

	return authToken.UID, nil
}

func (a *Auth) GetPrincipal(uid string) (*auth.UserRecord, error) {
	user, err := a.authClient.GetUser(a.context, uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAuthToken retrieves an authentication token from Firebase
func (a *Auth) GetAuthToken(email, password, APIKey string) (string, error) {
	payload := map[string]interface{}{
		"email":             email,
		"password":          password,
		"returnSecureToken": true,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	response, err := a.httpClient.SendRequest(
		http.MethodPost,
		fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", APIKey),
		bytes.NewReader(payloadBytes),
		map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return "", err
	}

	token, ok := responseBody["idToken"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected token response")
	}
	return token, nil
}

func createApp(ctx context.Context, projectID string, credentials []byte) (*firebase.App, error) {
	options := option.WithCredentialsJSON(credentials)
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID}, options)
	if err != nil {
		return nil, err
	}

	return app, nil
}
