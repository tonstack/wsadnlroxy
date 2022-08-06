package proxy

import (
	"context"
	"time"

	"github.com/xssnick/tonutils-go/liteclient"
)

func checkLiteServerConnection(cfg *liteServerConfig) error {
	client := liteclient.NewConnectionPool()

	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelCtx()

	err := client.AddConnection(ctx, cfg.hostport, cfg.pubkey)
	if err != nil {
		return err
	}

	return nil
}
