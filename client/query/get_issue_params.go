package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetIssueParams() (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.Params{})
	resp, code, err := c.baseClient.Get(IssueParamsEndpoint, emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
