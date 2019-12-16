package query

import (
	"encoding/json"
	"go-sdk/common"
)

func (c *client) GetSwapEstimate(sender, asset, targetDenom string) (*QueryResponse, error) {
	var obj = ResponseWithHeight()
	resp, code, err := c.baseClient.Get(SwapEstimateEndpoint(sender, asset, targetDenom), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, json.Unmarshal(resp, obj)
}
