package shared

import (
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
)

type LlamaStoreResponse[T any] struct {
	Data     T
	Metadata LlamaStoreResponseMetadata
}

type LlamaStoreResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewLlamaStoreResponse[T any](resp *httptransport.Response[T]) *LlamaStoreResponse[T] {
	return &LlamaStoreResponse[T]{
		Data: resp.Data,
		Metadata: LlamaStoreResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}
