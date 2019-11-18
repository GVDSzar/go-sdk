package transaction

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"log"
	"net/http"
	"time"

	"go-sdk/client/basic"
	"go-sdk/client/query"
	"go-sdk/common/types"
	"go-sdk/keys"
	"go-sdk/types/msg"
	"go-sdk/types/tx"
)

type Option = tx.Option

var (
	WithSource           = tx.WithSource
	WithMemo             = tx.WithMemo
	WithAcNumAndSequence = tx.WithAcNumAndSequence
)

type TransactionClient interface {
	CreateOrder(baseAssetSymbol, quoteAssetSymbol string, op int8, price, quantity int64, sync bool, options ...Option) (*CreateOrderResult, error)
	CancelOrder(baseAssetSymbol, quoteAssetSymbol, refId string, sync bool, options ...Option) (*CancelOrderResult, error)
	BurnToken(symbol string, amount int64, sync bool, options ...Option) (*BurnTokenResult, error)
	ListPair(proposalId int64, baseAssetSymbol string, quoteAssetSymbol string, initPrice int64, sync bool, options ...Option) (*ListPairResult, error)
	FreezeToken(symbol string, amount int64, sync bool, options ...Option) (*FreezeTokenResult, error)
	UnfreezeToken(symbol string, amount int64, sync bool, options ...Option) (*UnfreezeTokenResult, error)
	IssueToken(name, symbol string, supply int64, sync bool, mintable bool, options ...Option) (*IssueTokenResult, error)
	SendToken(transfers []msg.Transfer, sync bool, options ...Option) (*SendTokenResult, error)
	MintToken(symbol string, amount int64, sync bool, options ...Option) (*MintTokenResult, error)
	TimeLock(description string, amount types.Coins, lockTime int64, sync bool, options ...Option) (*TimeLockResult, error)
	TimeUnLock(id int64, sync bool, options ...Option) (*TimeUnLockResult, error)
	TimeReLock(id int64, description string, amount types.Coins, lockTime int64, sync bool, options ...Option) (*TimeReLockResult, error)
	SetAccountFlags(flags uint64, sync bool, options ...Option) (*SetAccountFlagsResult, error)
	AddAccountFlags(flagOptions []types.FlagOption, sync bool, options ...Option) (*SetAccountFlagsResult, error)
	HTLT(recipient types.AccAddress, recipientOtherChain, senderOtherChain string, randomNumberHash []byte, timestamp int64, amount types.Coins, expectedIncome string, heightSpan int64, crossChain bool, sync bool, options ...Option) (*HTLTResult, error)
	DepositHTLT(swapID []byte, amount types.Coins, sync bool, options ...Option) (*DepositHTLTResult, error)
	ClaimHTLT(swapID []byte, randomNumber []byte, sync bool, options ...Option) (*ClaimHTLTResult, error)
	RefundHTLT(swapID []byte, sync bool, options ...Option) (*RefundHTLTResult, error)

	SubmitListPairProposal(title string, param msg.ListTradingPairParams, initialDeposit int64, votingPeriod time.Duration, sync bool, options ...Option) (*SubmitProposalResult, error)
	SubmitProposal(title string, description string, proposalType msg.ProposalKind, initialDeposit int64, votingPeriod time.Duration, sync bool, options ...Option) (*SubmitProposalResult, error)
	DepositProposal(proposalID int64, amount int64, sync bool, options ...Option) (*DepositProposalResult, error)
	VoteProposal(proposalID int64, option msg.VoteOption, sync bool, options ...Option) (*VoteProposalResult, error)

	GetKeyManager() keys.KeyManager

	// Wrapper for a post/put transactions. Allows to request an arbitrary path relative to baseURL
	RestTransaction(txRequest msg.RestTransactionRequest, path, method string) (txRes *tx.TxCommitResult, err error)

	// Xar requests
	BuyNameTx(br *rest.BaseReq, name, amount, buyer string) (*tx.TxCommitResult, error)
	PlaceBidTx(r *msg.PlaceBidReq) (*tx.TxCommitResult, error)
	ModifyCsdtTx(r *msg.ModifyCsdtReq) (*tx.TxCommitResult, error)
	IssueTx(r *msg.PostIssueReq) (*tx.TxCommitResult, error)
	IssueApproveTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxCommitResult, error)
	IssueIncreaseApprovalTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxCommitResult, error)
	IssueDecreaseApprovalTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxCommitResult, error)
	IssueBurnTx(r *msg.PostIssueReq, issueId, amount string) (*tx.TxCommitResult, error)
	IssueBurnFromTx(r *msg.PostIssueReq, issueId, accAddress, amount string) (*tx.TxCommitResult, error)
	IssueFreeze(r *msg.PostIssueReq, freezeType, issueId, accAddress string) (*tx.TxCommitResult, error)
	IssueUnfreeze(r *msg.PostIssueReq, freezeType, issueId, accAddress string) (*tx.TxCommitResult, error)
	IssueSendFrom(r *msg.PostIssueReq, issueId, from, to, amount string) (*tx.TxCommitResult, error)
	IssueMint(r *msg.PostIssueReq, issueId, amount, to string) (*tx.TxCommitResult, error)
	IssueTransfer(r *msg.PostIssueReq, issueId, to string) (*tx.TxCommitResult, error)
	IssueDisableFeature(r *msg.PostIssueReq, issueId, feature string) (*tx.TxCommitResult, error)
	LiquidatorSieze(r *msg.SeizeAndStartCollateralAuctionRequest) (*tx.TxCommitResult, error)
	DebtAuction(r *msg.StartDebtAuctionRequest, issueId, accAddress, amount string) (*tx.TxCommitResult, error)
	PriceRequest(r *msg.PriceReq, issueId string) (*tx.TxCommitResult, error)
}

