package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostRequest(endpointURL string, data map[string]interface{}) error {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal json: %v", err)
	}

	res, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return fmt.Errorf("http post: %v", err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("post request failed: status: %d, body: %s", res.StatusCode, b)
	}

	return nil
}
