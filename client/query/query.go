package query

import (
	"go-sdk/client/basic"
	"go-sdk/common/types"
)

type QueryClient interface {
	GetClosedOrders(query *types.ClosedOrdersQuery) (*types.CloseOrders, error)
	GetDepth(query *types.DepthQuery) (*types.MarketDepth, error)
	GetKlines(query *types.KlineQuery) ([]types.Kline, error)
	GetMarkets(query *types.MarketsQuery) ([]types.TradingPair, error)
	GetOrder(orderID string) (*types.Order, error)
	GetOpenOrders(query *types.OpenOrdersQuery) (*types.OpenOrders, error)
	GetTicker24h(query *types.Ticker24hQuery) ([]types.Ticker24h, error)
	GetTrades(query *types.TradesQuery) (*types.Trades, error)
	GetAccount(string) (*types.BalanceAccount, error)
	GetTime() (*types.Time, error)
	GetTokens(query *types.TokensQuery) ([]types.Token, error)
	GetNodeInfo() (*types.ResultStatus, error)
	GetShortNodeInfo() (*types.ResultStatus, error)

	//temprorary endpoint for a testchain
	GetNames() (nl *types.NameList, err error)

	//xar query endpoinst
	GetAuctions() (*QueryResponse, error)
	CSDTs() (*QueryResponse, error)
	CSDTParams() (*QueryResponse, error)
	GetInterest() (*QueryResponse, error)
	GetIssueParams() (*QueryResponse, error)
	GetIssueId(issueId string) (*QueryResponse, error)
	GetIssueList() (*QueryResponse, error)
	SearchIssueWithSymbol(symbol string) (*QueryResponse, error)
	GetIssueFreeze(issueId, accAddress string) (*QueryResponse, error)
	GetIssueFreezes(issueId string) (*QueryResponse, error)
	GetIssueAllowance(issueId, accAddr, spenderAddr string) (*QueryResponse, error)
	GetRawPrices(assetCode string) (*QueryResponse, error)
	GetCurrentPrice(assetCode string) (*QueryResponse, error)
	GetPricefeedAssets() (*QueryResponse, error)
	SwapEstimate() (*QueryResponse, error)
}

type client struct {
	baseClient basic.BasicClient
}

func NewClient(c basic.BasicClient) QueryClient {
	return &client{baseClient: c}
}

var emptyQP = map[string]string{}
