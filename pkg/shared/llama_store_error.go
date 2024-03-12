package shared

import (
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
)

type LlamaStoreError struct {
	Err      error
	Body     []byte
	Metadata LlamaStoreErrorMetadata
}

type LlamaStoreErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewLlamaStoreError[T any](transportError *httptransport.ErrorResponse[T]) *LlamaStoreError {
	return &LlamaStoreError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: LlamaStoreErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *LlamaStoreError) Error() string {
	return e.Err.Error()
}
