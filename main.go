package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/schoolboybru/hockey-data/cmd/data"
	"github.com/schoolboybru/hockey-data/cmd/transformer"
	"github.com/schoolboybru/hockey-data/cmd/upload"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	bucketName := os.Getenv("BUCKET_NAME")
	region := os.Getenv("AWS_REGION")

	d := data.GameDayData{}

	d.Get()

	t := &transformer.Transformer{}

	csv, err := t.CreateCSVFile(d.Games)

	if err != nil {
		fmt.Println(err)
	}

	b := &upload.Bucket{}

	err = b.UploadGameFile(bucketName, csv.Name(), region)

	if err != nil {
		fmt.Println(err)
	}

	t.DeleteCSVFile(csv.Name())

}
