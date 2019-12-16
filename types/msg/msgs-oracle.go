package msg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go-sdk/common/types"
)

const oracleRouterKey = "oracle"

type MsgPostPrice struct {
	From      types.AccAddress // client that sent in this address
	AssetCode string           // asset code used by exchanges/api
	Price     types.Dec        // price in decimal (max precision 18)
	Expiry    types.Int        // block height
}

// Route Implements Msg.
func (msg MsgPostPrice) Route() string { return oracleRouterKey }

// Type Implements Msg
func (msg MsgPostPrice) Type() string { return "post_price" }

// GetSignBytes Implements Msg.
func (msg MsgPostPrice) GetSignBytes() []byte {
	return Cdc.MustMarshalJSON(msg)
}

// GetSigners Implements Msg.
func (msg MsgPostPrice) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.From}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgPostPrice) ValidateBasic() error {
	return nil
}

func (msg MsgPostPrice) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.From}
}

type MsgAddOracle struct {
	Oracle  types.AccAddress `json:"oracle" yaml:"oracle"`
	Nominee types.AccAddress `json:"nominee" yaml:"nominee"`
	Denom   string           `json:"denom" yaml:"denom"`
}

// MsgAddOracle creates a new add oracle message
func NewMsgAddOracle(
	nominee types.AccAddress,
	denom string,
	oracle types.AccAddress,
) MsgAddOracle {
	return MsgAddOracle{
		Oracle:  oracle,
		Denom:   denom,
		Nominee: nominee,
	}
}

// Route Implements Msg.
func (msg MsgAddOracle) Route() string { return oracleRouterKey }

// Type Implements Msg
func (msg MsgAddOracle) Type() string { return "add_oracle" }

// GetSignBytes Implements Msg.
func (msg MsgAddOracle) GetSignBytes() []byte {
	bz := Cdc.MustMarshalJSON(msg)
	return bz
}

// GetSigners Implements Msg.
func (msg MsgAddOracle) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgAddOracle) ValidateBasic() error {
	if msg.Oracle == nil {
		return sdk.ErrInvalidAddress("missing oracle address")
	}

	if msg.Denom == "" {
		return sdk.ErrInvalidCoins("missing denom")
	}

	if msg.Nominee == nil {
		return sdk.ErrInvalidAddress("missing nominee address")
	}
	return nil
}

func (msg MsgAddOracle) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}

type Oracle struct {
	Address sdk.AccAddress `json:"address" yaml:"address"`
}

type Oracles []Oracle

// MsgSetOracle struct representing a new nominee based oracle
type MsgSetOracles struct {
	Oracles Oracles          `json:"oracles" yaml:"oracles"`
	Nominee types.AccAddress `json:"nominee" yaml:"nominee"`
	Denom   string           `json:"denom" yaml:"denom"`
}

// MsgAddOracle creates a new add oracle message
func NewMsgSetOracles(
	nominee types.AccAddress,
	denom string,
	oracles Oracles,
) MsgSetOracles {
	return MsgSetOracles{
		Oracles: oracles,
		Denom:   denom,
		Nominee: nominee,
	}
}

// Route Implements Msg.
func (msg MsgSetOracles) Route() string { return oracleRouterKey }

// Type Implements Msg
func (msg MsgSetOracles) Type() string { return "set_oracles" }

// GetSignBytes Implements Msg.
func (msg MsgSetOracles) GetSignBytes() []byte {
	bz := Cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgSetOracles) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgSetOracles) ValidateBasic() error {
	if len(msg.Oracles) == 0 {
		return sdk.ErrInvalidAddress("missing oracle address")
	}

	if msg.Denom == "" {
		return sdk.ErrInvalidCoins("missing denom")
	}

	if msg.Nominee == nil {
		return sdk.ErrInvalidAddress("missing nominee address")
	}
	return nil
}

func (msg MsgSetOracles) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}

type Asset struct {
	AssetCode  string  `json:"asset_code" yaml:"asset_code"`
	BaseAsset  string  `json:"base_asset" yaml:"base_asset"`
	QuoteAsset string  `json:"quote_asset" yaml:"quote_asset"`
	Oracles    Oracles `json:"oracles" yaml:"oracles"`
	Active     bool    `json:"active" yaml:"active"`
}

// MsgSetAsset struct representing a new nominee based oracle
type MsgSetAsset struct {
	Nominee types.AccAddress `json:"nominee" yaml:"nominee"`
	Denom   string           `json:"denom" yaml:"denom"`
	Asset   Asset            `json:"asset" yaml:"asset"`
}

// NewMsgSetAsset creates a new add oracle message
func NewMsgSetAsset(
	nominee types.AccAddress,
	denom string,
	asset Asset,
) MsgSetAsset {
	return MsgSetAsset{
		Asset:   asset,
		Denom:   denom,
		Nominee: nominee,
	}
}

// Route Implements Msg.
func (msg MsgSetAsset) Route() string { return oracleRouterKey }

// Type Implements Msg
func (msg MsgSetAsset) Type() string { return "set_asset" }

// GetSignBytes Implements Msg.
func (msg MsgSetAsset) GetSignBytes() []byte {
	bz := Cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgSetAsset) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgSetAsset) ValidateBasic() error {
	return nil
}

func (msg MsgSetAsset) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}

// MsgAddAsset struct representing a new nominee based oracle
type MsgAddAsset struct {
	Nominee types.AccAddress `json:"nominee" yaml:"nominee"`
	Denom   string           `json:"denom" yaml:"denom"`
	Asset   Asset            `json:"asset" yaml:"asset"`
}

// NewMsgAddAsset creates a new add oracle message
func NewMsgAddAsset(
	nominee types.AccAddress,
	denom string,
	asset Asset,
) MsgAddAsset {
	return MsgAddAsset{
		Asset:   asset,
		Denom:   denom,
		Nominee: nominee,
	}
}

// Route Implements Msg.
func (msg MsgAddAsset) Route() string { return oracleRouterKey }

// Type Implements Msg
func (msg MsgAddAsset) Type() string { return "add_asset" }

// GetSignBytes Implements Msg.
func (msg MsgAddAsset) GetSignBytes() []byte {
	bz := Cdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgAddAsset) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgAddAsset) ValidateBasic() error {
	return nil
}

func (msg MsgAddAsset) GetInvolvedAddresses() []types.AccAddress {
	return []types.AccAddress{msg.Nominee}
}
