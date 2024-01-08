package utils

import "regexp"

func IsValidPhoneNumber(phoneNumber string) bool {
	pattern := `^\d{3}-\d{4}-\d{4}$`
	matched, err := regexp.MatchString(pattern, phoneNumber)
	if err != nil {
		return false
	}
	return matched
}