package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) NftsTransfer(r *msg.TransferNFTReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, nftsTransferEndpoint, http.MethodPost)
}
