package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetIssueFreezes(issueId string) (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.IssueAddressFreezeListWithHeight{})
	resp, code, err := c.baseClient.Get(IssueFreezesEndpoint(issueId), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
