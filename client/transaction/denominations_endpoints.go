package transaction

const (
	denominationsRoute = "/denominations"

	denominationsIssueToken       = denominationsRoute + "/tokens"
	denominationsMintEndpoint     = denominationsRoute + "/tokens/mint"
	denominationsBurnEndpoint     = denominationsRoute + "/tokens/burn"
	denominationsFreezeEndpoint   = denominationsRoute + "/tokens/freeze"
	denominationsUnfreezeEndpoint = denominationsRoute + "/tokens/unfreeze"
)
