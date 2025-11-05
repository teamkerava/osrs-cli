package main

import (
	"fmt"
	"strings"

	"osrs-cli/helpers"
)

type TableType string

const (
	TableAll        TableType = "all"
	TableSkills     TableType = "skills"
	TableActivities TableType = "activities"
)

func CreateTable(rsData *Hiscores, tableType TableType) error {
	showSkills := tableType == TableAll || tableType == TableSkills
	showActivities := tableType == TableAll || tableType == TableActivities

	// Initialize column widths with header lengths
	maxNameWidth := len("Name")
	maxLevelWidth := len("Level")
	maxXPWidth := len("XP")
	maxRankWidth := len("Rank")
	maxScoreWidth := len("Score")

	// Calculate maximum widths from skills
	if showSkills {
		for _, skill := range rsData.Skills {
			maxNameWidth = helpers.MaxLen(maxNameWidth, skill.Name)
			maxLevelWidth = helpers.MaxLen(maxLevelWidth, fmt.Sprintf("%d", skill.Level))
			maxXPWidth = helpers.MaxLen(maxXPWidth, helpers.FormatNumber(skill.XP))
			maxRankWidth = helpers.MaxLen(maxRankWidth, helpers.FormatNumber(skill.Rank))
		}
	}

	// Calculate maximum widths from activities
	if showActivities {
		for _, activity := range rsData.Activities {
			maxNameWidth = helpers.MaxLen(maxNameWidth, activity.Name)
			maxRankWidth = helpers.MaxLen(maxRankWidth, helpers.FormatNumber(activity.Rank))
			maxScoreWidth = helpers.MaxLen(maxScoreWidth, helpers.FormatNumber(activity.Score))
		}
	}

	var topBorder, middleBorder, bottomBorder string
	var headerFormat, skillFormat, activityFormat string

	switch tableType {
	case TableActivities:
		// 3 columns: Name, Rank, Score
		topBorder = "┌" + strings.Repeat("─", maxNameWidth+2) + "┬" +
			strings.Repeat("─", maxRankWidth+2) + "┬" +
			strings.Repeat("─", maxScoreWidth+2) + "┐"

		middleBorder = "├" + strings.Repeat("─", maxNameWidth+2) + "┼" +
			strings.Repeat("─", maxRankWidth+2) + "┼" +
			strings.Repeat("─", maxScoreWidth+2) + "┤"

		bottomBorder = "└" + strings.Repeat("─", maxNameWidth+2) + "┴" +
			strings.Repeat("─", maxRankWidth+2) + "┴" +
			strings.Repeat("─", maxScoreWidth+2) + "┘"

		headerFormat = "│ %-*s │ %-*s │ %-*s │\n"
		activityFormat = "│ %-*s │ %*s │ %*s │\n"
	case TableSkills:
		// 4 columns: Name, Level, XP, Rank
		topBorder = "┌" + strings.Repeat("─", maxNameWidth+2) + "┬" +
			strings.Repeat("─", maxLevelWidth+2) + "┬" +
			strings.Repeat("─", maxXPWidth+2) + "┬" +
			strings.Repeat("─", maxRankWidth+2) + "┐"

		middleBorder = "├" + strings.Repeat("─", maxNameWidth+2) + "┼" +
			strings.Repeat("─", maxLevelWidth+2) + "┼" +
			strings.Repeat("─", maxXPWidth+2) + "┼" +
			strings.Repeat("─", maxRankWidth+2) + "┤"

		bottomBorder = "└" + strings.Repeat("─", maxNameWidth+2) + "┴" +
			strings.Repeat("─", maxLevelWidth+2) + "┴" +
			strings.Repeat("─", maxXPWidth+2) + "┴" +
			strings.Repeat("─", maxRankWidth+2) + "┘"

		headerFormat = "│ %-*s │ %-*s │ %-*s │ %-*s │\n"
		skillFormat = "│ %-*s │ %*d │ %*s │ %*s │\n"
	default:
		// 5 columns: Name, Level, XP, Rank, Score
		topBorder = "┌" + strings.Repeat("─", maxNameWidth+2) + "┬" +
			strings.Repeat("─", maxLevelWidth+2) + "┬" +
			strings.Repeat("─", maxXPWidth+2) + "┬" +
			strings.Repeat("─", maxRankWidth+2) + "┬" +
			strings.Repeat("─", maxScoreWidth+2) + "┐"

		middleBorder = "├" + strings.Repeat("─", maxNameWidth+2) + "┼" +
			strings.Repeat("─", maxLevelWidth+2) + "┼" +
			strings.Repeat("─", maxXPWidth+2) + "┼" +
			strings.Repeat("─", maxRankWidth+2) + "┼" +
			strings.Repeat("─", maxScoreWidth+2) + "┤"

		bottomBorder = "└" + strings.Repeat("─", maxNameWidth+2) + "┴" +
			strings.Repeat("─", maxLevelWidth+2) + "┴" +
			strings.Repeat("─", maxXPWidth+2) + "┴" +
			strings.Repeat("─", maxRankWidth+2) + "┴" +
			strings.Repeat("─", maxScoreWidth+2) + "┘"

		headerFormat = "│ %-*s │ %-*s │ %-*s │ %-*s │ %-*s │\n"
		skillFormat = "│ %-*s │ %*d │ %*s │ %*s │ %-*s │\n"
		activityFormat = "│ %-*s │ %-*s │ %-*s │ %*s │ %*s │\n"
	}

	fmt.Println("\n" + topBorder)
	switch tableType {
	case TableActivities:
		fmt.Printf(headerFormat,
			maxNameWidth, "Name",
			maxRankWidth, "Rank",
			maxScoreWidth, "Score")
	case TableSkills:
		fmt.Printf(headerFormat,
			maxNameWidth, "Name",
			maxLevelWidth, "Level",
			maxXPWidth, "XP",
			maxRankWidth, "Rank")
	default:
		fmt.Printf(headerFormat,
			maxNameWidth, "Name",
			maxLevelWidth, "Level",
			maxXPWidth, "XP",
			maxRankWidth, "Rank",
			maxScoreWidth, "Score")
	}
	fmt.Println(middleBorder)

	if showSkills {
		for _, skill := range rsData.Skills {
			if tableType == TableSkills {
				fmt.Printf(skillFormat,
					maxNameWidth, skill.Name,
					maxLevelWidth, skill.Level,
					maxXPWidth, helpers.FormatNumber(skill.XP),
					maxRankWidth, helpers.FormatNumber(skill.Rank))
			} else {
				fmt.Printf(skillFormat,
					maxNameWidth, skill.Name,
					maxLevelWidth, skill.Level,
					maxXPWidth, helpers.FormatNumber(skill.XP),
					maxRankWidth, helpers.FormatNumber(skill.Rank),
					maxScoreWidth, "-")
			}
		}
	}

	// Print separator before activities if showing both
	if showSkills && showActivities && len(rsData.Activities) > 0 {
		fmt.Println(middleBorder)
	}

	if showActivities {
		for _, activity := range rsData.Activities {
			if tableType == TableActivities {
				fmt.Printf(activityFormat,
					maxNameWidth, activity.Name,
					maxRankWidth, helpers.FormatNumber(activity.Rank),
					maxScoreWidth, helpers.FormatNumber(activity.Score))
			} else {
				fmt.Printf(activityFormat,
					maxNameWidth, activity.Name,
					maxLevelWidth, "-",
					maxXPWidth, "-",
					maxRankWidth, helpers.FormatNumber(activity.Rank),
					maxScoreWidth, helpers.FormatNumber(activity.Score))
			}
		}
	}

	fmt.Println(bottomBorder)

	return nil
}
