package utils

import (
	"encoding/csv"
	"io"
	"log"
	. "manuth/types"
	"os"
)

func CsvToMapByColumnName(filePath string, columnNames []string) []Merchant {
	// open csv file
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// parse the file
	r := csv.NewReader(csvFile)

	// read the header
	_, err = r.Read()
	if err != nil {
		log.Fatalln("Couldn't read the header", err)
	}

	// read the body
	var merchants []Merchant
	for {
		// read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Couldn't read the record", err)
		}

		// create a merchant
		merchant := Merchant{
			MerchantName: record[0],
		}

		// append to merchants
		merchants = append(merchants, merchant)
	}

	return merchants
}
