package transaction

import (
	"github.com/pkg/errors"
	"go-sdk/client/query"
	"go-sdk/common/types"
	"net/http"
)

func (c *client) GetAccountInfo() (*types.AccountInfo, error) {
	var obj = query.ResponseWithHeight()
	addr := c.keyManager.GetAddr().String()
	if len(addr) == 0 {
		return nil, errors.New("client address not set")
	}

	res, err := c.basicClient.FastHttpRequest(accountInfoEndpoint(addr), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	obj.MustUnmarshal(res)
	accNum := obj.Json().GetInt64("result", "value", "account_number")
	seq := obj.Json().GetInt64("result", "value", "sequence")
	var accInfo = &types.AccountInfo{
		AccountNumber: accNum,
		Sequence:      seq,
	}

	return accInfo, err
}
