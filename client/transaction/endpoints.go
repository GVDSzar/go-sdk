package transaction

import "fmt"

const (
	buyNameEndpoint         = "/nameservice/names" //test endpoint for a local chain
	csdtParamsEndpoint      = "/csdts"
	postIssueEndpoint       = "/issue"
	liquidatorSeizeEndpoint = "/liquidator/seize"
	liquidatorMintEndpoint  = "/liquidator/mint"
)

func accountInfoEndpoint(accAddr string) string {
	return fmt.Sprintf("/auth/accounts/%s", accAddr)
}

func placeBidEndpoint(auctionId, bidder, bid, lot string) string {
	return fmt.Sprintf("/auction/bid/%s/%s/%s/%s", auctionId, bidder, bid, lot)
}