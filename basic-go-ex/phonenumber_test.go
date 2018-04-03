package phonenumber

import "testing"

var test = []struct{
	input string
	expected string
	expectedErr bool
}{
	{
	"(123) 456-7890", "1234567890", false,
	},
	{
		"123.456.7890", "1234567890", false,
	},
	{
		"1234567890", "1234567890", false,
	},
	{
		"12345678901234567", "", true,
	},
	{
		"21234567890", "", true,

		},
		{
			"123456789", "", true,
		},
	}


func TestNumber(t *testing.T) {
	for _, test := range numberTests {
		actual, actualErr := Number(test.input)
		if actual != test.expected {
			t.Errorf("Number(%s): expected [%s], actual: [%s]", test.input, test.expected, actual)
		}
		// if we expect an error and there isn't one
		if test.expectErr && actualErr == nil {
			t.Errorf("Number(%s): expected an error, but error is nil", test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectErr && actualErr != nil {
			t.Errorf("Number(%s): expected no error, but error is: %s", test.input, actualErr)
		}
	}
}
