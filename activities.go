package main

import (
	"encoding/json"
	"fmt"
)

type Activity struct {
	// ID    int    `json:"id"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Score int    `json:"score"`
}

func Activities(rsData *Hiscores) error {
	response, err := json.MarshalIndent(rsData.Activities, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(response))
	return nil
}
