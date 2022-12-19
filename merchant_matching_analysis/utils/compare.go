package utils

import (
	. "manuth/types"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/exp/slices"
)

// func CompareMerchantOneQrToOnePw(qr Merchant, pw Merchant) {
// 	// check if results folder exist
// 	if _, err := os.Stat("results"); os.IsNotExist(err) {
// 		// create results folder
// 		err := os.Mkdir("results", 0755)
// 		if err != nil {
// 			log.Fatalln("Couldn't create results folder", err)
// 		}
// 	}

// 	filename := "results/" + qr.MerchantName + "_" + pw.MerchantName + time.Now().Format("20060102150405") + ".csv"
// 	// replace space with underscore
// 	filename = strings.ReplaceAll(filename, " ", "_")

// 	if strings.EqualFold(qr.MerchantName, pw.MerchantName) {
// 		// save to csv
// 		SaveResult([]MerchantSimilarityResult{
// 			{
// 				QrMarchantName:  qr.MerchantName,
// 				PwMarchantName:  pw.MerchantName,
// 				SimilarityScore: 1,
// 			},
// 		}, filename)
// 	} else {
// 		// check if merchant name is similar
// 		similarityScore := strutil.Similarity(qr.MerchantName, pw.MerchantName, metrics.NewJaroWinkler())

// 		if similarityScore > 0.9 {
// 			// save to csv
// 			SaveResult([]MerchantSimilarityResult{
// 				{
// 					QrMarchantName:  qr.MerchantName,
// 					PwMarchantName:  pw.MerchantName,
// 					SimilarityScore: similarityScore,
// 				},
// 			}, filename)
// 		}
// 	}
// }

// func CompareMerchantOneQrToListPw(qr Merchant, pw []Merchant, bar *progressbar.ProgressBar) {
// 	for _, pwMerchant := range pw {
// 		// compare merchant one by one concurrently
// 		go CompareMerchantOneQrToOnePw(qr, pwMerchant)
// 		bar.Add(1)
// 	}
// }

// func CompareMerchantListQrToListPw(qr []Merchant, pw []Merchant) {
// 	bar := progressbar.Default(
// 		int64(len(qr)*len(pw)),
// 		"Comparing merchants",
// 	)
// 	for _, qrMerchant := range qr {
// 		// compare merchant one by one concurrently
// 		go CompareMerchantOneQrToListPw(qrMerchant, pw, bar)
// 	}
// }

func CompareMerchant(pw []Merchant, qr []Merchant) []MerchantSimilarityResult {
	bar := progressbar.Default(
		int64(len(pw)*len(qr)),
		"Comparing merchants",
	)
	result := []MerchantSimilarityResult{}
	existedName := []string{}

	// loop through pw
	for _, pwMerchant := range pw {
		// loop through qr
		for _, qrMerchant := range qr {

			// genMerchantName := qrMerchant.MerchantName
			// splitedMerchantName := strings.Split(qrMerchant.MerchantName, " ")
			// if len(splitedMerchantName) > 1 {
			// 	firstName := splitedMerchantName[0]
			// 	lastName := splitedMerchantName[len(splitedMerchantName)-1]
			// 	// get the first letter of last name
			// 	lastNameFirstLetter := lastName[0:1]
			// 	// concat first name and last name first letter
			// 	genMerchantName = lastNameFirstLetter + "." + firstName
			// }

			// trim space
			qrMerchantName := strings.TrimSpace(qrMerchant.MerchantName)
			pwMerchantName := strings.TrimSpace(pwMerchant.MerchantName)

			// compare merchant name
			if pwMerchantName == qrMerchantName {
				// check if merchant name is existed
				if slices.Contains(existedName, qrMerchantName) {
					continue
				}

				// append merchant name to existedName
				existedName = append(existedName, qrMerchantName)

				// create a merchant similarity result
				merchantSimilarityResult := MerchantSimilarityResult{
					QrMarchantName:  qrMerchantName,
					PwMarchantName:  pwMerchantName,
					SimilarityScore: 1,
				}

				// append to result
				result = append(result, merchantSimilarityResult)
			} else {
				similarScore := strutil.Similarity(pwMerchantName, qrMerchantName, metrics.NewJaroWinkler())
				if similarScore > 0.9 {
					// check if merchant name is existed
					if slices.Contains(existedName, qrMerchantName) {
						continue
					}

					// append merchant name to existedName
					existedName = append(existedName, qrMerchantName)

					// create a merchant similarity result
					merchantSimilarityResult := MerchantSimilarityResult{
						QrMarchantName:  qrMerchantName,
						PwMarchantName:  pwMerchantName,
						SimilarityScore: similarScore,
					}

					// append to result
					result = append(result, merchantSimilarityResult)
				}
			}

			// update progress bar
			bar.Add(1)
		}
	}

	return result
}
