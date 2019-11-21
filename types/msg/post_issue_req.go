package msg

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/common/types"
)

type PostIssueReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	IssueParams *IssueParams `json:"issue"`
}

func NewPostIssueReq(rest rest.BaseReq, issueParams *IssueParams) *PostIssueReq {
	return &PostIssueReq{
		rest,
		issueParams,
	}
}

func (r PostIssueReq) Route() string { return "undefined" }

func (r PostIssueReq) Type() string { return "undefined" }

func (r PostIssueReq) ValidateBasic() error {
	// add validaton if needed
	return nil
}

func (m *PostIssueReq) ValidateMsg(res Msg) error {
	// validate incoming msg
	return nil
}

func (r PostIssueReq) GetSignBytes() []byte {
	b, err := Cdc.MarshalJSON(r)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg PostIssueReq) GetSigners() []types.AccAddress {
	return nil //[]types.AccAddress{msg.Buyer}
}

func (msg PostIssueReq) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
