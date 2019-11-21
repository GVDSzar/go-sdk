package types

type CurrentPrice struct {
	AssetCode string  `json:"asset_code"`
	Price     string `json:"price"`
	Expiry    string `json:"expiry"`
}

type CurrentPriceWithHeight struct {
	Height int64          `json:"height"`
	Result []CurrentPrice `json:"result"`
}
