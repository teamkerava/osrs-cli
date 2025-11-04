package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Skill struct {
	// ID    int    `json:"id"`
	Level int    `json:"level"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	XP    int    `json:"xp"`
}

func GetSkill(rsData *Hiscores, skillName string) error {
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

func AllSkills(rsData *Hiscores) error {
	response, err := json.MarshalIndent(rsData.Skills, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(response))
	return nil
}
