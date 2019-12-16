package msg

import (
	"github.com/tendermint/go-amino"
)

var Cdc *amino.Codec

func RegisterCodec(cdc *amino.Codec) {
	Cdc = cdc
	cdc.RegisterInterface((*Msg)(nil), nil)

	cdc.RegisterConcrete(CreateOrderMsg{}, "dex/NewOrder", nil)
	cdc.RegisterConcrete(CancelOrderMsg{}, "dex/CancelOrder", nil)
	cdc.RegisterConcrete(TokenIssueMsg{}, "tokens/IssueMsg", nil)
	cdc.RegisterConcrete(TokenBurnMsg{}, "tokens/BurnMsg", nil)

	cdc.RegisterConcrete(TimeLockMsg{}, "tokens/TimeLockMsg", nil)
	cdc.RegisterConcrete(TokenFreezeMsg{}, "tokens/FreezeMsg", nil)
	cdc.RegisterConcrete(TokenUnfreezeMsg{}, "tokens/UnfreezeMsg", nil)

	cdc.RegisterConcrete(TimeUnlockMsg{}, "tokens/TimeUnlockMsg", nil)
	cdc.RegisterConcrete(TimeRelockMsg{}, "tokens/TimeRelockMsg", nil)

	cdc.RegisterConcrete(HTLTMsg{}, "tokens/HTLTMsg", nil)
	cdc.RegisterConcrete(DepositHTLTMsg{}, "tokens/DepositHTLTMsg", nil)
	cdc.RegisterConcrete(ClaimHTLTMsg{}, "tokens/ClaimHTLTMsg", nil)
	cdc.RegisterConcrete(RefundHTLTMsg{}, "tokens/RefundHTLTMsg", nil)

	cdc.RegisterConcrete(DexListMsg{}, "dex/ListMsg", nil)
	//cdc.RegisterConcrete(BuyNameReq{}, "nameservice/BuyNameReq", nil)
	cdc.RegisterConcrete(BuyName{}, "nameservice/BuyName", nil)
	//cdc.RegisterConcrete(MintMsg{}, "tokens/MintMsg", nil)
	//Must use cosmos-sdk.
	cdc.RegisterConcrete(SendMsg{}, "cosmos-sdk/Send", nil)

	cdc.RegisterConcrete(SubmitProposalMsg{}, "cosmos-sdk/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(DepositMsg{}, "cosmos-sdk/MsgDeposit", nil)
	cdc.RegisterConcrete(VoteMsg{}, "cosmos-sdk/MsgVote", nil)

	cdc.RegisterConcrete(SetAccountFlagsMsg{}, "scripts/SetAccountFlagsMsg", nil)

	cdc.RegisterConcrete(MsgCreateValidator{}, "cosmos-sdk/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgRemoveValidator{}, "cosmos-sdk/MsgRemoveValidator", nil)
	cdc.RegisterConcrete(MsgCreateValidatorProposal{}, "cosmos-sdk/MsgCreateValidatorProposal", nil)

	cdc.RegisterConcrete(MsgIssue{}, "issue/MsgIssue", nil)

	cdc.RegisterConcrete(PlaceBid{}, "auction/MsgPlaceBid", nil)
	cdc.RegisterConcrete(&ForwardAuction{}, "auction/ForwardAuction", nil)
	cdc.RegisterConcrete(&ReverseAuction{}, "auction/ReverseAuction", nil)
	cdc.RegisterConcrete(&ForwardReverseAuction{}, "auction/ForwardReverseAuction", nil)

	cdc.RegisterConcrete(MsgCreateCompound{}, "compound/CreateCompound", nil)
	cdc.RegisterConcrete(MsgSupplyMarket{}, "moneymarket/SupplyCompound", nil)
	cdc.RegisterConcrete(MsgBorrowFromMarket{}, "moneymarket/BorrowFromMarket", nil)
	cdc.RegisterConcrete(MsgRedeemFromMarket{}, "moneymarket/RedeemFromMarket", nil)
	cdc.RegisterConcrete(MsgRepayToMarket{}, "moneymarket/RepayToMarket", nil)

	cdc.RegisterConcrete(MsgCreateOrModifyCSDT{}, "csdt/MsgCreateOrModifyCSDT", nil)
	cdc.RegisterConcrete(MsgTransferCSDT{}, "csdt/MsgTransferCSDT", nil)

	cdc.RegisterConcrete(MsgIssueTransferOwnership{}, "issue/MsgIssueTransferOwnership", nil)
	cdc.RegisterConcrete(MsgIssueDescription{}, "issue/MsgIssueDescription", nil)
	cdc.RegisterConcrete(MsgIssueMint{}, "issue/MsgIssueMint", nil)
	cdc.RegisterConcrete(MsgIssueBurnOwner{}, "issue/MsgIssueBurnOwner", nil)
	cdc.RegisterConcrete(MsgIssueBurnHolder{}, "issue/MsgIssueBurnHolder", nil)
	cdc.RegisterConcrete(MsgIssueBurnFrom{}, "issue/MsgIssueBurnFrom", nil)
	cdc.RegisterConcrete(MsgIssueDisableFeature{}, "issue/MsgIssueDisableFeature", nil)
	cdc.RegisterConcrete(MsgIssueApprove{}, "issue/MsgIssueApprove", nil)
	cdc.RegisterConcrete(MsgIssueSendFrom{}, "issue/MsgIssueSendFrom", nil)
	cdc.RegisterConcrete(MsgIssueIncreaseApproval{}, "issue/MsgIssueIncreaseApproval", nil)
	cdc.RegisterConcrete(MsgIssueDecreaseApproval{}, "issue/MsgIssueDecreaseApproval", nil)
	cdc.RegisterConcrete(MsgIssueFreeze{}, "issue/MsgIssueFreeze", nil)
	cdc.RegisterConcrete(MsgIssueUnFreeze{}, "issue/MsgIssueUnFreeze", nil)

	cdc.RegisterConcrete(MsgIncreaseCredit{}, "issuer/MsgIncreaseCredit", nil)
	cdc.RegisterConcrete(MsgDecreaseCredit{}, "issuer/MsgDecreaseCredit", nil)
	cdc.RegisterConcrete(MsgRevokeLiquidityProvider{}, "issuer/MsgRevokeLiquidityProvider", nil)
	cdc.RegisterConcrete(MsgSetInterest{}, "issuer/MsgSetInterest", nil)

	cdc.RegisterConcrete(MsgSeizeAndStartCollateralAuction{}, "liquidator/MsgSeizeAndStartCollateralAuction", nil)
	cdc.RegisterConcrete(MsgStartDebtAuction{}, "liquidator/MsgStartDebtAuction", nil)

	cdc.RegisterConcrete(MsgMintTokens{}, "liquidityprovider/MsgMintTokens", nil)
	cdc.RegisterConcrete(MsgBurnTokens{}, "liquidityprovider/MsgBurnTokens", nil)
	cdc.RegisterConcrete(&LiquidityProviderAccount{}, "liquidityprovider/LiquidityProviderAccount", nil)

	cdc.RegisterConcrete(&BaseNFT{}, "cosmos-sdk/BaseNFT", nil)
	cdc.RegisterConcrete(&IDCollection{}, "cosmos-sdk/IDCollection", nil)
	cdc.RegisterConcrete(&Collection{}, "cosmos-sdk/Collection", nil)
	cdc.RegisterConcrete(&Owner{}, "cosmos-sdk/Owner", nil)
	cdc.RegisterConcrete(MsgTransferNFT{}, "cosmos-sdk/MsgTransferNFT", nil)
	cdc.RegisterConcrete(MsgEditNFTMetadata{}, "cosmos-sdk/MsgEditNFTMetadata", nil)
	cdc.RegisterConcrete(MsgMintNFT{}, "cosmos-sdk/MsgMintNFT", nil)
	cdc.RegisterConcrete(MsgBurnNFT{}, "cosmos-sdk/MsgBurnNFT", nil)

	cdc.RegisterConcrete(MsgPostPrice{}, "oracle/MsgPostPrice", nil)

	cdc.RegisterConcrete(MsgPost{}, "order/Post", nil)
	cdc.RegisterConcrete(MsgCancel{}, "order/Cancel", nil)

	cdc.RegisterConcrete(MsgDepositFund{}, "pool/MsgDepositFund", nil)
	cdc.RegisterConcrete(MsgWithdrawFund{}, "pool/MsgWithdrawFund", nil)
	cdc.RegisterConcrete(MsgRecord{}, "record/MsgRecord", nil)
	cdc.RegisterConcrete(&RecordInfo{}, "record/RecordInfo", nil)

	cdc.RegisterConcrete(MsgSwapOrder{}, "uniswap/MsgSwapOrder", nil)
	cdc.RegisterConcrete(MsgAddLiquidity{}, "uniswap/MsgAddLiquidity", nil)
	cdc.RegisterConcrete(MsgRemoveLiquidity{}, "uniswap/MsgRemoveLiquidity", nil)

	cdc.RegisterConcrete(MsgIssueToken{}, "denominations/MsgIssueToken", nil)
	cdc.RegisterConcrete(MsgMintCoins{}, "denominations/MsgMintCoins", nil)
	cdc.RegisterConcrete(MsgBurnCoins{}, "denominations/MsgBurnCoins", nil)
	cdc.RegisterConcrete(MsgFreezeCoins{}, "denominations/MsgFreezeCoins", nil)
	cdc.RegisterConcrete(MsgUnfreezeCoins{}, "denominations/MsgUnfreezeCoins", nil)

	cdc.RegisterConcrete(MsgPostPrice{}, "oracle/MsgPostPrice", nil)
	cdc.RegisterConcrete(MsgAddOracle{}, "oracle/MsgAddOracle", nil)
	cdc.RegisterConcrete(MsgSetOracles{}, "oracle/MsgSetOracles", nil)
	cdc.RegisterConcrete(MsgAddAsset{}, "oracle/MsgAddAsset", nil)
	cdc.RegisterConcrete(MsgSetAsset{}, "oracle/MsgSetAsset", nil)
}

func init() {
	//RegisterCodec(Cdc)
}
