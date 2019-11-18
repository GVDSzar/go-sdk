package msg

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/valyala/fastjson"
	"go-sdk/common/types"
)

type BuyNameReq struct {
	BaseReq *rest.BaseReq    `json:"base_req"`
	Name    string           `json:"name"`
	Amount  string           `json:"amount"`
	Buyer   types.AccAddress `json:"buyer"`
}

func NewBuyNameRequest(br *rest.BaseReq, name, amount, buyerAddr string) (*BuyNameReq, error) {
	addr, err := types.AccAddressFromBech32(buyerAddr)
	if err != nil {
		return nil, err
	}
	return &BuyNameReq{br, name, amount, addr}, nil
}

//uses fastjson parser to speed up computation
func (m *BuyNameReq) messageFromStdTx(b []byte) ([]byte, error) {
	var p fastjson.Parser
	v, err := p.ParseBytes(b)
	if err != nil {
		return nil, err
	}
	msgJson := v.GetObject("value", "msg", "0", "value").String()
	return []byte(msgJson), err
}

func (m *BuyNameReq) validateResponse(bnr *BuyName) error {
	if bnr.Name != m.Name {
		errmsg := fmt.Sprintf("expected msg name: %s, got: %s", m.Name, bnr.Name)
		return MsgContentError(errmsg)
	}

	if !bytes.Equal(bnr.Buyer, m.Buyer) {
		errmsg := fmt.Sprintf("expected msg buyer: %s, got: %s", string(m.Buyer), string(bnr.Buyer))
		return MsgContentError(errmsg)
	}

	if len(bnr.Bid) != 1 {
		errmsg := fmt.Sprintf("expected bid array length: %v, got: %v", 1, len(bnr.Bid))
		return MsgContentError(errmsg)
	}

	bidAmount := bnr.Bid[0].Amount + bnr.Bid[0].Denom
	if bidAmount != m.Amount {
		errmsg := fmt.Sprintf("expected coin data: %s, got: %s", m.Amount, bidAmount)
		return MsgContentError(errmsg)
	}

	return nil
}

func (m *BuyNameReq) ValidateMsg(res Msg) error {
	var bnr = res.(BuyName)

	err := m.validateResponse(&bnr)
	if err != nil {
		return err
	}

	return err
}

func (msg BuyNameReq) Route() string { return "nameservice" }

func (msg BuyNameReq) Type() string { return "buy_name" }

func (msg BuyNameReq) ValidateBasic() error {
	return nil
}

func (msg BuyNameReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg BuyNameReq) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Buyer}
}

func (msg BuyNameReq) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
