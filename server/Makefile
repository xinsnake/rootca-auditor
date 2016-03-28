build:
	go build -o main *.go

fetch: build
	./main --listen=false

s3: fetch
	aws --profile xinsnake s3 sync wwwroot/ s3://rootca-auditor/ --acl public-read
