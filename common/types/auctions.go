package types

type Auctions []string

type AuctionsWithHeight struct {
	Height int64 `json:"height"`
	Result Auctions `json:"result"`
}