package unmarshal

import (
	"encoding/json"

	"github.com/liblaber/llama-store-sdk-go/internal/clients/rest/httptransport"
)

func ToArray[R ~[]T, T any](r *httptransport.Response) (R, error) {
	result := make(R, 0)
	err := json.Unmarshal(r.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
