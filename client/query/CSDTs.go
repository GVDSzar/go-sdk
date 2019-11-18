package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) CSDTs() (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.CSDTs{})
	resp, code, err := c.baseClient.Get(CSDTsEndpoint, emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}

func (c *client) CSDTParams() (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.CsdtModuleParams{})
	resp, code, err := c.baseClient.Get(CSDTParamsEndpoint, emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
