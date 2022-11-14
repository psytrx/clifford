package main

import (
	"clifford/pkg/huemint"
	"fmt"
)

func randomGradient() ([]string, error) {
	return randomGradientBaseColor("-")
}

func randomGradientBaseColor(baseHex string) ([]string, error) {
	payload := huemint.Payload{
		NumColors:   6,
		Temperature: 1.3,
		NumResults:  1,
		Adjacency:   []string{"0", "45", "55", "65", "75", "95", "45", "0", "15", "25", "35", "55", "55", "15", "0", "15", "25", "45", "65", "25", "15", "0", "15", "35", "75", "35", "25", "15", "0", "25", "95", "55", "45", "35", "25", "0"},
		Palette:     []string{baseHex, "-", "-", "-", "-", "-"},
		Mode:        huemint.ModeTransformer,
	}

	res, err := huemint.Colors(payload)
	if err != nil {
		return nil, fmt.Errorf("could not get colors from Huemint: %w", err)
	}

	return res.Results[0].Palette, nil
}
