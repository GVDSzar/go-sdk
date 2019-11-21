package msg

type IDCollection struct {
	Denom string            `json:"denom" yaml:"denom"`
	IDs   SortedStringArray `json:"ids" yaml:"ids"`
}

type SortedStringArray []string

type NFTs []BaseNFT

type Collection struct {
	Denom string `json:"denom,omitempty" yaml:"denom"` // name of the collection; not exported to clients
	NFTs  NFTs   `json:"nfts" yaml:"nfts"`             // NFTs that belong to a collection
}
