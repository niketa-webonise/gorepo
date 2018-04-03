package phonenumber

import "errors"

func Number(input string) (string,error) {

	digits := ""

	for _,i := range input {

		if '0'<= i && i<='9' {

			digits = digits + string(i)
		}
	}

	switch {

	case len(digits) == 10 :
		return digits,nil

	case len(digits) == 11 && digits[1] == '1':
		return digits[1:],nil


	default:
		return "",errors.New("not a phone number"+ input)
	}
}


