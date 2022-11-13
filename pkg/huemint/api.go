package huemint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRandomGradient() ([]string, error) {
	reqPayload := reqPayload{
		NumColors:   6,
		Temperature: 1.3,
		NumResults:  1,
		Adjacency:   []string{"0", "35", "45", "55", "65", "75", "35", "0", "15", "25", "35", "45", "45", "15", "0", "15", "25", "35", "55", "25", "15", "0", "15", "25", "65", "35", "25", "15", "0", "15", "75", "45", "35", "25", "15", "0"},
		Palette:     []string{"-", "-", "-", "-", "-", "-"},
		Mode:        "transformer",
	}

	reqJson, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request payload: %w", err)
	}

	resp, err := http.Post("https://huemint.com/api/", "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("could not post HTTP request: %w", err)
	}

	var resBody respData
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		return nil, fmt.Errorf("could not decode response body: %w", err)
	}

	return resBody.Results[0].Palette, nil
}
