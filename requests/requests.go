package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// // GoRequest represents a request that moves rover forward/backward.
// type GoRequest struct {
// 	// Speed in km/h. Negative speeds means that the rover is going backward.
// 	Speed int `json:"speed"`
// }

// // TurnRequest represents a request that makes the rover turn.
// // Bear in mind that the rover should be moving while turning.
// type TurnRequest struct {
// 	// Negative degrees meant that the rover is turning left, positive – right.
// 	Degrees int `json:"degrees"`
// }

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
