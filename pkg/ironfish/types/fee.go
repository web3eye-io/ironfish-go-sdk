package types

const EstimateFeeRatesPath = "chain/estimateFeeRates"

type EstimateFeeRatesResponse struct {
	Slow    string `json:"slow"`
	Average string `json:"average"`
	Fast    string `json:"fast"`
}
