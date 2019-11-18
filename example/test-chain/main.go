package main

//#############################################
//## A file for a local chain tests          ##
//## Will be removed in a subsequent commits ##
//#############################################

import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"go-sdk/client"
	"go-sdk/common/types"
	"go-sdk/keys"
	"go-sdk/types/msg"
	"log"
	"net/http"
)

const baseUrl = "localhost:1317"

func main() {
	keyManager, _ := keys.NewMnemonicKeyManager("exile toilet mosquito dignity rival device toy wheat cherry ball empty snack topic chapter quiz fitness slab fault movie wave group addict sick pupil")

	types.Network = types.TestNetwork
	c, err := client.NewCustomClient(baseUrl, types.TestNetwork, keyManager)
	if err != nil {
		log.Fatal(err)
	}
	nl, err := c.GetNames()
	log.Println(nl)

	var br rest.BaseReq
	br.From = "cosmos1hse55at9f4ha5pm0tmjksdurc38jfzf6znzq0h"
	br.ChainID = "namechain"

	m, err := msg.NewBuyNameRequest(&br, "testname15", "5nametoken", br.From)
	if err != nil {
		panic(err)
	}

	result, err := c.RestTransaction(m, "/nameservice/names", http.MethodPost)
	if err != nil {
		panic(err)
	}

	log.Println("transaction response", result)
}
