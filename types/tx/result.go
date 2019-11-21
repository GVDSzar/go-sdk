package tx

const (
	CodeOk int32 = 0
)

// TxResult def
type TxResult struct {
	Hash string `json:"hash"`
	Log  string `json:"log"`
	Data string `json:"data"`
	Code int32  `json:"code"`
}

type QueryErrorResult struct {
	CodeSpace string `json:"codespace"`
	Code      uint64 `json:"code"`
	Message   string `json:"message"`
}

// TxCommitResult for POST tx results
type TxCommitResult struct {
	Ok     bool   `json:"ok"`
	Height string `json:"height"`
	Log    string `json:"log"`
	RaeLog string `json:"raw_log"`
	Hash   string `json:"hash"`
	Code   int32  `json:"code"`
	Data   string `json:"data"`
}

type TxBroadcastResult struct {
	Height string  `json:"height"`
	TxHash string  `json:"txhash"`
	RawLog string  `json:"raw_log"`
	Logs   []TxLog `json:"logs"`
}

type TxLog struct {
	MsgIndex uint64     `json:"msg_index"`
	Success  bool       `json:"success"`
	Log      string     `json:"log"`
	Events   []TxEvents `json:"events"`
}

type TxEvents struct {
	Type       string        `json:"type"`
	Attributes []TxAttribute `json:"attributes"`
}

type TxAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
