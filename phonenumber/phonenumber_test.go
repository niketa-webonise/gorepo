package phonenumber

import "testing"

var numberTests = []struct{
	input string
	expected string
	expectErr bool
}{
	{
	"(123) 456-7890","1234567890", false,
	},
	{
		"123.456.7890","1234567890", false,
	},
	{
		"1234567890","1234567890", false,
	},
	{
		"12345678901234567","", true,
	},
	{
		"21234567890","", true,

		},
		{
			"123456789","", true,
		},
	}


func TestNumber(t *testing.T) {
	for _, test := range numberTests {
		actual := Number(test.input)
		if actual != test.expected {
			t.Errorf("Number(%s): expected [%s], actual: [%s]", test.input, test.expected, actual)
		}
		
	}
}
