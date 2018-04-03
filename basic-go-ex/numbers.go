/*+1 (613)-995-0253
613-995-0253
1 613 995 0253
613.995.0253*/
package main

import ("fmt"
		"strings"
		"strconv")
var value1 int 
var value2 int
var value3 int
func main() {

input_str := "+1 (613)-995-0253"

split_str := strings.Split(input_str," ")

fmt.Println("split string",split_str[1])

len_str := len(split_str)

switch(len_str){

	case 1://

	case 2:
			split_dash := strings.Split(split_str[1],"-")
			value1 = split_dash[0]
			value2 = split_dash[1]
			value3 = split_dash[2]

			convert_str1,_ := strconv.Atoi(value1)

			convert_str2,_ := strconv.Atoi(value2)

			convert_str3,_ := strconv.Atoi(value3)

			fmt.Println(convert_str1+convert_str2+convert_str3)


}

}