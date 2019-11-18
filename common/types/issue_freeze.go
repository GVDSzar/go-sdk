package types

type IssueFreeze struct {
	Frozen bool `json:"frozen"`
}

type IssueFreezeWithHeight struct {
	Height int64         `json:"height"`
	Result []IssueFreeze `json:"result"`
}
