package main

type Stats struct {
	Score int    `json:"score"`
	Level int    `json:"level"`
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	XP    int    `json:"xp"`
}

func GetAllStats(rsData *Hiscores) error {
	return CreateTable(rsData, TableAll)
}
