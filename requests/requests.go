package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	return nil
}
