package ampub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestApi1PublishNoKey(t *testing.T) {
	wanted := recordingPublisher{topic: "atesttopic", key: "", data: "this is some data"}

	req, err := http.NewRequest("POST", "/apiv1/topics/t", strings.NewReader(wanted.data))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"t": wanted.topic})

	ampub := new(AmPub)
	ampub.publisher = new(recordingPublisher)
	handler := http.HandlerFunc(ampub.api1Publish)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assertStatusOK(t, rr)
	assertContextNotNil(t, ampub.publisher.(*recordingPublisher))
	assertEquals(t, &wanted, ampub.publisher.(*recordingPublisher))
}

func TestApi1PublishWithKey(t *testing.T) {
	wanted := recordingPublisher{topic: "atesttopic", key: "testingkey", data: "this is some data"}

	req, err := http.NewRequest("POST", "/apiv1/topics/t/key/k", strings.NewReader(wanted.data))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"t": wanted.topic, "k": wanted.key})

	ampub := new(AmPub)
	ampub.publisher = new(recordingPublisher)
	handler := http.HandlerFunc(ampub.api1Publish)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assertStatusOK(t, rr)
	assertContextNotNil(t, ampub.publisher.(*recordingPublisher))
	assertEquals(t, &wanted, ampub.publisher.(*recordingPublisher))
}

func assertStatusOK(t *testing.T, rr *httptest.ResponseRecorder) {
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func assertContextNotNil(t *testing.T, actual *recordingPublisher) {
	if actual.ctx == nil {
		t.Error("publisher received nil Context")
	}
}

func assertEquals(t *testing.T, wanted *recordingPublisher, actual *recordingPublisher) {
	if actual.topic != wanted.topic {
		t.Errorf("publisher received wrong topic: got %v want %v",
			actual.topic, wanted.topic)
	}

	if actual.key != wanted.key {
		t.Errorf("publisher received wrong key: got %v want %v",
			actual.key, wanted.key)
	}

	if actual.data != wanted.data {
		t.Errorf("publisher received wrong data: got %v want %v",
			actual.data, wanted.data)
	}
}

type recordingPublisher struct {
	ctx   context.Context
	topic string
	data  string
	key   string
}

func (record *recordingPublisher) Publish(ctx context.Context, topic string, key string, data []byte) error {
	record.ctx = ctx
	record.topic = topic
	record.key = key
	record.data = string(data)
	return nil
}
