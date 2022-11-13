package huemint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ModeTransformer = "transformer"
	ModeDiffusion   = "diffusion"
	ModeRandom      = "random"
)

type Mode string

type Payload struct {
	NumColors   int      `json:"num_colors"`
	Temperature float64  `json:"temperature"`
	NumResults  int      `json:"num_results"`
	Adjacency   []string `json:"adjacency"`
	Palette     []string `json:"palette"`
	Mode        Mode     `json:"mode"`
}

type Response struct {
	Results []struct {
		Palette []string `json:"palette"`
		Score   float64  `json:"score"`
	} `json:"results"`
}

func Colors(payload Payload) (*Response, error) {
	reqJson, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request payload: %w", err)
	}

	resp, err := http.Post("https://huemint.com/api/", "application/json", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("could not post HTTP request: %w", err)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("could not decode response body: %w", err)
	}

	return &response, nil
}
