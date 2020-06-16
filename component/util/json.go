package util

import "encoding/json"

func ToJsonString(input interface{}) (output string) {
	b, err := json.Marshal(input)
	if err != nil {
		output = ""
	} else {
		output = string(b)
	}
	return
}
