package llama

import (
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
	"github.com/liblaber/llama-store-sdk-go/internal/unmarshal"
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

// Get all the llamas.
func (api *LlamaService) GetLlamas() (*shared.LlamaStoreResponse[[]Llama], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/llama", config)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToArray[[]Llama](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[[]Llama]{
		Data: data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Create a new llama. Llama names must be unique.
func (api *LlamaService) CreateLlama(llamaCreate LlamaCreate) (*shared.LlamaStoreResponse[Llama], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("POST", "/llama", config)

	request.Body = llamaCreate

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Llama](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[Llama]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Get a llama by ID.
func (api *LlamaService) GetLlamaById(llamaId int64) (*shared.LlamaStoreResponse[Llama], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/llama/{llama_id}", config)

	request.SetPathParam("llama_id", llamaId)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Llama](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[Llama]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Update a llama. If the llama does not exist, create it.
//
// When updating a llama, the llama name must be unique. If the llama name is not unique, a 409 will be returned.
func (api *LlamaService) UpdateLlama(llamaId int64, llamaCreate LlamaCreate) (*shared.LlamaStoreResponse[Llama], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("PUT", "/llama/{llama_id}", config)

	request.Body = llamaCreate

	request.SetPathParam("llama_id", llamaId)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[Llama](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[Llama]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Delete a llama. If the llama does not exist, this will return a 404.
func (api *LlamaService) DeleteLlama(llamaId int64) error {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("DELETE", "/llama/{llama_id}", config)

	request.SetPathParam("llama_id", llamaId)

	_, err := client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil
}
