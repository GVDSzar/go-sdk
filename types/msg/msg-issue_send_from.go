package msg

import "github.com/cosmos/cosmos-sdk/types"

type IssueSendFrom struct {
	IssueId     string         `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	From        types.AccAddress `json:"from" yaml:"from"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      types.Int        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg IssueSendFrom) Route() string { return "issue" }

// Type Implements Msg.
func (msg IssueSendFrom) Type() string { return "issue_send_from" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg IssueSendFrom) ValidateBasic() error {
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
func (msg IssueSendFrom) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg IssueSendFrom) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
