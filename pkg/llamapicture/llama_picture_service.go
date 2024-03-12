package llamapicture

import (
	"context"
	restClient "github.com/liblaber/llama-store-sdk-go/internal/clients/rest"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/internal/configmanager"
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

func (api *LlamaPictureService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Get a llama's picture by the llama ID. Pictures are in PNG format.
func (api *LlamaPictureService) GetLlamaPictureByLlamaId(ctx context.Context, llamaId int64) (*shared.LlamaStoreResponse[[]byte], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[[]byte](config)

	request := httptransport.NewRequest(ctx, "GET", "/llama/{llama_id}/picture", config)

	request.SetPathParam("llama_id", llamaId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[[]byte](err)
	}

	return shared.NewLlamaStoreResponse[[]byte](resp), nil
}

// Create a picture for a llama. The picture is sent as a PNG as binary data in the body of the request.
func (api *LlamaPictureService) CreateLlamaPicture(ctx context.Context, llamaId int64, body []byte) (*shared.LlamaStoreResponse[LlamaId], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[LlamaId](config)

	request := httptransport.NewRequest(ctx, "POST", "/llama/{llama_id}/picture", config)

	request.Body = body

	request.SetPathParam("llama_id", llamaId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[LlamaId](err)
	}

	return shared.NewLlamaStoreResponse[LlamaId](resp), nil
}

// Update a picture for a llama. The picture is sent as a PNG as binary data in the body of the request.
//
// If the llama does not have a picture, one will be created. If the llama already has a picture,
// it will be overwritten.
// If the llama does not exist, a 404 will be returned.
func (api *LlamaPictureService) UpdateLlamaPicture(ctx context.Context, llamaId int64, body []byte) (*shared.LlamaStoreResponse[LlamaId], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[LlamaId](config)

	request := httptransport.NewRequest(ctx, "PUT", "/llama/{llama_id}/picture", config)

	request.Body = body

	request.SetPathParam("llama_id", llamaId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[LlamaId](err)
	}

	return shared.NewLlamaStoreResponse[LlamaId](resp), nil
}

// Delete a llama's picture by ID.
func (api *LlamaPictureService) DeleteLlamaPicture(ctx context.Context, llamaId int64) (*shared.LlamaStoreResponse[any], *shared.LlamaStoreError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/llama/{llama_id}/picture", config)

	request.SetPathParam("llama_id", llamaId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewLlamaStoreError[any](err)
	}

	return shared.NewLlamaStoreResponse[any](resp), nil
}
