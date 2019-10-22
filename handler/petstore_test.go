package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

// TestPetstoreHTTP spins up the full RPC/HTTP servers, and does a mini-integration test
// across HTTP.  Provides test coverage for the full round trip, does not needed to be
// extended if the API changes
func TestPetstoreHTTP(t *testing.T) {
	petServer := NewPetstoreServer(logrus.New(), testRpcPort, testRestPort, testAPIKey)

	t.Log("Testing Running the RPC/REST Listener - mini integration test")
	go petServer.Run()

	t.Log("Call the HTTP/REST endpoint after 200ms break")
	time.Sleep(time.Millisecond * 200)

	for i := range []int{1, 2, 3, 4, 5} {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d/pet/%d", testRestPort, i))
		if err != nil {
			t.Errorf("TestPetstoreListen unexpected Error calling GET %v", err.Error())
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("TestPetstoreListen unexpected Error reading response %v", err.Error())
			resp.Body.Close()
			return
		}
		t.Logf("TestPetstoreListen got a valid response %#v", string(body))
		resp.Body.Close()
	}
}