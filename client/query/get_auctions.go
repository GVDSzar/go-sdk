package query

import (
	"go-sdk/common"
)

func (c *client) GetAuctions() (*QueryResponse, error) {
	var aucts = ResponseWithHeight()
	resp, code, err := c.baseClient.Get(AuctionEndpoint, emptyQP)
	aucts.MustUnmarshal(resp)

	if err != nil {
		return aucts, common.NewRestyHttpError(resp, code, err)
	}

	return aucts, err
}
