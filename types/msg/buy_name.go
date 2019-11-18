package msg

import (
	"go-sdk/common/types"
)

type BuyName struct {
	Name  string           `json:"name"`
	Bid   types.CoinsWs    `json:"bid"`
	Buyer types.AccAddress `json:"buyer"`
}


// Route should return the name of the module
func (msg BuyName) Route() string { return "nameservice" }

// Type should return the action
func (msg BuyName) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg BuyName) ValidateBasic() error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg BuyName) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return b
}

// GetSigners defines whose signature is required
func (msg BuyName) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Buyer}
}

func (msg BuyName) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
