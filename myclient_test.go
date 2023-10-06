package myclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MyInfo struct {
	ID   string `json:"number"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func setupMockServer(handler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(handler)
}

func TestClient_CreateMyInfo(t *testing.T) {
	tests := []struct {
		id, name string
		age      int
		expected int
	}{
		{"1", "Alice", 30, http.StatusOK},
		{"2", "Bob", 25, http.StatusOK},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("id=%s name=%s age=%d", test.id, test.name, test.age), func(t *testing.T) {
			mockServer := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
					return
				}
				w.WriteHeader(http.StatusOK) // Use StatusOK for simplicity in this example
			})
			defer mockServer.Close()

			client := NewClient("mock-api-key")
			baseURL = mockServer.URL

			err := client.CreateMyInfo(test.id, test.name, test.age)
			if err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}

			if got, want := test.expected, http.StatusOK; got != want {
				t.Errorf("Expected status code %d, got %d", want, got)
			}
		})
	}
}
