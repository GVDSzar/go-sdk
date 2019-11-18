package xar_api_test

import (
	"encoding/json"
	"go-sdk/client"
	"go-sdk/common/types"
	"log"
	"testing"
)

// ############################################################
// ## this test file is only suitable for local rest testing ##
// ##          change constants to make it work              ##
// ############################################################

// it is assumed that std local rest-daemon is up
const baseUrl = "localhost:1317"

func TestQueries(t *testing.T) {
	c, err := client.NewCustomClient(baseUrl, types.TestNetwork, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	testGetAuctions(t, c)
	testCsdts(t, c)
	testCsdtParams(t, c)
	testGetInterest(t, c)
	testIssues(t, c)
	testGetRawPrices(t, c)
	testGetCurrentPrice(t, c)
	testGetPricefeedAssets(t, c)
	testSwapEstimate(t, c)
}

func testIssues(t *testing.T, c client.DexClient) {
	testIssueParams(t, c)
	testIssue(t, c)
	testIssueList(t, c)
	testIssueInfo(t, c)
	testIssueFreeze(t, c)
	testIssueFreezes(t, c)
	testIssueAllowance(t, c)
}

func testGetAuctions(t *testing.T, c client.DexClient) {
	aucts, err := c.GetAuctions()
	if err != nil {
		t.Errorf(err.Error())
	}
	res, err := json.Marshal(aucts)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("auctions res: ", string(res))
}
func testGetPricefeedAssets(t *testing.T, c client.DexClient) {
	pricesfeed, err := c.GetPricefeedAssets()
	if err != nil {
		t.Errorf(err.Error())
	}
	res, _ := json.Marshal(pricesfeed)
	log.Println("pricesfeed res: ", string(res))
}
func testSwapEstimate(t *testing.T, c client.DexClient) {
	swap, err := c.SwapEstimate()
	if err != nil {
		t.Errorf(err.Error())
	}
	res, _ := json.Marshal(swap)
	log.Println("swap res: ", string(res))
}

func testCsdts(t *testing.T, c client.DexClient) {
	csdt, err := c.CSDTs()
	if err != nil {
		t.Errorf(err.Error())
	}
	res, err := json.Marshal(csdt)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("csdt res: ", string(res))
}

func testGetRawPrices(t *testing.T, c client.DexClient) {
	rawPrices, err := c.GetRawPrices("assetCode")
	if err != nil {
		t.Errorf(err.Error())
	}
	res, err := json.Marshal(rawPrices)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("rawPrices res: ", string(res))
}

func testGetCurrentPrice(t *testing.T, c client.DexClient) {
	currprice, err := c.GetCurrentPrice("assetCode")
	if err != nil {
		t.Errorf(err.Error())
	}
	res, _ := json.Marshal(currprice)
	log.Println("currprice res: ", string(res))
}

func testCsdtParams(t *testing.T, c client.DexClient) {
	params, err := c.CSDTParams()
	if err != nil {
		t.Errorf(err.Error())
	}
	res, err := json.Marshal(params)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("params res: ", string(res))
}

func testGetInterest(t *testing.T, c client.DexClient) {
	interest, err := c.GetInterest()
	if err != nil {
		t.Errorf(err.Error())
	}
	res, err := json.Marshal(interest)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("interest res: ", string(res))
}

func testIssueParams(t *testing.T, c client.DexClient) {
	issuepar, err := c.GetIssueParams()
	if err != nil {
		t.Errorf(err.Error())
	}

	res, err := json.Marshal(issuepar)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("issuepar res: ", string(res))
}

func testIssueList(t *testing.T, c client.DexClient) {
	issuelist, err := c.GetIssueList()
	if err != nil {
		t.Errorf(err.Error())
	}

	res, err := json.Marshal(issuelist)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("issuelist res: ", string(res))
}

func testIssueInfo(t *testing.T, c client.DexClient) {
	symbs := []string{
		"s",
	}

	for _, v := range symbs {
		issueinfo, err := c.SearchIssueWithSymbol(v)
		if err != nil {
			t.Errorf(err.Error())
		}

		res, err := json.Marshal(issueinfo)
		if err != nil {
			t.Errorf(err.Error())
		}
		log.Println("issue res: ", string(res))
	}
}

func testIssueFreeze(t *testing.T, c client.DexClient) {
	issueFreeze := [][]string{
		{"issueid", "addr"},
	}

	for _, v := range issueFreeze {
		freeze, err := c.GetIssueFreeze(v[0], v[1])
		if err != nil {
			t.Errorf(err.Error())
		}
		res, err := json.Marshal(freeze)
		if err != nil {
			t.Errorf(err.Error())
		}
		log.Println("freeze res: ", string(res))
	}
}

func testIssueFreezes(t *testing.T, c client.DexClient) {
	issueFreeze := []string{
		"freezes",
	}

	for _, v := range issueFreeze {
		freezes, err := c.GetIssueFreezes(v)
		if err != nil {
			t.Errorf(err.Error())
		}
		res, err := json.Marshal(freezes)
		if err != nil {
			t.Errorf(err.Error())
		}
		log.Println("freezes res: ", string(res))
	}
}

func testIssueAllowance(t *testing.T, c client.DexClient) {
	issueFreeze := [][]string{
		{"issueid", "accaddr", "spenderaddr"},
	}

	for _, v := range issueFreeze {
		freeze, err := c.GetIssueAllowance(v[0], v[1], v[2])
		if err != nil {
			t.Errorf(err.Error())
		}
		res, err := json.Marshal(freeze)
		if err != nil {
			t.Errorf(err.Error())
		}
		log.Println("GetIssueAllowance res: ", string(res))
	}
}

func testIssue(t *testing.T, c client.DexClient) {
	errorExpected := []string{
		"issue",
	}
	for _, v := range errorExpected {
		issueid, err := c.GetIssueId(v)
		if err != nil {
			t.Errorf(err.Error())
		}

		j, err := json.Marshal(issueid)
		if err != nil {
			t.Errorf(err.Error())
			continue
		}
		log.Println("issue id res: ", string(j))
	}
}
