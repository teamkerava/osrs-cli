package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Activities struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Score int    `json:"score"`
}

type Skill struct {
	ID    int    `json:"id"`
	Level int    `json:"level"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	XP    int    `json:"xp"`
}

type RSResponse struct {
	Name       string       `json:"name"`
	Activities []Activities `json:"activities"`
	Skills     []Skill      `json:"skills"`
}

func rs_api(player string) (*RSResponse, error) {
	uri := "https://secure.runescape.com/m=hiscore_oldschool/index_lite.json?player=" + player

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}
	// Process the response body as needed
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON into RSResponse struct
	var rsData RSResponse
	if err := json.Unmarshal(body, &rsData); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &rsData, nil
}

func rs_name(rsData *RSResponse) error {
	fmt.Println("Player Name:", rsData.Name)
	return nil
}

func rs_skills(rsData *RSResponse) error {
	// Print only the skills part as nicely indented JSON
	response, err := json.MarshalIndent(rsData.Skills, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(response))

	return nil
}

func rs_activities(rsData *RSResponse) error {
	response, err := json.MarshalIndent(rsData.Activities, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(response))
	return nil
}

func rs_one_skill(rsData *RSResponse, skillName string) error {
	for _, skill := range rsData.Skills {
		if strings.EqualFold(skill.Name, skillName) {
			response, err := json.MarshalIndent(skill, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to format JSON: %w", err)
			}
			fmt.Println(string(response))
			return nil
		}
	}
	return nil
}

func main() {
	player := "Manly Bacon"
	replaced_player := strings.ReplaceAll(player, " ", "_")
	rsData, err := rs_api(replaced_player)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Now you can pass rsData to any function you want
	err = rs_name(rsData)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// err = rs_skills(rsData)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// err = rs_activities(rsData)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	err = rs_one_skill(rsData, "Woodcutting")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
