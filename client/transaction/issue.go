package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

//NewPostIssueReq

func (c *client) IssueTx(r *msg.PostIssueReq) (*tx.TxCommitResult, error) {

	return c.RestTransaction(r, postIssueEndpoint, http.MethodPost)
}