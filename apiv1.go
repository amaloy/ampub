package ampub

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func (ampub *AmPub) api1Publish(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
		return
	}

	vars := mux.Vars(r)
	topic := vars["t"]
	key := vars["k"]

	err = ampub.publisher.Publish(r.Context(), topic, key, body)
	if err != nil {
		writeError(w, err)
	} else {
		w.WriteHeader(200)
	}
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
