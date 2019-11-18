package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) SwapEstimate() (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.MsgSwap{})
	resp, code, err := c.baseClient.Get(InterestEndpoint, emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}