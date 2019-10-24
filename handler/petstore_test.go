package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/steveoc64/petstore/database/testdb"

	"github.com/sirupsen/logrus"
)

// TestPetstoreHTTP spins up the full RPC/HTTP servers, and does a mini-integration test
// across HTTP.  Provides test coverage for the full round trip, does not needed to be
// extended if the API changes
func TestPetstoreHTTP(t *testing.T) {
	petServer := NewPetstoreServer(logrus.New(), testdb.New(), testRPCPort, testRestPort, testAPIKey)

	t.Log("Testing Running the RPC/REST Listener - mini integration test")
	go petServer.Run()

	t.Log("Call the HTTP/REST endpoint after 200ms break")
	time.Sleep(time.Millisecond * 200)
	client := &http.Client{}
	for _, i := range []int{1, 2, 3, 4, 5} {
		url := fmt.Sprintf("http://localhost:%d/pet/%d", testRestPort, i)
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			t.Error("Error creating request", url, err.Error())
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
		//resp, err := http.Get(fmt.Sprintf("http://localhost:%d/pet/%d", testRestPort, i))
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
