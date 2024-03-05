package main

import (
	"fmt"

	"github.com/liblaber/llama-store-sdk-go/pkg/llamastore"
	"github.com/liblaber/llama-store-sdk-go/pkg/llamastoreconfig"
)

func main() {
	config := llamastoreconfig.NewConfig()

	llamaStore := llamastore.NewLlamaStore(config)

	res, err := llamaStore.Llama.GetLlamas()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}
