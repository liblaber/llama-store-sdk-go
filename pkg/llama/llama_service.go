package llama

import (
	"context"
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
	"github.com/liblaber/llama-store-sdk-go/pkg/shared"
)

type LlamaService struct {
	manager *configmanager.ConfigManager
}

func NewLlamaService(manager *configmanager.ConfigManager) *LlamaService {
	return &LlamaService{
		manager: manager,
	}
}

func (api *LlamaService) getConfig() *llamastoreconfig.Config {
	return api.manager.GetLlama()
}

func (api *LlamaService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *LlamaService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Get all the llamas.
func (api *LlamaService) GetLlamas(ctx context.Context) (*shared.LlamaStoreResponse[[]Llama], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[[]Llama](config)

	request := httptransport.NewRequest(ctx, "GET", "/llama", config)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[[]Llama](err)
	}

	return shared.NewLlamaStoreResponse[[]Llama](resp), nil
}

// Create a new llama. Llama names must be unique.
func (api *LlamaService) CreateLlama(ctx context.Context, llamaCreate LlamaCreate) (*shared.LlamaStoreResponse[Llama], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Llama](config)

	request := httptransport.NewRequest(ctx, "POST", "/llama", config)

	request.Body = llamaCreate

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[Llama](err)
	}

	return shared.NewLlamaStoreResponse[Llama](resp), nil
}

// Get a llama by ID.
func (api *LlamaService) GetLlamaById(ctx context.Context, llamaId int64) (*shared.LlamaStoreResponse[Llama], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Llama](config)

	request := httptransport.NewRequest(ctx, "GET", "/llama/{llama_id}", config)

	request.SetPathParam("llama_id", llamaId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[Llama](err)
	}

	return shared.NewLlamaStoreResponse[Llama](resp), nil
}

// Update a llama. If the llama does not exist, create it.
//
// When updating a llama, the llama name must be unique. If the llama name is not unique, a 409 will be returned.
func (api *LlamaService) UpdateLlama(ctx context.Context, llamaId int64, llamaCreate LlamaCreate) (*shared.LlamaStoreResponse[Llama], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Llama](config)

	request := httptransport.NewRequest(ctx, "PUT", "/llama/{llama_id}", config)

	request.Body = llamaCreate

	request.SetPathParam("llama_id", llamaId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[Llama](err)
	}

	return shared.NewLlamaStoreResponse[Llama](resp), nil
}

// Delete a llama. If the llama does not exist, this will return a 404.
func (api *LlamaService) DeleteLlama(ctx context.Context, llamaId int64) (*shared.LlamaStoreResponse[any], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/llama/{llama_id}", config)

	request.SetPathParam("llama_id", llamaId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[any](err)
	}

	return shared.NewLlamaStoreResponse[any](resp), nil
}
