package ampub

import "context"

// Publisher - Interface for a type than can publish
type Publisher interface {
	// Publish - publish the provided data to a topic, with an optional (may be empty) key attribute
	// Any error returned will result in a 500 status coded returned from the AmPub API server
	Publish(ctx context.Context, topic string, key string, data []byte) error
}
