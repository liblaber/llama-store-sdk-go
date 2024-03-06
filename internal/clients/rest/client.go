package rest

import (
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/handlers"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/hooks"
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
)

type RestClient struct {
	handlers *handlers.HandlerChain
}

func NewRestClient(config llamastoreconfig.Config) *RestClient {
	defaultHeadersHandler := handlers.NewDefaultHeadersHandler()
	retryHandler := handlers.NewRetryHandler()
	bearerTokenHandler := handlers.NewAccessTokenHandler()
	hookHandler := handlers.NewHookHandler(hooks.NewCustomHook())
	requestValidationHandler := handlers.NewRequestValidationHandler()
	terminatingHandler := handlers.NewTerminatingHandler()

	handlers := handlers.BuildHandlerChain().
		AddHandler(defaultHeadersHandler).
		AddHandler(retryHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(hookHandler).
		AddHandler(requestValidationHandler).
		AddHandler(terminatingHandler)

	return &RestClient{
		handlers: handlers,
	}
}

func (client *RestClient) Call(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	return client.handlers.CallApi(request)
}
