package main

import (
	"go-sdk/client"
	"go-sdk/common/types"
	"log"
)

const baseUrl = "localhost:1317"

func main() {
	types.Network = types.TestNetwork
	c, err := client.NewCustomClient(baseUrl, types.TestNetwork, nil)
	if err != nil {
		log.Fatal(err)
	}
	c.GetShortNodeInfo()
}
