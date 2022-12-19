package utils

import (
	. "manuth/types"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func CleanPwMerchant(merchants []Merchant) []Merchant {
	bar := progressbar.Default(
		int64(len(merchants)),
		"cleaning merchant",
	)
	result := make([]Merchant, 0)

	// clean merchant
	for _, merchant := range merchants {
		bar.Add(1)
		// allow only merchant name that not number and not inlcude "by"
		allow := strings.Contains(merchant.MerchantName, "by") == false && strings.ContainsAny(merchant.MerchantName, "0123456789") == false

		if allow {
			// trim space
			name := strings.TrimSpace(merchant.MerchantName)

			// append to result
			result = append(result, Merchant{
				MerchantName: name,
			})
		}
	}

	return result
}
