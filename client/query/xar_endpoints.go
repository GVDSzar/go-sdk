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
	GetDenominationSymbols  = "/denominations/tokens"
	GetNftDenomsEndpoint    = "/nft/denoms"
	//"/issue/query/" // + {issue id}
)

func GetNftSupplyEndpoint(denom string) string {
	return fmt.Sprintf("/nft/supply/%s", denom)
}

func GetNftOwnerByDenomEndpoint(delegatorAddr, denom string) string {
	return fmt.Sprintf("/nft/owner/%s/collection/%s", delegatorAddr, denom)
}

func GetNftsEndpoint(denom , id string) string {
	return fmt.Sprintf("/nft/collection/%s/nft/%s", denom, id)
}

func GetNftCollectionEndpoint(denom string) string {
	return fmt.Sprintf("/nft/collection/%s", denom)
}

func GetNftOwnerEndpoint(delegatorAddr string) string {
	return fmt.Sprintf("/nft/owner/%s", delegatorAddr)
}

func GetDenominationsTokenEndpoint(token string) string {
	return fmt.Sprintf("/denominations/tokens/%s", token)
}

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

func GetOracleRawPricesEndpoint(restName string) string {
	return fmt.Sprintf("/oracle/rawprices/%s", restName)
}

func GetOracleCurrentPricesEndpoint(restName string) string {
	return fmt.Sprintf("/oracle/currentprice/%s", restName)
}

func GetAssets() string {
	return fmt.Sprintf("/oracle/assets",)
}

func SwapEstimateEndpoint(sender, asset, targetDenom string) string {
	return fmt.Sprintf("/swap/estimate/%s/%s/%s", sender, asset, targetDenom)
}
