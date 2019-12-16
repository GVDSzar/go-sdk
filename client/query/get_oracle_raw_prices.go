package query

import (
	"encoding/json"
	"go-sdk/common"
)

func (c *client) GetOracleRawPrices(assetCode string) (*QueryResponse, error) {
	var obj = ResponseWithHeight()
	resp, code, err := c.baseClient.Get(GetOracleRawPricesEndpoint(assetCode), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, json.Unmarshal(resp, obj)
}
