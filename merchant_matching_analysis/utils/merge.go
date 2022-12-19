package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/schollz/progressbar/v3"
)

func MergeCsvFiles(pwFilePath string, qrFilePath string, outputPath string) {

	// read csv file
	pw := CsvToMapByColumnName(pwFilePath, []string{"merchant_name"})
	qr := CsvToMapByColumnName(qrFilePath, []string{"merchant_name"})

	bar := progressbar.Default(
		int64(len(pw)*len(qr)),
		"Merging merchants",
	)

	// write to csv
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write header
	_ = writer.Write([]string{"qr_merchant_name", "pw_merchant_name"})

	// write body
	for _, pwMerchant := range pw {
		for _, qrMerchant := range qr {
			err = writer.Write([]string{qrMerchant.MerchantName, pwMerchant.MerchantName})
			bar.Add(1)
			if err != nil {
				log.Fatalln("Couldn't write to the csv file", err)
			}
		}
	}

	fmt.Println("Done")
}
