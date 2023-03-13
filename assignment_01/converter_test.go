package assignment01_test

import (
	"tesfayprep/assignment_01/converter"
	"testing"
)

func TestRomanConverter(t *testing.T) {
	testCases := []struct {
		Name  string
		Input int
		Want  string
	}{
		{"1 gets coverted to I", 1, "I"},
		{"2 gets coverted to II", 2, "II"},
		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V (can't repeat more than 3 times)", 5, "V"},
		{"7 gets converted to VII (can't repeat more than 3 times)", 7, "VII"},
		{"9 gets converted to IX (can't repeat more than 3 times)", 9, "IX"},
		{"10 gets converted to X (can't repeat more than 3 times)", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3999 gets converted to MMMCMXCIX", 3999, "MMMCMXCIX"},
		{"2014 gets converted to MMXIV", 2014, "MMXIV"},
		{"1006 gets converted to MVI", 1006, "MVI"},
		{"798 gets converted to DCCXCVIII", 798, "DCCXCVIII"},
	}
	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			got := converter.ConvertToRoman(test.Input)

			if got != test.Want {
				t.Errorf("Want'%s' got '%s'", test.Want, got)
			}
		})
	}
}
