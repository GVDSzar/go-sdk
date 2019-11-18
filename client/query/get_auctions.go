package query

import (
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetAuctions() (*QueryResponse, error) {
	var aucts = WrapQueryResponse(types.Auctions{})
	resp, code, err := c.baseClient.Get(AuctionEndpoint, emptyQP)
	aucts.MustUnmarshal(resp)

	if err != nil {
		return aucts, common.NewRestyHttpError(resp, code, err)
	}

	return aucts, err
}
