package test

import (
	"os"
	"testing"
	"time"

	"github.com/schoolboybru/hockey-data/cmd/data"
	"github.com/schoolboybru/hockey-data/cmd/transformer"
)

func createCSVFileTest(t *testing.T) {
	tr := transformer.Transformer{}

	currDate := time.Now().Format("2006-01-02")

	games := []data.Game{
		{
			GameDate: currDate,
			AwayTeam: data.Team{
				Abbreviation: "Tor",
				Name:         data.Name{Default: "Toronto Maple Leafs"},
				Score:        0,
				ShotOnGoal:   30,
			},
			HomeTeam: data.Team{
				Abbreviation: "Edm",
				Name:         data.Name{Default: "Edmonton Oilers"},
				Score:        3,
				ShotOnGoal:   30,
			},
			GameOutcome: data.GameOutcome{LastPeriodType: "Regulation"},
			Goals: []data.Goal{
				{
					Name:     data.Name{Default: "Connor Mcdavid"},
					Period:   1,
					Strength: "Even",
					Time:     "17:05",
				},
				{
					Name:     data.Name{Default: "Leon Draisaitl"},
					Period:   2,
					Strength: "Power Play",
					Time:     "10:11",
				},
				{
					Name:     data.Name{Default: "Zach Hyman"},
					Period:   3,
					Strength: "Even",
					Time:     "2:05",
				},
			},
		},
		{
			GameDate: currDate,
			AwayTeam: data.Team{
				Abbreviation: "Bos",
				Name:         data.Name{Default: "Boston Bruins"},
				Score:        2,
				ShotOnGoal:   33,
			},
			HomeTeam: data.Team{
				Abbreviation: "Det",
				Name:         data.Name{Default: "Detroit Red Wings"},
				Score:        1,
				ShotOnGoal:   28,
			},
			GameOutcome: data.GameOutcome{LastPeriodType: "Regulation"},
			Goals: []data.Goal{
				{
					Name:     data.Name{Default: "David Pastrnak"},
					Period:   1,
					Strength: "Even",
					Time:     "9:05",
				},
				{
					Name:     data.Name{Default: "David Pastrnak"},
					Period:   1,
					Strength: "Even",
					Time:     "1:35",
				},
				{
					Name:     data.Name{Default: "Klim Kostin"},
					Period:   3,
					Strength: "Even Strength",
					Time:     "7:40",
				},
			},
		},
	}

	file, err := tr.CreateCSVFile(games)

	if err != nil {
		t.Error(err)
	}

	file.Close()

	os.Remove(file.Name())

}
