package logic

import (
	"api/pkg/log"
	"context"
	"encoding/json"
)

func ParseJSON[T any](input any, ctx context.Context) (result T, err error) {
	logger := log.New(ctx)
	jsonValue, err := json.Marshal(input)

	if err != nil {
		logger.Error(err, "Failed to marshal input")
		return
	}

	result, err = Unmarshal[T](jsonValue, ctx)
	return
}

func Unmarshal[T any](input []byte, ctx context.Context) (result T, err error) {
	logger := log.New(ctx)
	err = json.Unmarshal(input, &result)

	if err != nil {
		logger.Error(err, "Failed to unmarshal")
	}

	return
}
