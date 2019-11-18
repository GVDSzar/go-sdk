package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) DebtAuction(r *msg.StartDebtAuctionRequest, issueId, accAddress, amount string) (*tx.TxCommitResult, error) {
	return c.RestTransaction(r, liquidatorMintEndpoint, http.MethodPost)
}
