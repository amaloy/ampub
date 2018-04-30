package ampub

import (
	"context"
	"log"
)

type logPublisher struct {
}

func (p *logPublisher) Publish(ctx context.Context, topic string, key string, data []byte) error {
	log.Printf("topic=%s key=%s data=%s", topic, key, string(data))
	return nil
}
