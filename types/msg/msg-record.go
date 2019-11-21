package msg

import (
	"go-sdk/common/types"
)

const recordRouterKey = "record"

type MsgRecord struct {
	Sender        types.AccAddress `json:"sender"`
	*RecordParams `json:"params"`
}

type RecordParams struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Hash        string `json:"hash"`
	RecordNo    string `json:"record_number"`
	RecordType  string `json:"record_type"`
	Description string `json:"description"`
}

//New MsgRecord Instance
func NewMsgRecord(sender types.AccAddress, params *RecordParams) MsgRecord {
	return MsgRecord{sender, params}
}

// Route Implements Msg.
func (msg MsgRecord) Route() string { return recordRouterKey }

// Type Implements Msg.
func (msg MsgRecord) Type() string { return "record" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgRecord) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgRecord) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgRecord) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}

func (msg MsgRecord) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Sender}
}
