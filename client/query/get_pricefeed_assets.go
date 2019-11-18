package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetPricefeedAssets() (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.PricefeedAssets{})
	resp, code, err := c.baseClient.Get(PricefeedAssetsEndpoint, emptyQP)
	obj.MustUnmarshal(resp)
	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
