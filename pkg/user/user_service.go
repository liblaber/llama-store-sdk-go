package user

import (
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
	"github.com/liblaber/llama-store-sdk-go/internal/unmarshal"
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

// Get a user by email.
//
// This endpoint will return a 404 if the user does not exist. Otherwise, it will return a 200.
func (api *UserService) GetUserByEmail(email string) (*shared.LlamaStoreResponse[User], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/user/{email}", config)

	request.SetPathParam("email", email)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[User](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[User]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Register a new user.
//
// This endpoint will return a 400 if the user already exists. Otherwise, it will return a 201.
func (api *UserService) RegisterUser(userRegistration UserRegistration) (*shared.LlamaStoreResponse[User], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("POST", "/user", config)

	request.Body = userRegistration

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[User](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[User]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}
