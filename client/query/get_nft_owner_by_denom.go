package query

import "go-sdk/common"

func (c *client) GetNftOwnerByDenom(delegatorAddr, denom string) (*QueryResponse, error) {
	var obj = ResponseWithHeight()
	resp, code, err := c.baseClient.Get(GetNftOwnerByDenomEndpoint(delegatorAddr, denom), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}