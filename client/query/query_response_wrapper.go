package query

import (
	"github.com/valyala/fastjson"
)

type QueryResponse struct {
	json *fastjson.Value
}

type queryResponseSimplified struct {
	Height string      `json:"height"`
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}

func ResponseWithHeight() *QueryResponse {
	return &QueryResponse{}
}

func (q *QueryResponse) Json() *fastjson.Value {
	return q.json
}

func (q *QueryResponse) MustUnmarshal(j []byte) {
	var p fastjson.Parser
	v, err := p.ParseBytes(j)
	if err != nil {
		panic(err)
	}

	q.json = v
	//qs := new(queryResponseSimplified)
	//err := json.Unmarshal(j, qs)
	//if err != nil {
	//	if len(j) > 0 {
	//		q.Error.Message = string(j)
	//		return
	//	}
	//	panic(err)
	//}
	//
	//if qs.Error != "" {
	//	err = json.Unmarshal([]byte(qs.Error), &q.Error)
	//	if err != nil {
	//		q.Error.Message = qs.Error
	//	}
	//}
	//
	//if qs.Height != "" {
	//	h, err := strconv.ParseInt(qs.Height, 10, 64)
	//	if err != nil {
	//		panic(err)
	//	}
	//	q.Height = uint64(h)
	//}
	//q.Result = qs.Result
}
