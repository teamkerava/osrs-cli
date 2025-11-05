package main

import (
	"fmt"
	"strings"
)

type Activity struct {
	// ID    int    `json:"id"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Score int    `json:"score"`
}

func GetAllActivities(rsData *Hiscores) error {
	return CreateTable(rsData, TableActivities)
}

func GetActivity(rsData *Hiscores, activityName string) error {
	for _, activity := range rsData.Activities {
		if strings.EqualFold(activity.Name, activityName) {
			return CreateTable(&Hiscores{Activities: []Activity{activity}}, TableActivities)
		}
	}
	return fmt.Errorf("Activity '%s' not found", activityName)
}
