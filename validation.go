package country_phone_codes

import (
	"regexp"
)

var (
	phoneNumberRegexp *regexp.Regexp
)

func init() {
	phoneNumberRegexp = regexp.MustCompile("\\+[0-9]{7,15}")
}

func isValidPhoneNumber(phoneNumber string) bool {
	regexMatched := phoneNumberRegexp.MatchString(phoneNumber)
	lengthLimited := len(phoneNumber) <= 16
	return regexMatched && lengthLimited
}
