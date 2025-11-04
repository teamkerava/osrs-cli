package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Hiscores struct {
	Name       string     `json:"name"`
	Activities []Activity `json:"activities"`
	Skills     []Skill    `json:"skills"`
}

func Api(playerName string) (*Hiscores, error) {
	uri := "https://secure.runescape.com/m=hiscore_oldschool/index_lite.json?player=" + playerName

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		text := "Invalid player name: %s\nFailed to fetch data: status code %d"
		return nil, fmt.Errorf(text, playerName, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rsData Hiscores
	if err := json.Unmarshal(body, &rsData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &rsData, nil
}
