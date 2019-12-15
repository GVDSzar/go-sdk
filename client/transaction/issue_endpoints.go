package transaction

import "fmt"

const (
	issueRoute = "issue"
)


func issueApprove(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/%s/approve/%s/%s/%s", issueRoute, issueId, accAddress, amount)
}

func issueIncreaseApproval(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/%s/approve/increase/%s/%s/%s", issueRoute, issueId, accAddress, amount)
}

func issueDecreaseApproval(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/%s/approve/decrease/%s/%s/%s", issueRoute, issueId, accAddress, amount)
}

func issueBurn(issueId, amount string) string {
	return fmt.Sprintf("/%s/burn/%s/%s", issueRoute, issueId, amount)
}

func issueBurnFrom(issueId, accAddress, amount string) string {
	return fmt.Sprintf("/%s/burn-from/{%s}/{%s}/{%s}", issueRoute, issueId, accAddress, amount)
}

func issueFreeze(freezeType, issueId, accAddress string) string {
	return fmt.Sprintf("/%s/freeze/%s/%s/%s", issueRoute, freezeType, issueId, accAddress)
}

func issueUnfreeze(freezeType, issueId, accAddress string) string {
	return fmt.Sprintf("/%s/unfreeze/%s/%s/%s", issueRoute, freezeType, issueId, accAddress)
}

func issueSendFrom(issueId, from, to, amount string) string {
	return fmt.Sprintf("/%s/send-from/%s/%s/%s/%s", issueRoute, issueId, from, to, amount)
}

func issueMint(issueId, amount, to string) string {
	return fmt.Sprintf("/%s/mint/%s/%s/%s", issueRoute, issueId, amount, to)
}

func issueTransfer(issueId, to string) string {
	return fmt.Sprintf("/%s/ownership/transfer/%s/%s", issueRoute, issueId, to)
}

func issueDescription(issueId string) string {
	return fmt.Sprintf("/%s/description/%s", issueRoute, issueId)
}

func issueDisableFeature(issueId, feature string) string {
	return fmt.Sprintf("/%s/feature/disable/%s/%s", issueRoute, issueId, feature)
}