package main

import (
	"encoding/hex"
	"github.com/tendermint/go-amino"
	"go-sdk/common/types"
	"go-sdk/keys"
	"log"
)

func main() {
	types.Network = types.TestNetwork
	keyManager, _ := keys.NewMnemonicKeyManager("exile toilet mosquito dignity rival device toy wheat cherry ball empty snack topic chapter quiz fitness slab fault movie wave group addict sick pupil")

	priv, err := keyManager.ExportAsPrivateKey()
	if err != nil {
		panic(err)
	}
	println("priv", priv)
	println("bech32 addr", keyManager.GetAddr().String())
	println("hex addr", hex.EncodeToString(keyManager.GetAddr().Bytes()))
	println("hex addr", keyManager)

	b := []byte(`{"account_number":"0","chain_id":"namechain","fee":{"amount":[],"gas":"200000"},"memo":"","msgs":[{"type":"nameservice/BuyName","value":{"bid":[{"amount":"5","denom":"nametoken"}],"buyer":"cosmos1hse55at9f4ha5pm0tmjksdurc38jfzf6znzq0h","name":"testname15"}}],"sequence":"0"}`)
	privSig, err := keyManager.GetPrivKey().Sign(b)
	if err != nil {
		panic(err)
	}

	v, err := amino.MarshalJSON(privSig)
	if err != nil {
		panic(err)
	}

	log.Println("json for sign:", string(b))
	log.Println("privSig amino s", string(v))
	log.Println("privSig encode", hex.EncodeToString(privSig))
	log.Println("privSig raw", privSig)
	log.Println("")
}
//{\"type\":\"cosmos-sdk/StdTx\",\"value\":{\"msg\":[{\"type\":\"nameservice/BuyName\",\"value\":{\"name\":\"testname15\",\"bid\":[{\"denom\":\"nametoken\",\"amount\":\"5\"}],\"buyer\":\"cosmos1hse55at9f4ha5pm0tmjksdurc38jfzf6znzq0h\"}}],\"fee\":{\"amount\":[],\"gas\":\"200000\"},\"signatures\":null,\"memo\":\"\"}}

//{\"account_number\":\"2\",\"chain_id\":\"namechain\",\"fee\":{\"amount\":[],\"gas\":\"200000\"},\"memo\":\"\",\"msgs\":[{\"type\":\"nameservice/BuyName\",\"value\":{\"bid\":[{\"amount\":\"5\",\"denom\":\"nametoken\"}],\"buyer\":\"cosmos1hse55at9f4ha5pm0tmjksdurc38jfzf6znzq0h\",\"name\":\"testname15\"}}],\"sequence\":\"0\"}

//{\"type\":\"cosmos-sdk/StdTx\",\"value\":{\"msg\":[{\"type\":\"nameservice/BuyName\",\"value\":{\"name\":\"testname15\",\"bid\":[{\"denom\":\"nametoken\",\"amount\":\"5\"}],\"buyer\":\"cosmos1hse55at9f4ha5pm0tmjksdurc38jfzf6znzq0h\"}}],\"fee\":{\"amount\":[],\"gas\":\"200000\"},\"signatures\":null,\"memo\":\"\"}}