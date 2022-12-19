package types

type Merchant struct {
	MerchantName string `json:"merchant_name"`
}

type MerchantSimilarityResult struct {
	QrMarchantName  string  `json:"qr_marchant_name"`
	PwMarchantName  string  `json:"pw_marchant_name"`
	SimilarityScore float64 `json:"similarity_score"`
}
