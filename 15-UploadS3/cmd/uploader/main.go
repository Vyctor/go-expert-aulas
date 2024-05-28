package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"",
				"",
				"",
			),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "vyctor-guimaraes-go-expert-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 20)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case file := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(file, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s\n", completeFileName)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", completeFileName, err)
		errorFileUpload <- completeFileName
		<-uploadControl
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %s\n", completeFileName, err)
		errorFileUpload <- completeFileName
		<-uploadControl
		return
	}
	<-uploadControl
	fmt.Printf("File %s uploaded successfully\n", completeFileName)
}
