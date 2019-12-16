package transaction

import "fmt"

const (
	nftsTransferEndpoint = "/nfts/transfer"
	nftsMintEndpoint     = "/nfts/mint"
)

func EditNftsMetadataEndpont(denom, id string) string {
	return fmt.Sprintf("/nfts/collection/%s/nft/%s/metadata", denom, id)
}

func BurnNftEndpoint(denom, id string) string {
	return fmt.Sprintf("/nfts/collection/%s/nft/%s/burn", denom, id)
}