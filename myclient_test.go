package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	client "github.com/srikanthbhandary-teach/my-client"
	"github.com/stretchr/testify/assert"
)

func setupTestServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}

func createClientForTest(serverURL string) *client.Client {
	return client.NewClient(serverURL, "test-api-key")
}

func TestGetMyInfo(t *testing.T) {
	info := client.MyInfo{
		ID:   "1",
		Name: "Alice",
		Age:  30,
	}

	infoJSON, err := json.Marshal(info)
	assert.NoError(t, err)

	testServer := setupTestServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(infoJSON)
	}))

	defer testServer.Close()

	client := createClientForTest(testServer.URL)

	// Test for a single MyInfo
	myInfo, err := client.GetMyInfo("1")
	assert.NoError(t, err)
	assert.Len(t, myInfo, 1)
	assert.Equal(t, info, myInfo[0])

	// Test for an array of MyInfo
	myInfoArray, err := client.GetMyInfo("all")
	assert.NoError(t, err)
	assert.Len(t, myInfoArray, 1)
	assert.Equal(t, info, myInfoArray[0])
}
