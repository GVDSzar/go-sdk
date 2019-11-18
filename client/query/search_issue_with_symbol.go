package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) SearchIssueWithSymbol(symbol string) (*QueryResponse, error) {
	var obj = WrapQueryResponse(types.CoinIssueInfo{})
	resp, code, err := c.baseClient.Get(IssueSearchEndpoint(symbol), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
