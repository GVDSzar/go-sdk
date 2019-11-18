package msg

import (
	"github.com/pkg/errors"
	"go-sdk/common/types"
)

const CoinDescriptionMaxLength = 1024

type IssueDescription struct {
	IssueId     string           `json:"issue_id" yaml:"issue_id"`
	FromAddress types.AccAddress `json:"from_address" yaml:"from_address"`
	Description []byte           `json:"description" yaml:"description"`
}

func (msg IssueDescription) Route() string { return "issue" }

// Type Implements Msg.
func (msg IssueDescription) Type() string { return "issue_description" }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg IssueDescription) ValidateBasic() error {
	if len(msg.IssueId) == 0 {
		return errors.New("issueId cannot be empty")
	}
	if len(msg.Description) > CoinDescriptionMaxLength {
		return errors.New("msg length is incorrect")
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg IssueDescription) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg IssueDescription) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.FromAddress}
}
