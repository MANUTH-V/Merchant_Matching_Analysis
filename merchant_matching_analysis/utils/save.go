package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	. "manuth/types"
	"os"
)

func SaveMerchant(merchants []Merchant, outputPath string) {
	// save result to csv
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write header
	_ = writer.Write([]string{"merchant_name"})

	// write body
	for _, merchant := range merchants {
		err = writer.Write([]string{merchant.MerchantName})
		if err != nil {
			log.Fatalln("Couldn't write to the csv file", err)
		}
	}
}

func SaveResult(result []MerchantSimilarityResult, outputPath string) {
	// save result to csv
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalln("Couldn't create the csv file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write header
	_ = writer.Write([]string{"qr_marchant_name", "pw_marchant_name", "similarity_score"})

	// write body
	for _, merchantSimilarityResult := range result {
		err = writer.Write([]string{
			merchantSimilarityResult.QrMarchantName,
			merchantSimilarityResult.PwMarchantName,
			fmt.Sprintf("%.2f", merchantSimilarityResult.SimilarityScore)})
		if err != nil {
			log.Fatalln("Couldn't write to the csv file", err)
		}
	}
}
