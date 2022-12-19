package main

import (
	. "manuth/utils"
)

// --------------------------------------------------------------------------------
// Merge
// func main() {
// 	var (
// 		qrFilePath = "qr.csv"
// 		pwFilePath = "pw.csv"
// 		outputPath = "merged.csv"
// 	)

// 	// merge csv files
// 	MergeCsvFiles(pwFilePath, qrFilePath, outputPath)
// }

// --------------------------------------------------------------------------------
// Compare
func main() {
	var (
		qrFilePath = "qr_cleaned.csv"
		pwFilePath = "pw_cleaned.csv"
		columnName = "merchant_name"
		outputPath = "result.csv"
	)

	// read csv file
	pw := CsvToMapByColumnName(pwFilePath, []string{columnName})
	qr := CsvToMapByColumnName(qrFilePath, []string{columnName})

	// compare merchant
	result := CompareMerchant(pw, qr)

	// save result to csv
	SaveResult(result, outputPath)
}

// --------------------------------------------------------------------------------
// Clean pw
// func main() {
// 	var (
// 		inputPath  = "pw.csv"
// 		outputPath = "pw_cleaned.csv"
// 	)

// 	// read csv file
// 	pw := CsvToMapByColumnName(inputPath, []string{"merchant_name"})

// 	// clean merchant
// 	pw = CleanPwMerchant(pw)

// 	// save result to csv
// 	SaveMerchant(pw, outputPath)
// }

// --------------------------------------------------------------------------------
// Split pw
// func main() {
// 	var (
// 		inputPath  = "qr_cleaned.csv"
// 		outputPath = "executions"
// 	)

// 	// read csv file
// 	qr := CsvToMapByColumnName(inputPath, []string{"merchant_name"})

// 	// split merchant
// 	results := SplitMerchant(qr, 20)

// 	// save result to csv
// 	SaveQrSplitMerchant(results, outputPath)
// }
