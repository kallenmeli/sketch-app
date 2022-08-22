package tests

import "github.com/goccy/go-json"

func ToJSON[T any](input T) []byte {
	data, _ := json.Marshal(input)
	return data
}
