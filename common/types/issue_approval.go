package types

type Approval struct {
	Amount string `json:"amount"`
}

type IssueApprovalWithHeight struct {
	Height int64      `json:"height"`
	Result []Approval `json:"result"`
}
