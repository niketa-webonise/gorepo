package phnum

import ("strings"
		"fmt"
		"regexp"
		"errors"
		)

var v1 string
var v2 string
var v3 string
var temp string

func Number(input string) (string, error) {




   re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
   re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)

   final := re_leadclose_whtsp.ReplaceAllString(input, "")
   
   final = re_inside_whtsp.ReplaceAllString(final, " ")



   err := validate(final)


	
	if(err == nil) {
		return (temp+v2+v3), nil
	} else {
		return "-1", err
	}
}


func Format(input string) (string, error) {

	err := validate(input)
	temp1 := v1
	
	if !strings.Contains(temp1,"(") {
		temp1 = "("+temp1+")"
	}

	if(err == nil) {
		return (temp1+" "+v2+"-"+v3), nil
	} else {
		return "-1", err
	}
}


func AreaCode(input string) (string, error) {

	err := validate(input)

	if(err == nil) {
		return temp, nil
	} else {
		return "-1", err
	}
}

func validate(input string) (error) {

	if strings.Contains(input," "){

	split_str := strings.Split(input," ")

	len_split_str := len(split_str)

	switch(len_split_str){

	case 1:	


		split_convert := strings.Replace(input,".","-",-1)
		split_dash := strings.Split(split_convert,"-")
		temp = split_dash[0]
		v2 = split_dash[1]
		v3 = split_dash[2]


	case 2:

		str1 :=split_str[0]
		split_dash := strings.Split(split_str[1],"-")
		if len(split_dash) == 3 {

				v1 = split_dash[0]
				fmt.Println("v1",v1)
				v2 = split_dash[1]
				v3 = split_dash[2]

			}else if len(split_dash) == 2{

					v1 = str1
					v2 = split_dash[0]
					v3 = split_dash[1]


			}
				temp = strings.Replace(v1,"(","",-1)
				temp = strings.Replace(temp,")","",-1)
		//fmt.Println("temp",temp)
	case 3:

	
		str1 := split_str[0]
		if strings.HasPrefix(str1,"+1") && strings.Contains("+1 (223) 456-7890","+1") {


			split_dash := strings.Split(split_str[2],"-")

				v1 = split_str[1]
				v2 = split_dash[0]
				v3 = split_dash[1]
			}

			temp = strings.Replace(v1,"(","",-1)
			temp = strings.Replace(temp,")","",-1)

	case 4:

			temp = split_str[1]
			v2 = split_str[2]
			v3 = split_str[3]

		}
	
	}else{
		  if len(input) > 10 && !strings.HasPrefix(input,"1") && 
		  !strings.Contains(input,".") && !strings.Contains(input,"-") && len(input)==9{


		  	return errors.New("invalid code")
		  }
	}

	if string(v2[0]) == "1" {

		return errors.New("exchange code  starts with 1!")

	}else if string(v2[0]) == "0" {

			return errors.New("exhange code starts with 0!")

	}else if string(temp[0]) == "1"  {

		return errors.New("area code starts with 1")

	}else if string(temp[0]) == "0" {

		return errors.New("area code starts with 0")

	}else if strings.Contains(v2,"@:!#$%^&*"){

		return errors.New("contains punctuations")
		
	}else if strings.Contains(v2,"abcdefghijklmnopqrstuuvwxyz"){

		return errors.New("contains letters")
	}
	return nil
}