package transaction

import (
	ctypes "go-sdk/common/types"
	"go-sdk/types"
	"go-sdk/types/msg"
	"go-sdk/types/tx"
)

type DepositProposalResult struct {
	tx.TxCommitResult
}

func (c *client) DepositProposal(proposalID int64, amount int64, sync bool, options ...Option) (*DepositProposalResult, error) {
	fromAddr := c.keyManager.GetAddr()
	coins := ctypes.Coins{ctypes.Coin{Denom: types.NativeSymbol, Amount: amount}}
	depositMsg := msg.NewDepositMsg(fromAddr, proposalID, coins)
	commit, err := c.broadcastMsg(depositMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &DepositProposalResult{*commit}, err

}
