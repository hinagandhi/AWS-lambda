package main
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "fmt"
	"os"
)
func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}	

func main(){

	if len(os.Args) != 2 {
		exitErrorf("file name required: %s filename",
			os.Args[1])
	}
	
	filename := os.Args[1]
	bucket := "hina-csv-files-bucket"

	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}
	
	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key: aws.String(filename),
		Body: file,
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}
	
	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}	