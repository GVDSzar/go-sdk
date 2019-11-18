package types

type RawPrices []string

type RawPricesWithHeight struct {
	Height int64       `json:"height"`
	Result []RawPrices `json:"result"`
}
