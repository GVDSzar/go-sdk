package query

import (
	"github.com/tendermint/go-amino"
	"go-sdk/common"
)

func (c *client) GetIssueId(issueId string) (*QueryResponse, error) {
	var obj = ResponseWithHeight()
	resp, code, requestErr := c.baseClient.Get(IssueIdEndpoint(issueId), emptyQP)
	obj.MustUnmarshal(resp)

	if requestErr != nil {
		return obj, common.NewRestyHttpError(resp, code, requestErr)
	}

	return obj, amino.UnmarshalJSON(resp, obj)
}
