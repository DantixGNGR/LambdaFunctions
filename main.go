package main

//WHEN BUILDING, RUN: GOARCH=amd64 GOOS=linux go build main.go && zip lambda.zip main

import (
	"awesomeProject/assembly"
	"flag"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

var xlsxPath = flag.String("o", "Users/danicabezuidenhout/Documents/code/GolandProjects/lambdaFunction", "Path to the XLSX output file")// this is what will need to be pushed to the upload service?
var delimiter = flag.String("d", ",", "Delimiter for fields in the CSV input.")

func main() {
	lambda.Start(assembly.Handler(*xlsxPath, *delimiter))
	fmt.Println("it worked!!")
}
