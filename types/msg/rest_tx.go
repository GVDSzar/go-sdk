package msg

//Rest Transaction is an interface
type RestTransactionRequest interface {
	Msg
	// Takes a response message for a transaction request and validates it
	ValidateMsg(Msg) error
}
