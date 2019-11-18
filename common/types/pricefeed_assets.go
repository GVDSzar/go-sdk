package types

type PricefeedAssets []string

type PricefeedAssetsWithHeight struct {
	Height int64             `json:"height"`
	Result PricefeedAssets `json:"result"`
}
