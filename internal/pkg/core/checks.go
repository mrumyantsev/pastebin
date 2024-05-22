package core

import "net/mail"

func IsInputMissing(fields ...string) bool {
	i := len(fields) - 1

	for ; i > 0; i-- {
		if len(fields[i]) == 0 {
			return true
		}
	}

	return false
}

func IsInputExceeds(lim int, fields ...string) bool {
	i := len(fields) - 1

	for ; i > 0; i-- {
		if len(fields[i]) > lim {
			return true
		}
	}

	return false
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}
