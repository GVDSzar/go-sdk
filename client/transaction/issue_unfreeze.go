package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) IssueUnfreeze(r *msg.PostIssueReq, freezeType, issueId, accAddress string) (*tx.TxCommitResult, error) {

	return c.RestTransaction(r, issueUnfreeze(freezeType, issueId, accAddress), http.MethodPost)
}
