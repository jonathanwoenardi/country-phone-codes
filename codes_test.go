package country_phone_codes

import (
	"fmt"
	"strings"
	"testing"
)

// TestCountryCodeMapPrefixTree ensure that the phone country codes in the map satisfy prefix code property
func TestCountryCodeMapPrefixCodeProperty(t *testing.T) {
	for i := 100; i < 1000; i++ {
		phone := fmt.Sprintf("+%d", i)
		matchingPhoneCodes := []string{}
		for _, phoneCodes := range countryCodeToPhoneCodeMap {
			for _, phoneCode := range phoneCodes {
				if strings.HasPrefix(phone, phoneCode) {
					matchingPhoneCodes = append(matchingPhoneCodes, phoneCode)
				}
			}
		}
		if len(matchingPhoneCodes) > 0 {
			checkPhoneCode := matchingPhoneCodes[0]
			for _, phoneCode := range matchingPhoneCodes {
				if phoneCode != checkPhoneCode {
					t.Errorf("more than one matching phone code | phone: %s | matchingPhoneCodes: %s", phone, matchingPhoneCodes)
				}
			}
		}
	}
}
