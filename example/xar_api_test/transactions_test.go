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
	"go-sdk/types/tx"
	"log"
	"testing"
)

// it is assumed that std local rest-daemon is up
const baseUrl = "localhost:1317"
const localUserMnemonic = "canal diesel swim expand key puppy agent airport ride shrug walnut erosion toss inside oblige square sort square inspire tackle achieve promote urge field"
const localUserAddr = "xar14p760judpsaengg0zp8g04tgvw8x6dq9zjupse"

func TestTransactions(t *testing.T) {
	km, err := keys.NewMnemonicKeyManager(localUserMnemonic)
	if err != nil {
		panic(err)
	}

	c, err := client.NewCustomClient(baseUrl, types.TestNetwork, km)
	if err != nil {
		t.Errorf(err.Error())
	}
	br.From = c.GetKeyManager().GetAddr().String()

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
	auctionId := "0"
	bidder := "bidder"
	bid := "bid"
	lot := "lot"
	mpb := msg.NewPlaceBidReq(br, auctionId, bidder, bid, lot)
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	tx, err := c.PlaceBidTx(mpb)
	if err != nil {
		t.Error(err.Error())
	}

	j, err := json.Marshal(tx)
	if err != nil {
		t.Error(err.Error())
	}

	log.Println("PlaceBid tx", string(j))
}

func testModifyCsdtTx(t *testing.T, c client.DexClient) {
	collateralDenom := "test"
	collateralAmount := "0"
	debt := "0"
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
	name := "testinggname"
	symbol := "symbolll"
	description := "{\"desc\":\"ription\"}"
	totalSupply := sdk.NewInt(1)

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

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
	name := "testinggname"
	symbol := "SYMBOLLL"
	description := "{\"desc\":\"ription\"}"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	r := msg.NewPostIssueReq(br, params)

	issueId := "xar174876e802"
	accAddr := "xar14azd4p7ks2lmmq8xpxdztwe67rsfnpzt0hnqpl"
	amount := "1"

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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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

	log.Println("MsgIssueSendFrom tx", string(j))
}

func testIssueMint(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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

	log.Println("MsgIssueMint tx", string(j))
}

func testIssueTransfer(t *testing.T, c client.DexClient) {
	name := "issueName"
	symbol := "symbol"
	description := "description"
	totalSupply := sdk.Int{}

	params := msg.NewIssueParamsBm(name, symbol, description, totalSupply, 0)
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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

// temporary method. will be removed in a following commits
func TestPostRequestsVer(t *testing.T) {
	km, err := keys.NewMnemonicKeyManager(localUserMnemonic)
	if err != nil {
		panic(err)
	}
	log.Println(km.GetAddr().String())

	c, err := client.NewCustomClient(baseUrl, types.TestNetwork, km)
	if err != nil {
		t.Errorf(err.Error())
	}

	//tx.RegisterCodec(amino.NewCodec())
	var b = []byte(`{
        "type": "issue/MsgIssueApprove",
        "value": {
          "issue_id": "xar174876e802",
          "from_address": "xar14p760judpsaengg0zp8g04tgvw8x6dq9zjupse",
          "to_address": "xar14azd4p7ks2lmmq8xpxdztwe67rsfnpzt0hnqpl",
          "amount": "1"
        }
}`)
	b = msg.MustSortJSON(b)
	log.Println("b sorted:", string(b))
	//stdTxBasic := new(tx.StdTxBasic)
	msgIssueApprove := msg.MsgIssueApprove{}
	tx.Cdc.MustUnmarshalJSON(b, &msgIssueApprove)

	br.From = c.GetKeyManager().GetAddr().String()
	//
	//err = c.CollectAccountInfo()
	//if err != nil {
	//	t.Errorf(err.Error())
	//}
	testIssueApproveTx(t, c)

	//testDebtAuction(t, c)
}

func testDebtAuction(t *testing.T, c client.DexClient) {
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	r := msg.NewDebtAuctionRequest(br, c.GetKeyManager().GetAddr())
	tx, err := c.DebtAuction(r)
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
	err := c.CollectAccountInfo()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
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
