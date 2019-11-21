package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

const (
	BurnOwner  = "burn-owner"
	BurnHolder = "burn-holder"
	BurnFrom   = "burn-from"
	Freeze     = "freeze"
	Minting    = "minting"
)

var Features = map[string]int{BurnOwner: 1, BurnHolder: 1, BurnFrom: 1, Freeze: 1, Minting: 1}

type MsgIssueDisableFeature struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	Feature     string           `json:"feature" yaml:"feature"`
}

// Route Implements Msg.
func (msg MsgIssueDisableFeature) Route() string { return "issue" }

// Type Implements Msg.
func (msg MsgIssueDisableFeature) Type() string { return "issue_disable_feature" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueDisableFeature) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("issueId cannot be empty")
	}
	_, ok := Features[msg.Feature]
	if !ok {
		return ErrUnknownFeatures
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueDisableFeature) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgIssueDisableFeature) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
