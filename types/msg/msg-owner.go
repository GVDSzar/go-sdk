package msg

import "go-sdk/common/types"

type IDCollections []IDCollection

type Owner struct {
	Address       types.AccAddress `json:"address" yaml:"address"`
	IDCollections IDCollections    `json:"idCollections" yaml:"idCollections"`
}
