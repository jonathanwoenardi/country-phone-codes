package country_phone_codes

import (
	"testing"
)

func TestIsValidPhoneNumber(t *testing.T) {
	type test struct {
		phone string
		want  bool
	}

	tests := []test{
		{"", false},
		{"phone", false},
		{"testing@email.com", false},
		{"0123456789", false},
		{"+1-234-5678", false},
		{"+1 234 5678", false},
		{"(1)2345678", false},
		{"+123456", false},
		{"+1234567", true},
		{"+123456789012345", true},
		{"+1234567890123456", false},
		{"+00000000", true},
	}

	for _, test := range tests {
		got := isValidPhoneNumber(test.phone)
		if got != test.want {
			t.Errorf("%s | want: %v | got: %v", test.phone, test.want, got)
		}
	}
}
