package upload

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Bucket struct{}

func (b Bucket) UploadGameFile(bucketName string, fileName string, region string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("os.Open - filename: %s, err: %v", fileName, err)
	}
	defer file.Close()

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Name()),
		Body:   io.Reader(file),
		ACL:    "public-read",
	})

	if err != nil {
		log.Printf("Couldn't upload file %v to %v. Here's why: %v\n", file.Name(), bucketName, err)
	}

	return err
}
