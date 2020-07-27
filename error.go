package country_phone_codes

import (
	"errors"
)

// Error constants
var (
	ErrorInvalidCountryCode     = errors.New("invalid country code")     // country codes is not a ISO 3166 alpha-2 two letter country codes
	ErrorInvalidPhoneNumber     = errors.New("invalid phone number")     // phone number is not a string started with "+" symbols followed by 7 to 15 numerical digits
	ErrorUnsupportedPhoneNumber = errors.New("unsupported phone number") // phone number country code is not associated with a geographical region
)
