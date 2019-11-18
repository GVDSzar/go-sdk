package query

import (
	"encoding/json"
	"go-sdk/types/tx"
	"strconv"
)

type QueryResponse struct {
	Height uint64              `json:"height"`
	Result interface{}         `json:"result"`
	Error  tx.QueryErrorResult `json:"error"`
}

type queryResponseSimplified struct {
	Height string      `json:"height"`
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

func WrapQueryResponse(resp interface{}) *QueryResponse {
	return &QueryResponse{
		Result: resp,
	}
}

func (q *QueryResponse) MustUnmarshal(j []byte) {
	qs := new(queryResponseSimplified)
	err := json.Unmarshal(j, qs)
	if err != nil {
		if len(j) > 0 {
			q.Error.Message = string(j)
			return
		}
		panic(err)
	}

	if qs.Error != "" {
		err = json.Unmarshal([]byte(qs.Error), &q.Error)
		if err != nil {
			q.Error.Message = qs.Error
		}
	}

	if qs.Height != "" {
		_, err = strconv.ParseInt(qs.Height, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	q.Result = qs.Result
}
