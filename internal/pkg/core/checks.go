package core

func IsMissingInput(fields ...string) bool {
	i := len(fields) - 1

	for ; i > 0; i-- {
		if len(fields[i]) == 0 {
			return true
		}
	}

	return false
}

func IsInputLengthTooLong(lim int, fields ...string) bool {
	i := len(fields) - 1

	for ; i > 0; i-- {
		if len(fields[i]) > lim {
			return true
		}
	}

	return false
}
