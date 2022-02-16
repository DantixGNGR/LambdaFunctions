package main

//WHEN BUILDING, RUN: GOARCH=amd64 GOOS=linux go build main.go && zip lambda.zip main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)


func main() {
	lambda.Start(handler)
	fmt.Println("it worked!!")
}
