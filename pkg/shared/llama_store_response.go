package shared

type LlamaStoreResponse[T any] struct {
	Data     T
	Metadata LlamaStoreResponseMetadata
}

type LlamaStoreResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}
