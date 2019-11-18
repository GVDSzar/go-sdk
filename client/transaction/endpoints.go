package transaction

import "fmt"

const (
	buyNameEndpoint         = "/nameservice/names" //test endpoint for a local chain
	csdtParamsEndpoint      = "/csdts"
	postIssueEndpoint       = "/issue"
	liquidatorSeizeEndpoint = "/liquidator/seize"
	liquidatorMintEndpoint  = "/liquidator/mint"
)

func placeBidEndpoint(auctionId, bidder, bid, lot string) string {
	return fmt.Sprintf("/auction/bid/%s/%s/%s/%s", auctionId, bidder, bid, lot)
}

func issueApprove(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/issue/approve/%s/%s/%s", issueId, accAddress, amount)
}

func issueIncreaseApproval(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/issue/approve/increase/%s/%s/%s", issueId, accAddress, amount)
}

func issueDecreaseApproval(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/issue/approve/decrease/%s/%s/%s", issueId, accAddress, amount)
}

func issueBurn(issueId, amount string) string {
	return fmt.Sprintf("/issue/burn/%s/%s", issueId, amount)
}

func issueBurnFrom(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/issue/burn-from/{%s}/{%s}/{%s}", issueId, accAddress, amount)
}

func issueFreeze(freezeType, issueId, accAddress string) string {
	return fmt.Sprintf("/issue/freeze/%s/%s/%s", freezeType, issueId, accAddress)
}

func issueUnfreeze(freezeType, issueId, accAddress string) string {
	return fmt.Sprintf("/issue/unfreeze/%s/%s/%s", freezeType, issueId, accAddress)
}

func issueSendFrom(issueId, from, to, amount string) string {
	return fmt.Sprintf("/issue/send-from/%s/%s/%s/%s", issueId, from, to, amount)
}

func issueMint(issueId, amount, to string) string {
	return fmt.Sprintf("/issue/mint/%s/%s/%s", issueId, amount, to)
}

func issueTransfer(issueId, to string) string {
	return fmt.Sprintf("/issue/ownership/transfer/%s/%s", issueId, to)
}

func issueDescription(issueId string) string {
	return fmt.Sprintf("/issue/description/%s", issueId)
}

func issueDisableFeature(issueId, feature string) string {
	return fmt.Sprintf("/issue/feature/disable/%s/%s", issueId, feature)
}
