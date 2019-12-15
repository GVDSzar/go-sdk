package query

import "go-sdk/common"

func (c *client) GetDenominationsToken(token string) (*QueryResponse, error) {
	var obj = ResponseWithHeight()
	resp, code, err := c.baseClient.Get(GetDenominationsTokenEndpoint(token), emptyQP)
	obj.MustUnmarshal(resp)

	if err != nil {
		return obj, common.NewRestyHttpError(resp, code, err)
	}

	return obj, nil
}
