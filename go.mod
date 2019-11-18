module go-sdk

require (
	github.com/binance-chain/go-sdk v1.1.3
	github.com/binance-chain/ledger-cosmos-go v0.9.9-binance.1
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/cosmos/cosmos-sdk v0.37.4
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
<<<<<<< HEAD
=======
	github.com/fatih/structs v1.1.0
>>>>>>> 16a09dfaea01e4a04f95037c10a54d6c2ee4e454
	github.com/gorilla/websocket v1.4.1
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.7
	github.com/valyala/fasthttp v1.6.0
	github.com/valyala/fastjson v1.4.1
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	gopkg.in/resty.v1 v1.12.0
)

replace github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1

replace github.com/zondax/ledger-go => github.com/binance-chain/ledger-go v0.9.1

go 1.13
