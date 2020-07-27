package country_phone_codes

import (
	"sort"
	"strings"
	"testing"
)

func checkEquality(As, Bs []string) bool {
	if len(As) != len(Bs) {
		return false
	}

	if len(As) == 0 {
		return true
	}

	sort.Slice(As, func(i, j int) bool { return strings.Compare(As[i], As[j]) == -1 })
	sort.Slice(Bs, func(i, j int) bool { return strings.Compare(Bs[i], Bs[j]) == -1 })

	for idx := 0; idx < len(As); idx++ {
		if As[idx] != Bs[idx] {
			return false
		}
	}

	return true
}

func TestIsValidCountryCode(t *testing.T) {
	type test struct {
		country string
		want    bool
	}

	tests := []test{
		{"", false},
		{"  ", false},
		{"a", false},
		{"A", false},
		{"1", false},
		{"A1", false},
		{"1A", false},
		{"id", true},
		{"iD", true},
		{"Id", true},
		{"ID", true},
		{"IDN", false},
	}

	for _, test := range tests {
		got := IsValidCountryCode(test.country)
		if got != test.want {
			t.Errorf("%s | want: %v | got: %v", test.country, test.want, got)
		}
	}
}

func TestIsValidCountryCodeTotal(t *testing.T) {
	var count int
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			code := alphabet[i:i+1] + alphabet[j:j+1]
			if IsValidCountryCode(code) {
				count++
			}
		}
	}

	if count != 249 {
		t.Errorf("missing country code")
	}
}

func TestGetPhoneCodesFromCountryCode(t *testing.T) {
	type test struct {
		country string
		want    []string
		err     error
	}

	tests := []test{
		{"", nil, ErrorInvalidCountryCode},
		{"A", nil, ErrorInvalidCountryCode},
		{"AA", nil, ErrorInvalidCountryCode},
		{"AD", []string{"+376"}, nil},
		{"ID", []string{"+62"}, nil},
		{"CA", []string{"+1"}, nil},
		{"US", []string{"+1"}, nil},
		{"PR", []string{"+1"}, nil},
		{"JM", []string{"+1"}, nil},
		{"KZ", []string{"+7"}, nil},
		{"RU", []string{"+7"}, nil},
		{"BV", []string{}, nil},
		{"HM", []string{}, nil},
		{"UM", []string{}, nil},
		{"SH", []string{"+247", "+290"}, nil},
		{"VA", []string{"+39", "+379"}, nil},
	}

	for _, test := range tests {
		got, err := GetPhoneCodesFromCountryCode(test.country)
		if err != nil && test.err == nil {
			t.Errorf("%s | want err: %v | got err: %v", test.country, test.err, err)
		}
		if err == nil && test.err != nil {
			t.Errorf("%s | want err: %v | got err: %v", test.country, test.err, err)
		}
		if !checkEquality(got, test.want) {
			t.Errorf("%s | want: %v | got: %v", test.country, test.want, got)
		}
	}
}

func TestGetPhoneCodeFromPhoneNumber(t *testing.T) {
	type test struct {
		phone string
		want  string
		err   error
	}

	tests := []test{
		{"ABC", "", ErrorInvalidPhoneNumber},
		{"12345678", "", ErrorInvalidPhoneNumber},
		{"+1234", "", ErrorInvalidPhoneNumber},
		{"+0001234", "", ErrorUnsupportedPhoneNumber},
		{"+8001234", "", ErrorUnsupportedPhoneNumber},
		{"+1001234", "+1", nil},
		{"+7991234", "+7", nil},
		{"+6201234", "+62", nil},
		{"+6701234", "+670", nil},
	}

	for _, test := range tests {
		got, err := GetPhoneCodeFromPhoneNumber(test.phone)
		if err != nil && test.err == nil {
			t.Errorf("%s | want err: %v | got err: %v", test.phone, test.err, err)
		}
		if err == nil && test.err != nil {
			t.Errorf("%s | want err: %v | got err: %v", test.phone, test.err, err)
		}
		if got != test.want {
			t.Errorf("%s | want: %v | got: %v", test.phone, test.want, got)
		}
	}
}

func TestGetCountryCodesFromPhoneNumber(t *testing.T) {
	type test struct {
		phone string
		want  []string
		err   error
	}

	tests := []test{
		// TODO
	}

	for _, test := range tests {
		got, err := GetCountryCodesFromPhoneNumber(test.phone)
		if err != nil && test.err == nil {
			t.Errorf("%s | want err: %v | got err: %v", test.phone, test.err, err)
		}
		if err == nil && test.err != nil {
			t.Errorf("%s | want err: %v | got err: %v", test.phone, test.err, err)
		}
		if !checkEquality(got, test.want) {
			t.Errorf("%s | want: %v | got: %v", test.phone, test.want, got)
		}
	}
}
