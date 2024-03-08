package user

import (
	"context"
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
	"github.com/liblaber/llama-store-sdk-go/pkg/shared"
)

type UserService struct {
	manager *configmanager.ConfigManager
}

func NewUserService(manager *configmanager.ConfigManager) *UserService {
	return &UserService{
		manager: manager,
	}
}

func (api *UserService) getConfig() *llamastoreconfig.Config {
	return api.manager.GetUser()
}

func (api *UserService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *UserService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Get a user by email.
//
// This endpoint will return a 404 if the user does not exist. Otherwise, it will return a 200.
func (api *UserService) GetUserByEmail(ctx context.Context, email string) (*shared.LlamaStoreResponse[User], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[User](config)

	request := httptransport.NewRequest(ctx, "GET", "/user/{email}", config)

	request.SetPathParam("email", email)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[User](err)
	}

	return shared.NewLlamaStoreResponse[User](resp), nil
}

// Register a new user.
//
// This endpoint will return a 400 if the user already exists. Otherwise, it will return a 201.
func (api *UserService) RegisterUser(ctx context.Context, userRegistration UserRegistration) (*shared.LlamaStoreResponse[User], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[User](config)

	request := httptransport.NewRequest(ctx, "POST", "/user", config)

	request.Body = userRegistration

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[User](err)
	}

	return shared.NewLlamaStoreResponse[User](resp), nil
}
