package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"osrs-cli/helpers"
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
			formattedSkill := map[string]interface{}{
				"name":  skill.Name,
				"level": skill.Level,
				"xp":    helpers.FormatNumber(skill.XP),
				"rank":  helpers.FormatNumber(skill.Rank),
			}

			response, err := json.MarshalIndent(formattedSkill, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to format JSON: %w", err)
			}
			fmt.Println(string(response))
			return nil
		}
	}
	return fmt.Errorf("skill '%s' not found", skillName)
}

func GetAllSkills(rsData *Hiscores) error {
	response, err := json.MarshalIndent(rsData.Skills, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(response))
	return nil
}
