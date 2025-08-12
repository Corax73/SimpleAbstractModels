package customStructs

import (
	"encoding/json"
	"strconv"

	goutils "github.com/Corax73/goUtils"
)

type SimpleResponse struct {
	Success bool
	Message map[string]any
}

type ListResponse struct {
	Success bool
	Message []map[string]any `json:"data"`
	Total   int `json:"total"`
}

func (listResponse *ListResponse) ToString() string {
	var resp string
	jsonData, err := json.Marshal(listResponse.Message)
	if err != nil {
		goutils.Logging(err)
	} else {
		resp = string(jsonData)
		resp = goutils.ConcatSlice([]string{ "{\"data\":", resp, ",", "\"total\":", strconv.Itoa(listResponse.Total), "}"})
	}
	return resp
}
