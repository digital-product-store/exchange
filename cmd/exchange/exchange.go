//go:generate oapi-codegen -package gen -generate types,server,spec -o ../../pkg/server/gen/specs.gen.go ../../specs/openapi3.yaml

package main

import (
	"os"
	"userservice/pkg/config"
	"userservice/pkg/server"
	"userservice/pkg/storage"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("error on during configuration", zap.Error(err))
		os.Exit(-1)
	}

	roDB := createROMemStorage()
	handler := server.NewHandler(logger, roDB)
	srvr := server.NewServer(&handler, conf)

	srvr.Listen()
}

func createROMemStorage() storage.ROMemStorage {
	rates := []storage.ExchangeRate{
		{
			From: "EUR",
			To:   "USD",
			Rate: 1.12,
		},
	}

	return storage.NewROMemStorage(rates)
}
