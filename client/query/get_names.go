package query

import (
	"encoding/json"
	"go-sdk/common"
	"go-sdk/common/types"
)

func (c *client) GetNames() (*types.NameList, error) {
	var nl = new(types.NameList)
	resp, code, err := c.baseClient.Get("/nameservice/names", emptyQP)
	if err != nil {
		return nl, common.NewRestyHttpError(resp, code, err)
	}

	return nl, json.Unmarshal(resp, nl)
}
