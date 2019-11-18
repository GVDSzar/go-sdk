package types

type IssueAddressFreeze struct {
	Address string `json:"address"`
}

type IssueAddressFreezeList []IssueAddressFreeze

type IssueAddressFreezeListWithHeight struct {
	Height int64                  `json:"height"`
	Result IssueAddressFreezeList `json:"result"`
}
