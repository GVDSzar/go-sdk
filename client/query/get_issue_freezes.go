package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetIssueFreeze(issueId, accAddress string) (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.IssueFreeze{})
	resp, code, err := c.baseClient.Get(IssueFreezeEndpoint(issueId, accAddress), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
