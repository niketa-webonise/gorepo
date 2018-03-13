package phnum

import ("testing")

var numberTests = []struct{  

description string
input string
number string
areaCode string
formatted string
expectErr bool
}	{

{
input:"(223) 456-7890",
number:"2234567890",
areaCode:"223",
formatted: "(223) 456-7890",

		expectErr:   false,
},
{

input:       "223.456.7890",
number:      "2234567890",
areaCode:    "223",
formatted:   "(223) 456-7890",
		expectErr:   false,
},
{

input:       "223 456 7890",
number:      "2234567890",
areaCode:    "223",
formatted:   "(223) 456-7890",
		expectErr:   false,
},


{

input:       "+1 (223) 456-7890",
number:      "2234567890",
areaCode:    "223",
formatted:   "(223) 456-7890",
		expectErr:   false,
},

{
input:       "223 456   7890   ",
number:      "2234567890",
areaCode:    "223",
formatted:   "(223) 456-7890",
		expectErr:   false,
},
	
	{
		description: "invalid if exchange code starts with 1",
		input:       "(223) 156-7890",
		expectErr:   true,
	},
	
	
}

func TestNumber(t *testing.T) {

	for _,test := range numberTests {

		actual, actualError := Number(test.input)
		
		if !test.expectErr {
			if actual != test.number {
				t.Errorf("Number(%s): expected %s, actual %s", test.input, test.number, actual)
			}
		} else if actualError == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nNumber(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}


func TestAreaCode(t *testing.T) {

	for _,test := range numberTests {

		actual, actualError := AreaCode(test.input)
		if !test.expectErr {
			if actual != test.areaCode {
				t.Errorf("AreaCode(%s): expected %s, actual %s", test.input, test.areaCode, actual)
			} 
		} else if actualError == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nAreaCode(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}


func TestFormatted(t *testing.T) {

	for _,test := range numberTests {

		actual, actualError := Format(test.input)
		if !test.expectErr {
			if actual != test.formatted {
				t.Errorf("AreaCode(%s): expected %s, actual %s", test.input, test.formatted, actual)
			}
		} else if actualError == nil {
			// if we expect an error and there isn't one
			t.Errorf("FAIL: %s\nAreaCode(%q): expected an error, but error is nil", test.description, test.input)
		}
	}
}