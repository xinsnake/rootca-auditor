build:
	go build -o main *.go

s3:
	aws --profile xinsnake s3 sync wwwroot/ s3://rootca-auditor/ --acl public-read
