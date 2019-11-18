package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
	"go-sdk/types/msg"
)

func (c *client) GetCurrentPrice(assetCode string) (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.CurrentPrice{})
	resp, code, err := c.baseClient.Get(CurrentPricesEndpoint(assetCode), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, msg.Cdc.UnmarshalJSON(resp, obj)
}
