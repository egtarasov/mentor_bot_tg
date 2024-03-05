package daemons

import (
	"context"
	"log"
	"time"
)

type daemon interface {
	start(ctx context.Context) error
}

func startDaemon(ctx context.Context, duration time.Duration, daemon daemon) {
	log.Println("Start the daemon!")
	err := daemon.start(ctx)
	if err != nil {
		// TODO Log error
	}
	ticker := time.NewTicker(duration)
	for {
		select {
		// TODO graceful shutdown.
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := daemon.start(ctx)
			if err != nil {
				// TODO Log error
			}
		}
	}
}
