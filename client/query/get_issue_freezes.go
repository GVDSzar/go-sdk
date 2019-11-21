package query

import (
	"go-sdk/common"
)

func (c *client) GetIssueFreeze(issueId, accAddress string) (*QueryResponse, error) {
	var obj = ResponseWithHeight()
	resp, code, err := c.baseClient.Get(IssueFreezeEndpoint(issueId, accAddress), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
