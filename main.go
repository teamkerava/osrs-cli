package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"osrs-cli/helpers"
)

const help = `
Usage: ./osrs-cli <player_name> [flags]

Flags:
  -/--activities         # Show all activities (bosses, minigames, etc.)
  -/--skills             # Show all skills only
  -/--activity string    # Show specific activity (e.g., "Kalphite Queen")
  -/--skill string       # Show specific skill (e.g., "Woodcutting")

Examples:
  ./osrs-cli "Manly Bacon"                             # Show all stats (default)
  ./osrs-cli "Manly Bacon" --activities                # Show only activities
  ./osrs-cli "Manly Bacon" --skills                    # Show only skills
  ./osrs-cli "Manly Bacon" --activity "Kalphite Queen" # Show specific activity
  ./osrs-cli "Manly Bacon" --skill "Woodcutting"       # Show specific skill`

func Name(rsData *Hiscores) (string, error) {
	formattedName := helpers.FormatName(rsData.Name)
	return formattedName, nil
}

func main() {
	// Check if we have at least a player name
	if len(os.Args) < 2 {
		fmt.Println(help)
		return
	}

	// First argument is always the player name
	playerName := os.Args[1]

	// Parse flags from the remaining arguments (after player name)
	flagSet := flag.NewFlagSet("osrs-cli", flag.ExitOnError)
	showActivities := flagSet.Bool("activities", false, "Show all activities (bosses, minigames, etc.)")
	showSkills := flagSet.Bool("skills", false, "Show all skills")
	activityName := flagSet.String("activity", "", "Show specific activity (e.g., 'Kalphite Queen')")
	skillName := flagSet.String("skill", "", "Show specific skill (e.g., 'Woodcutting')")

	// Parse flags from arguments after player name
	if len(os.Args) > 2 {
		flagSet.Parse(os.Args[2:])
	}

	// Replace spaces with underscores for the API request
	replacedPlayerName := strings.ReplaceAll(playerName, " ", "_")
	rsData, err := Api(replacedPlayerName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	name, err := Name(rsData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Player Name:", name)

	// Handle different display options based on flags
	switch {
	case *showActivities:
		err = GetAllActivities(rsData)
	case *showSkills:
		err = GetAllSkills(rsData)
	case *activityName != "":
		*activityName = helpers.FormatName(*activityName)
		err = GetActivity(rsData, *activityName)
	case *skillName != "":
		*skillName = helpers.FormatName(*skillName)
		err = GetSkill(rsData, *skillName)
	default:
		err = GetAllStats(rsData)
	}

	if err != nil {
		fmt.Println("Error:", err)
	}

}
