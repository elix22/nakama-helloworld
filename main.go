package main

//  My configurtion
//  sudo rm -rf /usr/local/go
//  sudo tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz
// .bashrc , export PATH="$PATH:/usr/local/go/bin"
// "name":"nakama-gce-1","version":"3.3.0+83fc6fbc","runtime":"go1.16.4"

// go mod init elix22.com/hello-world
// go get github.com/heroiclabs/nakama-common/runtime:v1.14.0
// go mod vendor
// env GO111MODULE=on
// env CGO_ENABLED=1
// go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so
// sudo cp backend.so  /var/lib/nakama-data/modules/

// data_dir: /var/lib/nakama-data/modules/
// ps -aux | grep cockroach
// ps -aux | grep nakama
// sudo kill -9 ...

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()

	err := initializer.RegisterRpc("healthcheck", RpcHealthCheck)

	if err != nil {
		return err
	}

	logger.Info("Eli Plugin loaded in '%d' msec.", time.Now().Sub(initStart).Milliseconds())
	return nil
}
