package msg

import "github.com/cosmos/cosmos-sdk/types"

type MsgIssueSendFrom struct {
	IssueId     string         `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	From        types.AccAddress `json:"from" yaml:"from"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Int        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueSendFrom) Route() string { return "issue" }

// Type Implements Msg.
func (msg MsgIssueSendFrom) Type() string { return "issue_send_from" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueSendFrom) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return types.ErrInvalidAddress("issueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.Amount.IsNegative() {
		return types.ErrInvalidAddress("can't send negative amount")
	}
	if msg.From.Equals(msg.ToAddress) {
		return types.ErrInvalidAddress("can't send yourself")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueSendFrom) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueSendFrom) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
