package transaction

import (
	"go-sdk/types/msg"
	"go-sdk/types/tx"
	"net/http"
)

func (c *client) NftsEditMetadata(r *msg.EditNFTMetadataReq, denom, id string) (*tx.TxBroadcastResult, error) {
	return c.RestTransaction(r, EditNftsMetadataEndpont(denom, id), http.MethodPut)
}
