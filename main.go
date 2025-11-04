package main

import (
	"fmt"
	"os"
	"strings"
)

const help = `
Usage: ./rs-cli <player_name> [skill_name]

Examples:
  ./rs-cli "Manly Bacon"               # Show player info
  ./rs-cli "Manly Bacon" "Activities"  # Show all activities (bosses, minigames, etc.)
  ./rs-cli "Manly Bacon" "Skills"      # Show all skills (only skills)
  ./rs-cli "Manly Bacon" "Stats"       # Show all hiscore stats
  ./rs-cli "Manly Bacon" "Woodcutting" # Optional: Show specific skill`

func Name(rsData *Hiscores) error {
	fmt.Println("Player Name:", rsData.Name)
	return nil
}

func main() {
	// TODO: Use a proper CLI parser like cobra or urfave/cli for better argument handling
	if len(os.Args) < 2 {
		fmt.Println(help)
		return
	}

	playerName := os.Args[1]
	var skillName string
	if len(os.Args) > 2 {
		skillName = os.Args[2]
	}

	// Replace spaces with underscores for the API request
	replaced_player := strings.ReplaceAll(playerName, " ", "_")
	rsData, err := Api(replaced_player)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = Name(rsData)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if skillName == "" {
		err, err = AllSkills(rsData), Activities(rsData)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// Call specific skill if provided
	if skillName != "" {
		err = GetSkill(rsData, skillName)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}
