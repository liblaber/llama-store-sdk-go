package token

import (
	"context"
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
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

func (api *TokenService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Create an API token for a user. These tokens expire after 30 minutes.
//
// Once you have this token, you need to pass it to other endpoints in the Authorization header as a Bearer token.
func (api *TokenService) CreateApiToken(ctx context.Context, apiTokenRequest ApiTokenRequest) (*shared.LlamaStoreResponse[ApiToken], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ApiToken](config)

	request := httptransport.NewRequest(ctx, "POST", "/token", config)

	request.Body = apiTokenRequest

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[ApiToken](err)
	}

	return shared.NewLlamaStoreResponse[ApiToken](resp), nil
}
