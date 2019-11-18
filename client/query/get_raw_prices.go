package query

import (
	"encoding/json"
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetRawPrices(assetCode string) (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.RawPrices{})
	resp, code, err := c.baseClient.Get(RawPricesEndpoint(assetCode), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, json.Unmarshal(resp, obj)
}
