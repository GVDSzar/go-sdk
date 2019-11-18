package transaction

type RestTransactionError string

func (r RestTransactionError) Error() string {
	return string(r)
}