type client struct {
	basicClient basic.BasicClient
	queryClient query.QueryClient
	keyManager  keys.KeyManager
	chainId     string
}

func NewClient(chainId string, keyManager keys.KeyManager, queryClient query.QueryClient, basicClient basic.BasicClient) TransactionClient {
	return &client{basicClient, queryClient, keyManager, chainId}
}

// a common function to perform arbitrary transactions.
// requests an stdTx message, validates it, signs and broadcasts
func (c *client) RestTransaction(txRequest msg.RestTransactionRequest, path, method string) (txRes *tx.TxCommitResult, err error) {
	stdTx, err := c.getValidStdTx(txRequest, path, method)
	if err != nil {
		return
	}

	sTx, err := c.signTransaction(stdTx, "sync")
	if err != nil {
		return
	}

	return c.getCommitResult(sTx)
}

func (c *client) getValidStdTx(txRequest msg.RestTransactionRequest, path, method string) (t *tx.StdTxBasic, err error) {
	if txRequest == nil {
		return nil, RestTransactionError("request is nil")
	}
	t = new(tx.StdTxBasic)

	reqBody := txRequest.GetSignBytes()
	res, err := c.basicClient.FastHttpRequest(path, method, reqBody)
	if err != nil {
		return
	}

	if err = tx.Cdc.UnmarshalJSON(res, t); err != nil {
		return
	}

	for _, v := range t.GetMsgs() {
		err = v.ValidateBasic()
		if err != nil {
			return
		}

		err = txRequest.ValidateMsg(v)
		if err != nil {
			return
		}
	}
	return
}

func (c *client) getSignatures(stdTx *tx.StdTxBasic) (sigs []tx.StdSignatureBasic, err error) {
	bytesToSign := tx.StdSignBytesWithFee(c.chainId, 0, 0, stdTx.GetMsgs(), stdTx.Memo, stdTx.Fee)
	sig, err := c.keyManager.GetPrivKey().Sign(bytesToSign)
	if err != nil {
		return
	}

	return []tx.StdSignatureBasic{
		{
			PubKey:    c.keyManager.GetPrivKey().PubKey(),
			Signature: sig,
		},
	}, nil
}

func (c *client) signTransaction(stdTx *tx.StdTxBasic, transactionMode string) (*tx.TxBasic, error) {
	sigs, err := c.getSignatures(stdTx)
	if err != nil {
		return nil, err
	}

	return tx.NewDefaultTransaction(stdTx.GetMsgs(), sigs, stdTx.Fee, stdTx.Memo, transactionMode), nil
}

func (c *client) getCommitResult(stdTx *tx.TxBasic) (txRes *tx.TxCommitResult, err error) {
	rawResult, err := c.broadcastRawMsg(stdTx)
	if err != nil {
		return nil, err
	}

	txRes = new(tx.TxCommitResult)
	return txRes, tx.Cdc.UnmarshalJSON(rawResult, txRes)
}

func (c *client) broadcastRawMsg(stdTx *tx.TxBasic) ([]byte, error) {
	// temprorary solution.
	// constant variable is expected to become either a package-level singleton or a config-defined param
	const targetUrl = "/txs"

	rawMsg := tx.Cdc.MustMarshalJSON(stdTx)
	log.Println("rawMsg final", string(rawMsg))
	res, err := c.basicClient.FastHttpRequest(targetUrl, http.MethodPost, rawMsg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) GetKeyManager() keys.KeyManager {
	return c.keyManager
}

func (c *client) broadcastMsg(m msg.Msg, sync bool, options ...Option) (*tx.TxCommitResult, error) {
	// prepare message to sign
	signMsg := &tx.StdSignMsg{
		ChainID:       c.chainId,
		AccountNumber: -1,
		Sequence:      -1,
		Memo:          "",
		Msgs:          []msg.Msg{m},
		Source:        tx.Source,
	}

	for _, op := range options {
		signMsg = op(signMsg)
	}

	if signMsg.Sequence == -1 || signMsg.AccountNumber == -1 {
		fromAddr := c.keyManager.GetAddr()
		acc, err := c.queryClient.GetAccount(fromAddr.String())
		if err != nil {
			return nil, err
		}
		signMsg.Sequence = acc.Sequence
		signMsg.AccountNumber = acc.Number
	}

	// special logic for createOrder, to save account query
	if orderMsg, ok := m.(msg.CreateOrderMsg); ok {
		orderMsg.ID = msg.GenerateOrderID(signMsg.Sequence+1, c.keyManager.GetAddr())
		signMsg.Msgs[0] = orderMsg
	}

	for _, m := range signMsg.Msgs {
		if err := m.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	rawBz, err := c.keyManager.Sign(*signMsg)
	if err != nil {
		return nil, err
	}
	// Hex encoded signed transaction, ready to be posted to BncChain API
	hexTx := []byte(hex.EncodeToString(rawBz))
	param := map[string]string{}
	if sync {
		param["sync"] = "true"
	}
	commits, err := c.basicClient.PostTx(hexTx, param)
	if err != nil {
		return nil, err
	}
	if len(commits) < 1 {
		return nil, fmt.Errorf("Len of tx Commit result is less than 1 ")
	}
	return &commits[0], nil
}

//b32Public, err := bech32.ConvertAndEncode("cosmospub", c.keyManager.GetPrivKey().PubKey().Bytes())
//if err != nil {
//panic(err)
//}
//
//log.Println("public bench32", b32Public)
