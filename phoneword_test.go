package main

import "testing"

func TestPhoneword(t *testing.T) {
	testcases := []struct{ word, expected string }{
		{
			" ",
			"0",
		},
		{
			"adgjmptw ",
			"234567890",
		},
		{
			" abcdefghijklmnopqrstuvwxyz",
			"022233344455566677778889999",
		},
		{
			"ADGJMPTW ",
			"234567890",
		},
		{
			" ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"022233344455566677778889999",
		},
		{
			"0123456789!@#$%^&*()",
			"0123456789**********",
		},
	}

	for i, tc := range testcases {
		phone := wordToNumber(tc.word)
		if phone != tc.expected {
			t.Errorf("%d: %s: %s != %s", i, tc.word, tc.expected, phone)
		}
	}
}
