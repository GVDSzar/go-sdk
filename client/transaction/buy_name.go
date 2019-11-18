package transaction

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) BuyNameTx(br *rest.BaseReq, name, amount, buyer string) (*tx.TxCommitResult, error) {
	r, err := msg.NewBuyNameRequest(br, name, amount, buyer)
	if err != nil {
		return nil, err
	}

	return c.RestTransaction(r, buyNameEndpoint, http.MethodPost)
}