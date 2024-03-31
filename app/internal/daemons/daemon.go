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
		log.Println(err)
	}
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			log.Println("Trigger the daemon")
			err := daemon.start(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
