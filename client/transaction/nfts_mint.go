package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) NftsMint(r *msg.MintNFTReq) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, nftsMintEndpoint, http.MethodPost)
}
