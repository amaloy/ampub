# ampub
AmPub is a very simple framework to help abstract the specifics of publishing to some message broker. It provides a simple HTTP API server and you provide the publishing implementation. Your client application can then communicate with the HTTP API without knowing the details about the implementation that was provided.

The envisioned use case is for an [**am**bassador/adaptor](https://kubernetes.io/blog/2015/06/the-distributed-system-toolkit-patterns) container in a composite-container deployment.

# Example
```go
package main

import (
	"context"
	"log"

	"github.com/amaloy/ampub"
)

func main() {
	// Create AmPub
	ampub := new(ampub.AmPub)

	// Create your publisher
	publisher := new(examplePublisher)

	// Run AmPub with your publisher
	ampub.Run(publisher)
}

// examplePublisher - Implementation of ampub.Publisher that publishes to log
type examplePublisher struct {
}

func (p *examplePublisher) Publish(ctx context.Context, topic string, key string, data []byte) error {
	// It's up to your implementation what to do with the values
	log.Printf("topic=%s key=%s data=%s", topic, key, string(data))
	return nil
}
```

# HTTP API
`POST /apiv1/topics/{T}` - Post the bytes in the request body to topic T.

`POST /apiv1/topics/{T}/key/{K}`  - Post the bytes in the request body to topic T given key K. A key is a common attribute in message systems and so it provided as an option.

# Environment Variables
* `AMPUB_ADDR` - The address to listen on, e.g. `0.0.0.0:4567`, default is `:8000`
* `AMPUB_LOGONLY` - If `true`, the provided Publisher will be ignored and publishing will go to `log`

# See Also
[ampub-gcppubsub](https://github.com/amaloy/ampub-gcppubsub) - An implementation for Google Cloud Pub/Sub
