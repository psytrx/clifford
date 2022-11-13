package main

import (
	"clifford/pkg/huemint"
	"fmt"
)

func randomGradient() ([]string, error) {
	payload := huemint.Payload{
		NumColors:   6,
		Temperature: 1.3,
		NumResults:  1,
		Adjacency:   []string{"0", "35", "45", "55", "65", "75", "35", "0", "15", "25", "35", "45", "45", "15", "0", "15", "25", "35", "55", "25", "15", "0", "15", "25", "65", "35", "25", "15", "0", "15", "75", "45", "35", "25", "15", "0"},
		Palette:     []string{"-", "-", "-", "-", "-", "-"},
		Mode:        huemint.ModeTransformer,
	}

	res, err := huemint.Colors(payload)
	if err != nil {
		return nil, fmt.Errorf("could not get colors from Huemint: %w", err)
	}

	return res.Results[0].Palette, nil
}
