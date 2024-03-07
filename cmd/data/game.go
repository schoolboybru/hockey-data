package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var url = "https://api-web.nhle.com/v1/score/"

type Name struct {
	Default string `json:"default"`
}

type Team struct {
	Abbreviation string `json:"abbrev"`
	Name         Name   `json:"name"`
	Score        int    `json:"score"`
	ShotOnGoal   int    `json:"sog"`
}

type GameOutcome struct {
	LastPeriodType string `json:"lastPeriodType"`
}

type Goal struct {
	Name     Name   `json:"name"`
	Period   int    `json:"period"`
	Strength string `json:"strength"`
	Time     string `json:"timeInPeriod"`
}

type Game struct {
	GameDate    string      `json:"gameDate"`
	AwayTeam    Team        `json:"awayTeam"`
	HomeTeam    Team        `json:"homeTeam"`
	GameOutcome GameOutcome `json:"gameOutcome"`
	Goals       []Goal      `json:"goals"`
}

type GameDayData struct {
	CurrentDate string `json:"currentDate"`
	Games       []Game `json:"games"`
}

func (g *GameDayData) Get() error {

	currDate := time.Now().Format("2006-01-02")

	req := fmt.Sprintf("%s%s", url, currDate)

	resp, err := http.Get(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &g)

	if err != nil {
		return err
	}

	return nil
}
