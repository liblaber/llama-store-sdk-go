package rest

import (
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/handlers"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/hooks"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
)

type RestClient[T any] struct {
	handlers *handlers.HandlerChain[T]
}

func NewRestClient[T any](config llamastoreconfig.Config) *RestClient[T] {
	defaultHeadersHandler := handlers.NewDefaultHeadersHandler[T]()
	retryHandler := handlers.NewRetryHandler[T]()
	bearerTokenHandler := handlers.NewAccessTokenHandler[T]()
	responseValidationHandler := handlers.NewResponseValidationHandler[T]()
	unmarshalHandler := handlers.NewUnmarshalHandler[T]()
	requestValidationHandler := handlers.NewRequestValidationHandler[T]()
	hookHandler := handlers.NewHookHandler[T](hooks.NewCustomHook())
	terminatingHandler := handlers.NewTerminatingHandler[T]()

	handlers := handlers.BuildHandlerChain[T]().
		AddHandler(defaultHeadersHandler).
		AddHandler(retryHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(responseValidationHandler).
		AddHandler(unmarshalHandler).
		AddHandler(requestValidationHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient[T]{
		handlers: handlers,
	}
}

func (client *RestClient[T]) Call(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	return client.handlers.CallApi(request)
}
