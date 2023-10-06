package myclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:8080"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

// CreateMyInfo sends a POST request to create a new MyInfo entity.
func (c *Client) CreateMyInfo(id, name string, age int) error {
	url := fmt.Sprintf("%s/?id=%s", baseURL, id)

	info := map[string]interface{}{
		"number": id,
		"name":   name,
		"age":    age,
	}

	payload, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("error marshaling request payload: %v", err)
	}

	return c.sendRequest("POST", url, payload)
}

// GetMyInfo sends a GET request to retrieve a MyInfo entity by its ID.
func (c *Client) GetMyInfo(id string) ([]byte, error) {
	url := fmt.Sprintf("%s/?id=%s", baseURL, id)

	return nil, c.sendRequest("GET", url, nil)
}

// UpdateMyInfo sends a PUT request to update an existing MyInfo entity.
func (c *Client) UpdateMyInfo(id, name string, age int) error {
	url := fmt.Sprintf("%s/?id=%s", baseURL, id)

	info := map[string]interface{}{
		"number": id,
		"name":   name,
		"age":    age,
	}

	payload, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("error marshaling request payload: %v", err)
	}

	return c.sendRequest("PUT", url, payload)
}

// DeleteMyInfo sends a DELETE request to delete a MyInfo entity by its ID.
func (c *Client) DeleteMyInfo(id string) error {
	url := fmt.Sprintf("%s/?id=%s", baseURL, id)

	return c.sendRequest("DELETE", url, nil)
}

func (c *Client) sendRequest(method, url string, payload []byte) error {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("error creating %s request: %v", method, err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending %s request: %v", method, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
