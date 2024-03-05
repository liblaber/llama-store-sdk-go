package unmarshal

import (
	"encoding/json"

	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
)

func ToObject[T any](r *httptransport.Response) (*T, error) {
	result := new(T)
	err := json.Unmarshal(r.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
