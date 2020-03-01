# AWS-lambda

# Summary:
Two files are written, the file upload-csv.go is to upload any CSV to s3 bucket which will trigger AWS lambda. Second file lambda-function.go is
to AWS lambda code that will copy file uploaded to other s3 bucket. This was just a simple code that I wanted to wite to learn about AWS lambda and Go.

# Steps taken to run AWS lambda:
1. Add a trigger 
  a) Select s3 from drop down
  b) Add bucket name, event type(in this case put) and then after filling all details, click on add
  
2. Handler name should be the file name. In this case, lambda-function
3. GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o lambda-function lambda-functio.go -> building binary file
4. zip the file `zip lambda.zip lambda-function`
5. Upload the zip file
6. Create permissions for function by going to:
  a) https://console.aws.amazon.com/iam/home#/roles and create a new role and attach policies: `AmazonS3FullAccess` 
  and `AWSLambdaBasicExecutionRole` and other policy which will have policy to create log stream and put log events.
  b) Once role is created then select that role in execution role setion.
  


