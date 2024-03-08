package main

import (
	"context"
	"fmt"

	"github.com/liblaber/llama-store-sdk-go/pkg/llamastore"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
)

func main() {
	config := llamastoreconfig.NewConfig()

	llamaStore := llamastore.NewLlamaStore(config)

	res, err := llamaStore.Llama.GetLlamas(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}
