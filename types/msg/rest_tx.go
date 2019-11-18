package msg

//Rest Transaction is an interface
type RestTransactionRequest interface {
	Msg
	// Takes a response message for a transaction request
	// validates it and signs via PK from client key manager
	ValidateMsg(Msg) error
}
