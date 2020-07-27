package country_phone_codes

import (
	"strings"
)

// IsValidCountryCode return whether country code is ISO 3166 alpha-2 two letter country codes.
func IsValidCountryCode(countryCode string) bool {
	upperCaseCountryCode := strings.ToUpper(countryCode)
	_, found := countryCodeToPhoneCodeMap[upperCaseCountryCode]
	return found
}

// GetPhoneCodesFromCountryCode return all phone codes associated with the country code.
func GetPhoneCodesFromCountryCode(countryCode string) ([]string, error) {
	upperCaseCountryCode := strings.ToUpper(countryCode)
	phoneCodes, found := countryCodeToPhoneCodeMap[upperCaseCountryCode]
	if !found {
		return nil, ErrorInvalidCountryCode
	}

	return phoneCodes, nil
}

// GetPhoneCodeFromPhoneNumber return the phone code of a phone number.
// The prefix code property of the phone code system ensure that the answer is always unique.
// The test TestCountryCodeMapPrefixCodeProperty ensure the prefix code property.
func GetPhoneCodeFromPhoneNumber(phoneNumber string) (string, error) {
	if !isValidPhoneNumber(phoneNumber) {
		return "", ErrorInvalidPhoneNumber
	}

	for _, phoneCodes := range countryCodeToPhoneCodeMap {
		for _, phoneCode := range phoneCodes {
			if strings.HasPrefix(phoneNumber, phoneCode) {
				return phoneCode, nil
			}
		}
	}

	return "", ErrorUnsupportedPhoneNumber
}

// GetCountryCodesFromPhoneNumber return all country codes associated with the phone code.
// It is possible for a phone number to be assoiciated with more than one country code
// because the country code cannot be decided from the phone country code itself.
// For example: Kazakhstan and Russia both use +7 as their phone country code.
// To achieve better result for phone number in North American Numbering Plan (NANP),
// we will use area code to deduce the country code of the phone number.
// However, regions of Canada and United States is not included at the moment.
// Thus, phone number from Canada and United States cannot be differentiated at the moment.
func GetCountryCodesFromPhoneNumber(phoneNumber string) ([]string, error) {
	if !isValidPhoneNumber(phoneNumber) {
		return nil, ErrorInvalidPhoneNumber
	}

	if isNANPPhoneNumber(phoneNumber) {
		return getCountryCodesFromNANPPhoneNumber(phoneNumber), nil
	}

	var countryCodes []string
	for countryCode, phoneCodes := range countryCodeToPhoneCodeMap {
		for _, phoneCode := range phoneCodes {
			if strings.HasPrefix(phoneNumber, phoneCode) {
				countryCodes = append(countryCodes, countryCode)
			}
			break
		}
	}

	return countryCodes, nil
}

func isNANPPhoneNumber(phoneNumber string) bool {
	return strings.HasPrefix(phoneNumber, "+1")
}

// getCountryCodesFromNANPPhoneNumber return all country codes from NANP associated with the phone code.
// If the phone number matches any of the countries in NANP except Canada or United States,
// the function should return just one country code. Otherwise, it will return both Canada and United States.
func getCountryCodesFromNANPPhoneNumber(phoneNumber string) []string {
	var countryCodes []string
	for countryCode, phoneCodes := range nanpAreaCodeToPhoneCodeMap {
		// skip Canada or US
		// if there is no area match, then it is Canada or US
		if countryCode == "CA" || countryCode == "US" {
			continue
		}
		for _, phoneCode := range phoneCodes {
			if strings.HasPrefix(phoneNumber, phoneCode) {
				countryCodes = append(countryCodes, countryCode)
			}
			break
		}
	}

	if len(countryCodes) == 0 {
		return []string{"CA", "US"}
	} else {
		return countryCodes
	}
}
