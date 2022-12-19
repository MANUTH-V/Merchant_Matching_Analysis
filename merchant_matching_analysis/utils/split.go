package utils

import (
	. "manuth/types"
	"os"

	"github.com/google/uuid"
)

func SplitMerchant(merchants []Merchant, count int) [][]Merchant {
	var (
		result = make([][]Merchant, count)
	)

	for i, merchant := range merchants {
		result[i%count] = append(result[i%count], merchant)
	}

	return result
}

func SaveQrSplitMerchant(merchants [][]Merchant, outputFolder string) {
	for _, merchant := range merchants {
		// create folder
		folderName := uuid.New().String()
		os.Mkdir(outputFolder+"/"+folderName, os.ModePerm)

		filename := outputFolder + "/" + folderName + "/" + "qr_cleaned.csv"
		SaveMerchant(merchant, filename)
	}
}
