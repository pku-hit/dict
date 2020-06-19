package util

type stringUtil struct {
}

var String stringUtil

func (u stringUtil) IsEmptyString(input interface{}) bool {
	if input == nil {
		return true
	}

	str, ok := input.(string)
	if !ok {
		return false
	}

	return len(str) == 0
}

