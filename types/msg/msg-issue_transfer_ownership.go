package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

type IssueTransferOwnership struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	ToAddress   types.AccAddress `json:"to_address" yaml:"to_address"`
}

// Route Implements Msg.
func (msg IssueTransferOwnership) Route() string { return "issue" }

// Type Implements Msg.
func (msg IssueTransferOwnership) Type() string { return "issue_transfer_ownership" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg IssueTransferOwnership) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("IssueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg IssueTransferOwnership) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg IssueTransferOwnership) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
