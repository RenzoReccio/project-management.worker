package azureapi_utils

import (
	"encoding/json"
	"net/http"
)

func GetJson(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
