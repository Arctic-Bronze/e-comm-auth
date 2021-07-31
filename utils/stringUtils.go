package utils

import (
	"net/mail"
)

func IsEmpty(s ...string) bool {
	for _, v := range s {
		if len(v) == 0 {
			return true
		}
	}
	return false
}

func IsGreaterThanEq(s string, lowerBoundry int) bool {
	return len(s) >= lowerBoundry
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
