package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"

	"encoding/csv"
	"os"

	"github.com/tealeg/xlsx"
)

var (
	s3session *s3.S3
)

const (
	BUCKET_NAME = "sol-dev-output"
	REGION = "eu-central-1"
	KEY = "TruCape-Invoices/TRU_Invoices.csv"
)

var xlsxPath = flag.String("o", "", "Path to the XLSX output file")// this is what will need to be pushed to the upload service?
var delimiter = flag.String("d", ",", "Delimiter for fields in the CSV input.")

func handler(delimiter string, XLSXPath string) error {
	//Create S3 session
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
				Region: aws.String(REGION),
			})))
	//Get output file from s3 output bucket
	fmt.Println("Printing: ", KEY)

		resp, err := s3session.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(BUCKET_NAME),
			Key: aws.String(KEY),
		})

		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)// is the body, bytes?
		csvBody := string(body[:])
		fmt.Println(string(body))
		if err != nil {
			panic(err)
		}

		// Convert CSV file to Excel file - don't know if this is the right approach
	csvFile, err := os.Open(csvBody)
	if err != nil {
		//return error
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	} else {
		reader.Comma = rune(',')
	}
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet(csvBody)
	if err != nil {
		//return err
	}
	fields, err := reader.Read()
	for err == nil {
		row := sheet.AddRow()
		for _, field := range fields {
			cell := row.AddCell()
			cell.Value = field
		}
		fields, err = reader.Read()
	}
	if err != nil {
		fmt.Printf(err.Error())
	}

	return xlsxFile.Save(XLSXPath)

}
