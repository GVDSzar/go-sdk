package xar_api

// ############################################################
// ## this test file is only suitable for local rest testing ##
// ##          change constants to make it work              ##
// ############################################################

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/client"
	"go-sdk/common/types"
	"go-sdk/keys"
	"go-sdk/types/msg"
	"log"
	"testing"
)

// it is assumed that std local rest-daemon is up
const baseUrl = "localhost:1317"
const localUserMnemonic = "hospital year blade virus verify win sell hope sauce effort acquire alien cradle area draw cube assault leg medal direct album light cargo income"
const localUserAddr = "xar1mu4sgq98tx3mxweknzmqt95vp3zx07gmuqk6js"

func TestTransactions(t *testing.T) {
	km, err := keys.NewMnemonicKeyManager(localUserMnemonic)
	if err != nil {
		panic(err)
	}

	c, err := client.NewCustomClient(baseUrl, types.TestNetwork, km)
	if err != nil {
		t.Errorf(err.Error())
	}

	testPlaceBid(t, c)
	testModifyCsdtTx(t, c)
	testIssueTx(t, c)
	testIssueApproveTx(t, c)
	testIssueIncreaseApprovalTx(t, c)
	testIssueDecreaseApprovalTx(t, c)
	testIssueBurnTx(t, c)
	testIssueBurnFromTx(t, c)
	testIssueFreeze(t, c)
	testIssueUnfreeze(t, c)
	testIssueSendFrom(t, c)
	testIssueMint(t, c)
	testIssueTransfer(t, c)
	testIssueDisableFeature(t, c)
	testLiquidatorSieze(t, c)
	testDebtAuction(t, c)
	testPriceRequest(t, c)
}

var testingAccAddress = sdk.AccAddress([]byte(localUserAddr))
var br = rest.BaseReq{ChainID: "testing"}

func testPlaceBid(t *testing.T, c client.DexClient) {
	auctionId := "auction_id"
	bidder := "bidder"
	bid := "bid"
	lot := "lot"
	mpb := msg.NewPlaceBidReq(br, auctionId, bidder, bid, lot)
	tx, err := c.PlaceBidTx(mpb)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("PlaceBid tx", j)
}

func testModifyCsdtTx(t *testing.T, c client.DexClient) {
	collateralDenom := "test"
	collateralAmount := sdk.Int{}
	debt := sdk.Int{}
	csdt := types.NewCSDT(testingAccAddress, collateralDenom, collateralAmount, debt)
	r := msg.NewModifyCsdtReq(br, csdt)
	tx, err := c.ModifyCsdtTx(r)
	j1, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	log.Println(string(j1))
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("PlaceBid tx", string(j))
}

func testIssueTx(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)
	tx, err := c.IssueTx(r)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("Issue tx", string(j))
}

func testIssueApproveTx(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	accAddr := "someAdress"
	amount := "amount"

	tx, err := c.IssueApproveTx(r, issueId, accAddr, amount)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueApprove tx", string(j))
}

func testIssueIncreaseApprovalTx(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	accAddr := "someAdress"
	amount := "amount"
	tx, err := c.IssueIncreaseApprovalTx(r, issueId, accAddr, amount)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueIncreaseApproval tx", string(j))
}

func testIssueDecreaseApprovalTx(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	accAddr := "someAdress"
	amount := "amount"
	tx, err := c.IssueDecreaseApprovalTx(r, issueId, accAddr, amount)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueDecreaseApproval tx", string(j))
}

func testIssueBurnTx(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	amount := "amount"
	tx, err := c.IssueBurnTx(r, issueId, amount)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueBurn tx", string(j))
}

func testIssueBurnFromTx(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	accAddr := "someAdress"
	amount := "amount"
	tx, err := c.IssueBurnFromTx(r, issueId, accAddr, amount)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueBurnFrom tx", string(j))
}

func testIssueFreeze(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	accAddr := "someAdress"
	freezeType := "freezeType"
	tx, err := c.IssueFreeze(r, freezeType, issueId, accAddr)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueFreeze tx", string(j))
}

func testIssueUnfreeze(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	accAddr := "someAdress"
	freezeType := "freezeType"
	tx, err := c.IssueUnfreeze(r, freezeType, issueId, accAddr)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueUnfreeze tx", string(j))
}

func testIssueSendFrom(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	from := "from"
	to := "to"
	amount := "amount"
	tx, err := c.IssueSendFrom(r, issueId, from, to, amount)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueSendFrom tx", string(j))
}

func testIssueMint(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	to := "to"
	amount := "amount"
	tx, err := c.IssueMint(r, issueId, amount, to)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueMint tx", string(j))
}

func testIssueTransfer(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	to := "to"
	tx, err := c.IssueTransfer(r, issueId, to)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueTransfer tx", string(j))
}

func testIssueDisableFeature(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	r := msg.NewPostIssueReq(br, params)

	issueId := "issueId"
	feature := "feature"
	tx, err := c.IssueDisableFeature(r, issueId, feature)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("IssueDisableFeature tx", string(j))
}

func testLiquidatorSieze(t *testing.T, c client.DexClient) {
	collateralDenom := "collateralDenom"
	accAddr := types.AccAddress([]byte("xar1hcz69efwpsttfx0a3p8xxrkjrmgrm3rwsceut3"))
	r := msg.NewSASCARequest(br, accAddr, accAddr, collateralDenom)
	tx, err := c.LiquidatorSieze(r)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("LiquidatorSieze tx", string(j))
}

func testDebtAuction(t *testing.T, c client.DexClient) {
	accAddr := types.AccAddress([]byte("xar1hcz69efwpsttfx0a3p8xxrkjrmgrm3rwsceut3"))
	issueId := "issueId"
	accAddress := "accAddress"
	amount := "amount"
	r := msg.NewDebtAuctionRequest(br, accAddr)
	tx, err := c.DebtAuction(r, issueId, accAddress, amount)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("DebtAuction tx", string(j))
}

func testPriceRequest(t *testing.T, c client.DexClient) {
	issueId := "issueId"
	assetCode := "accAddress"
	price := "price"
	expiry := "expiry"
	r := msg.NewPriceReq(br, assetCode, price, expiry)
	tx, err := c.PriceRequest(r, issueId)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("PriceRequest tx", string(j))
}
