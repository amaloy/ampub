package ampub

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

const version = "0"
const defaultAddr = ":8000"
const envVarPrefix = "AMPUB_"

// AmPub - The point from which an AmPub service can be executed
type AmPub struct {
	initOnce  sync.Once
	publisher Publisher
}

// Run - Begin running an AmPub service with the given Publisher
func (ampub *AmPub) Run(publisher Publisher) {
	ampub.initOnce.Do(func() {
		log.Printf("ampub version %s", version)

		addr := defaultAddr
		if value, found := getEnvVar("ADDR"); found {
			addr = value
		}

		if value, found := getEnvVar("LOGONLY"); found {
			if logOnly, _ := strconv.ParseBool(value); logOnly {
				publisher = new(logPublisher)
			}
		}

		ampub.publisher = publisher

		router := mux.NewRouter()
		router.HandleFunc("/apiv1/topics/{t}", ampub.api1Publish).Methods("POST")
		router.HandleFunc("/apiv1/topics/{t}/key/{k}", ampub.api1Publish).Methods("POST")
		log.Printf("Listening on %s", addr)
		log.Fatal(http.ListenAndServe(addr, router))
	})
}

func getEnvVar(name string) (value string, found bool) {
	name = envVarPrefix + name
	value, found = os.LookupEnv(name)
	log.Printf("env: %s=%s", name, value)
	return
}
