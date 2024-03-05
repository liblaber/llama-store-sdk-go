package token

import (
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
	"github.com/liblaber/llama-store-sdk-go/internal/unmarshal"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
	"github.com/liblaber/llama-store-sdk-go/pkg/shared"
)

type TokenService struct {
	manager *configmanager.ConfigManager
}

func NewTokenService(manager *configmanager.ConfigManager) *TokenService {
	return &TokenService{
		manager: manager,
	}
}

func (api *TokenService) getConfig() *llamastoreconfig.Config {
	return api.manager.GetToken()
}

func (api *TokenService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

// Create an API token for a user. These tokens expire after 30 minutes.
//
// Once you have this token, you need to pass it to other endpoints in the Authorization header as a Bearer token.
func (api *TokenService) CreateApiToken(apiTokenRequest ApiTokenRequest) (*shared.LlamaStoreResponse[ApiToken], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("POST", "/token", config)

	request.Body = apiTokenRequest

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[ApiToken](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[ApiToken]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}
