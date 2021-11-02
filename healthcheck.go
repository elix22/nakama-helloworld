package main

//go mod init elix22.com/hello-world
//go get github.com/heroiclabs/nakama-common/runtime
//go mod vendor

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HealthCheckResponse struct {
	Success bool `json : "success"`
}

func RpcHealthCheck(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("Healthcheck RPC called")

	response := &HealthCheckResponse{Success: true}

	out, err := json.Marshal(response)

	if err != nil {
		logger.Error("Error marshaling response type to JSON : %v", err)
		return "", runtime.NewError("Can't marshal type", 13)
	}

	return string(out), nil
}
