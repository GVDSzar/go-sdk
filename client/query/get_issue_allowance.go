package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetIssueAllowance(issueId, accAddr, spenderAddr string) (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.IssueFreeze{})
	resp, code, err := c.baseClient.Get(IssueAlowanceEndpoint(issueId, accAddr, spenderAddr), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
