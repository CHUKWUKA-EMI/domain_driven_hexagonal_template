package secretmanager

import (
	httputils "backend_api_template/internal/infrastructure/http"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// SecretManager struct
type SecretManager struct {
	apiBaseURL string
	apiKey     string
	httpClient *httputils.HTTPClient
}

// NewSecretManager creates a new secret manager instance
func NewSecretManager(apiBaseURL string, apiKey string, httpClient *httputils.HTTPClient) *SecretManager {
	return &SecretManager{
		apiBaseURL: apiBaseURL,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// GetSecret retrieves a secret from the secret manager
func (s SecretManager) GetSecret(secretID string, version *int) (string, error) {
	httpClient := s.httpClient
	if httpClient == nil {
		httpClient = httputils.NewHTTPClient()
	}

	secretEndpoint := s.apiBaseURL + "/secret/" + secretID
	if version != nil {
		secretEndpoint = fmt.Sprintf("%s?version=%d", secretEndpoint, *version)
	}

	response, err := httpClient.SendRequest(http.MethodGet, secretEndpoint, nil, map[string]string{
		"Authorization": "Bearer " + s.apiKey,
	})
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", errors.New("failed to retrieve secret")
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responseBodyPayload map[string]string

	err = json.Unmarshal(responseBody, &responseBodyPayload)
	if err != nil {
		return "", err
	}

	return responseBodyPayload["secret"], nil
}
