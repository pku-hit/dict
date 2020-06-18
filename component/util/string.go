package util

func IsEmptyString(input interface{}) bool {
	if input == nil {
		return true
	}

	str, ok := input.(string)
	if !ok {
		return false
	}

	return len(str) == 0
}

func NilIfEmpty(input string) (out string) {
	if IsEmptyString(input) {
		out = "nil"
	} else {
		out = input
	}
	return
}
