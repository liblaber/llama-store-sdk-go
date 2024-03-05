package llamapicture

import (
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
	"github.com/liblaber/llama-store-sdk-go/internal/unmarshal"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
	"github.com/liblaber/llama-store-sdk-go/pkg/shared"
)

type LlamaPictureService struct {
	manager *configmanager.ConfigManager
}

func NewLlamaPictureService(manager *configmanager.ConfigManager) *LlamaPictureService {
	return &LlamaPictureService{
		manager: manager,
	}
}

func (api *LlamaPictureService) getConfig() *llamastoreconfig.Config {
	return api.manager.GetLlamaPicture()
}

func (api *LlamaPictureService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

// Get a llama's picture by the llama ID. Pictures are in PNG format.
func (api *LlamaPictureService) GetLlamaPictureByLlamaId(llamaId int64) error {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("GET", "/llama/{llama_id}/picture", config)

	request.SetPathParam("llama_id", llamaId)

	_, err := client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil
}

// Create a picture for a llama. The picture is sent as a PNG as binary data in the body of the request.
func (api *LlamaPictureService) CreateLlamaPicture(llamaId int64) (*shared.LlamaStoreResponse[LlamaId], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("POST", "/llama/{llama_id}/picture", config)

	request.SetPathParam("llama_id", llamaId)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[LlamaId](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[LlamaId]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Update a picture for a llama. The picture is sent as a PNG as binary data in the body of the request.
//
// If the llama does not have a picture, one will be created. If the llama already has a picture,
// it will be overwritten.
// If the llama does not exist, a 404 will be returned.
func (api *LlamaPictureService) UpdateLlamaPicture(llamaId int64) (*shared.LlamaStoreResponse[LlamaId], error) {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("PUT", "/llama/{llama_id}/picture", config)

	request.SetPathParam("llama_id", llamaId)

	httpResponse, err := client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	data, unmarshalError := unmarshal.ToObject[LlamaId](httpResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	response := shared.LlamaStoreResponse[LlamaId]{
		Data: *data,
		Metadata: shared.LlamaStoreResponseMetadata{
			Headers:    httpResponse.Headers,
			StatusCode: httpResponse.StatusCode,
		},
	}

	return &response, nil
}

// Delete a llama's picture by ID.
func (api *LlamaPictureService) DeleteLlamaPicture(llamaId int64) error {
	config := *api.getConfig()

	client := restClient.NewRestClient(config)

	request := httptransport.NewRequest("DELETE", "/llama/{llama_id}/picture", config)

	request.SetPathParam("llama_id", llamaId)

	_, err := client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil
}
