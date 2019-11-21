package msg

import (
	"go-sdk/common/types"
)

type RecordInfo struct {
	ID          string           `json:"id"`
	Hash        string           `json:"hash"`
	RecordNo    string           `json:"record_number"`
	Sender      types.AccAddress `json:"sender"`
	RecordTime  int64            `json:"record_time"`
	Name        string           `json:"name"`
	Author      string           `json:"author"`
	RecordType  string           `json:"record_type"`
	Description string           `json:"description"`
}
