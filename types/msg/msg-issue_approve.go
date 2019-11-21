package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type MsgIssueApprove struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	Amount      string        `json:"amount" yaml:"amount"`
}

// Route Implements Msg.
func (msg MsgIssueApprove) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueApprove) Type() string { return "issue_approve" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueApprove) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("issueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	//if msg.Amount.I.Int64() < 0 {
	//	return errors.New("can't approve negative coin amount")
	//}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueApprove) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueApprove) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

func (msg MsgIssueApprove) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
