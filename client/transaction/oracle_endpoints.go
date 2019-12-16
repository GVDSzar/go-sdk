package transaction

import "fmt"

const oracleRoute = "/oracle"

func PostPriceEndpoint() string {
	return fmt.Sprintf(oracleRoute + "/rawprices")
}