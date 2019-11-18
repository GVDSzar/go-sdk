package msg

import (
	"fmt"
)

const CodeIssueTotalSupplyNotValid = 3507
const CodeUnknownFeature = 3513
const CodeInvalidAddress = 7

type msgsNumberError struct {
	expected int
	got      int
}

func MsgsNumberError(e, g int) error {
	return msgsNumberError{e, g}
}

func (i msgsNumberError) Error() string {
	return fmt.Sprintf("incorrect length of message array. Expected: %v, got: %v", i.expected, i.got)
}

type msgContentError struct {
	msg string
}

func MsgContentError(m string) error {
	return msgContentError{m}
}

func (i msgContentError) Error() string {
	return fmt.Sprintf("stdTx message content differs from original message. %s", i.msg)
}

type coinTotalSupplyMaxValueNotValid string

var ErrorCoinTotalSupplyMaxValueNotValid = coinTotalSupplyMaxValueNotValid(fmt.Sprintf("code: %v, greater than total supply max value", CodeIssueTotalSupplyNotValid))

func (c coinTotalSupplyMaxValueNotValid) Error() string {
	return string(c)
}

type unknownFeatures string

var ErrUnknownFeatures = unknownFeatures(fmt.Sprintf("code: %v, unknown feature", CodeUnknownFeature))

func (c unknownFeatures) Error() string {
	return string(c)
}
