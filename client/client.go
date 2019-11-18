package client

import (
	conf "go-sdk/types"
	"gopkg.in/resty.v1"

	"go-sdk/client/basic"
	"go-sdk/client/query"
	"go-sdk/client/transaction"
	"go-sdk/client/websocket"
	"go-sdk/common/types"
	"go-sdk/keys"
)

// dexClient wrapper
type dexClient struct {
	query.QueryClient
	websocket.WSClient
	transaction.TransactionClient
	basic.BasicClient
}

// DexClient methods
type DexClient interface {
	basic.BasicClient
	query.QueryClient
	websocket.WSClient
	transaction.TransactionClient
}

func init() {
	resty.DefaultClient.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
}

func NewDexClient(baseUrl string, network types.ChainNetwork, keyManager keys.KeyManager) (DexClient, error) {
	types.Network = network
	c := basic.NewClient(baseUrl)
	w := websocket.NewClient(c)
	q := query.NewClient(c)
	n, err := q.GetNodeInfo()
	if err != nil {
		return nil, err
	}
	t := transaction.NewClient(n.NodeInfo.Network, keyManager, q, c)
	return &dexClient{BasicClient: c, QueryClient: q, TransactionClient: t, WSClient: w}, nil
}

func NewCustomClient(baseUrl string, network types.ChainNetwork, keyManager keys.KeyManager) (DexClient, error) {
	types.Network = network
	c := basic.NewCustomClient(baseUrl, conf.UnsafeApiSchema, conf.NoPrefix)
	w := websocket.NewClient(c)
	q := query.NewClient(c)
	n, err := q.GetShortNodeInfo()
	if err != nil {
		return nil, err
	}
	t := transaction.NewClient(n.NodeInfo.Network, keyManager, q, c)
	return &dexClient{BasicClient: c, QueryClient: q, TransactionClient: t, WSClient: w}, nil
}
