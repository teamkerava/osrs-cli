package main

import (
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

func GetAllSkills(rsData *Hiscores) error {
	return CreateTable(rsData, TableSkills)
}

func GetSkill(rsData *Hiscores, skillName string) error {
	for _, skill := range rsData.Skills {
		if strings.EqualFold(skill.Name, skillName) {
			return CreateTable(&Hiscores{Skills: []Skill{skill}}, TableSkills)
		}
	}
	return fmt.Errorf("Skill '%s' not found", skillName)
}
