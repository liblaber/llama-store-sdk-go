package handlers

import (
	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
)

type Handler interface {
	Handle(req httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse)
	SetNext(handler Handler)
}

type HandlerChain struct {
	head Handler
	tail Handler
}

func BuildHandlerChain() *HandlerChain {
	return &HandlerChain{}
}

func (chain *HandlerChain) AddHandler(handler Handler) *HandlerChain {
	if chain.head == nil {
		chain.head = handler
		chain.tail = handler
		return chain
	}

	chain.tail.SetNext(handler)
	chain.tail = handler

	return chain
}

func (chain *HandlerChain) CallApi(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	return chain.head.Handle(request)
}
