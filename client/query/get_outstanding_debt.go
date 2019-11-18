package query

import (
	"encoding/json"
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetOutstandingDebt() (*types.OutstandingDebtWithHeight, error) {
	var obj = new(types.OutstandingDebtWithHeight)
	resp, code, err := c.baseClient.Get(OutstandingDebtEndpoint, emptyQP)
	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, json.Unmarshal(resp, obj)
}