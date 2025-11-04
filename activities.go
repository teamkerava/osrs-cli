package main

import (
	"encoding/json"
	"fmt"
	"osrs-cli/helpers"
)

type Activity struct {
	// ID    int    `json:"id"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Score int    `json:"score"`
}

func GetAllActivities(rsData *Hiscores) error {
	response, err := json.MarshalIndent(rsData.Activities, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	fmt.Println(string(response))
	return nil
}

func GetActivity(rsData *Hiscores, activityName string) error {
	for _, activity := range rsData.Activities {
		if activity.Name == activityName {
			formattedActivity := map[string]interface{}{
				"name":  activity.Name,
				"score": helpers.FormatNumber(activity.Score),
				"rank":  helpers.FormatNumber(activity.Rank),
			}

			response, err := json.MarshalIndent(formattedActivity, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to format JSON: %w", err)
			}
			fmt.Println(string(response))
			return nil
		}
	}
	return fmt.Errorf("activity '%s' not found", activityName)
}
