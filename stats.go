package main

import (
	"encoding/json"
	"fmt"
)

type Stats struct {
	Score int    `json:"score"`
	Level int    `json:"level"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	XP    int    `json:"xp"`
}

func GetAllStats(rsData *Hiscores) error {
	response, err := json.MarshalIndent(rsData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(response))
	return nil
}
