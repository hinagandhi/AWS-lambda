package main
import (
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-lambda-go/events"
	"context"
	"os"
	"log" 
)

func handler(ctx context.Context, s3Event events.S3Event) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	log.Print("session", sess)
	if err != nil {
		log.Print("successfully created session")
	}
	for _, record := range s3Event.Records {
			s3_record := record.S3
			log.Print("s3=",s3_record)
			log.Print("s3 object", s3_record.Object)
			download_bucket := "hina-csv-files-bucket"
			upload_bucket := "hina-copy-csv-files"
			key := s3_record.Object.Key
			file_path := "/tmp/"+key
			log.Print("filepath",file_path)
			f, err := os.Create(file_path)
			log.Print("create file",f)
			downloader := s3manager.NewDownloader(sess)
			numBytes, err := downloader.Download(f,
				&s3.GetObjectInput{
					Bucket: aws.String(download_bucket),
					Key:    aws.String(key),
				})
			log.Print("numbytes ",numBytes, " f=", f)	
				
			if err != nil {
				log.Print("Unable to download item", f, err)
			}
			
			log.Print("uploading file",key," in bucket",upload_bucket, " body", f)
			uploader := s3manager.NewUploader(sess)
			_, err = uploader.Upload(&s3manager.UploadInput{
					 Bucket: aws.String(upload_bucket),
					 Key: aws.String(key),
					 Body: f,
	})
	if err != nil {
		// Print the error and exit.
		log.Print("Unable to upload", f, upload_bucket, err)
	}
	log.Print("Successfully uploaded", f, upload_bucket)
	}
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	log.Print("in main")
	lambda.Start(handler)
}