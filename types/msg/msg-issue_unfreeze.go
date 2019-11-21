package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type MsgIssueUnFreeze struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	FreezeType  string           `json:"freeze_type" yaml:"freeze_type"`
}

// Route Implements Msg.
func (msg MsgIssueUnFreeze) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueUnFreeze) Type() string { return "issue_unfreeze" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueUnFreeze) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("issueId cannot be empty")
	}
	_, ok := FreezeTypes[msg.FreezeType]
	if !ok {
		return errors.New("unknown freeze type")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueUnFreeze) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueUnFreeze) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

func (msg MsgIssueUnFreeze) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
