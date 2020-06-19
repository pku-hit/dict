package util

import "encoding/json"

type jsonUtil struct {
}

var Json jsonUtil

func (u jsonUtil) ToJsonString(input interface{}) (output string) {
	b, err := json.Marshal(input)
	if err != nil {
		output = ""
	} else {
		output = string(b)
	}
	return
}

func (u jsonUtil) StructToMap(input interface{}) (output map[string]interface{}) {
	output = make(map[string]interface{})
	b, err := json.Marshal(input)
	if err != nil {
		return
	}
	json.Unmarshal(b, &output)
	return
}
