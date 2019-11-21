package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type MsgIssueFreeze struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
	FreezeType  string           `json:"freeze_type" yaml:"freeze_type"`
}

const (
	FreezeIn       = "in"
	FreezeOut      = "out"
	FreezeInAndOut = "in-out"
)

var FreezeTypes = map[string]int{FreezeIn: 1, FreezeOut: 1, FreezeInAndOut: 1}

// Route Implements Msg.
func (msg MsgIssueFreeze) Route() string { return issueRouterKey }

// Type Implements Msg.
func (msg MsgIssueFreeze) Type() string { return "issue_freeze" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueFreeze) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("issueId cannot be empty")
	}
	_, ok := FreezeTypes[msg.FreezeType]
	if !ok {
		return errors.New("unknown freeze type")
	}
	return nil
}

func (msg MsgIssueFreeze) ValidateService() error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueFreeze) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueFreeze) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}

func (msg MsgIssueFreeze) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
