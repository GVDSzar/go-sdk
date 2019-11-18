package query

import (
	"encoding/json"
	"go-sdk/common/types"
)

func (c *client) GetShortNodeInfo() (*types.ResultStatus, error) {
	qp := map[string]string{}
	resp, _, err := c.baseClient.Get("/node_info", qp)
	if err != nil {
		return nil, err
	}
	var resultStatus types.ResultStatus
	if err := json.Unmarshal(resp, &resultStatus); err != nil {
		return nil, err
	}

	return &resultStatus, nil
}
