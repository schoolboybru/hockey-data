package transformer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/schoolboybru/hockey-data/cmd/data"
)

type Transformer struct{}

func (t *Transformer) CreateCSVFile(games []data.Game) (*os.File, error) {

	date := games[0].GameDate

	fileName := fmt.Sprintf("%s.csv", date)
	csvFile, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("Couldn't create csv file for game on %v. Here's why: %v\n", date, err)
	}

	csvWriter := csv.NewWriter(csvFile)

	err = gocsv.MarshalCSV(games, csvWriter)

	if err != nil {
		fmt.Printf("Couldn't marshal csv file for game on %v. Here's why: %v\n", date, err)
	}

	return csvFile, err
}

func (t *Transformer) DeleteCSVFile(fileName string) {
	err := os.Remove(fileName)

	if err != nil {
		log.Println("Failed to remove file:", err)
	}
}
