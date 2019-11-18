package query

import (
	"fmt"
)

const (
	AuctionEndpoint         = "/auction/getauctions"
	CSDTsEndpoint           = "/auction/getauctions"
	CSDTParamsEndpoint      = "/csdts/params"
	InterestEndpoint        = "/interest/current"
	IssueParamsEndpoint     = "/issue/params"
	IssueListEndpoint       = "/issue/list"
	OutstandingDebtEndpoint = "/liquidator/outstandingdebt"
	PricefeedAssetsEndpoint = "/assets"
	//"/issue/query/" // + {issue id}
)

func IssueIdEndpoint(issueId string) string {
	return fmt.Sprintf("/issue/query/%s", issueId)
}

func IssueSearchEndpoint(symbol string) string {
	return fmt.Sprintf("/issue/search/%s", symbol)
}

func IssueFreezeEndpoint(issueID, accAddress string) string {
	return fmt.Sprintf("/issue/freeze/%s/%s", issueID, accAddress)
}

func IssueFreezesEndpoint(issueID string) string {
	return fmt.Sprintf("/issue/freezes/%s", issueID)
}

func IssueAlowanceEndpoint(issueID, address, spenderAddress string) string {
	return fmt.Sprintf("/issue/allowance/%s/%s/%s", issueID, address, spenderAddress)
}

func RawPricesEndpoint(assetCode string) string {
	return fmt.Sprintf("/pricefeed/rawprices/%s", assetCode)
}

func CurrentPricesEndpoint(assetCode string) string {
	return fmt.Sprintf("/pricefeed/currentprice/%s", assetCode)
}

func SwapEstimateEndpoint(sender, asset, targetDenom string) string {
	return fmt.Sprintf("/swap/estimate/%s/%s/%s", sender, asset, targetDenom)
}
