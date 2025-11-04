# osrs-cli

A simple Go command-line tool for fetching Old School RuneScape hiscores.

## Features
- Fetches player stats from OSRS hiscores
- Supports activities, skills, and individual skill lookup
## Usage

```bash
./osrs-cli <player_name> [flags]
```

### Flags
- `-/--activities`   Show all activities (bosses, minigames, etc.)
- `-/--skills`       Show all skills only
- `-/--skill NAME`   Show specific skill (e.g., "Woodcutting")

### Examples
```bash
./osrs-cli "Manly Bacon"                              # Show all stats (default)
./osrs-cli "Manly Bacon" --skills                     # Show only skills
./osrs-cli "Manly Bacon" --activities                 # Show only activities
./osrs-cli "Manly Bacon" --activity "Kalphite Queen"  # Show specific activity
./osrs-cli "Manly Bacon" --skill "Woodcutting"        # Show specific skill
```

## Build

```
go build -o osrs-cli
```

## License
MIT
