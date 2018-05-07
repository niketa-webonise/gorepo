/*+1 (613)-995-0253 done
613-995-0253  done
1 613 995 0253 done
613.995.0253 done
(223) 456-7890 // done
223.456.7890 //done
223 456   7890  
123456789 //in-between
*/


package main

import ("fmt"
"strings"
"regexp")

var v1 string
var v2 string
var v3 string
var temp string
func Number() {

	fmt.Println(temp+v2+v3)
}

func AreaCode() {	
	fmt.Println(temp)
}
func Format() {

	temp1 := v1
	if !strings.Contains(temp1,"(") {
		temp1 = "("+temp1+")"
	}
	
	fmt.Println(temp1+" "+v2+"-"+v3)
}
func main() {

	input := "223 456   7890"
	//fmt.Println("length:",len(input))

	 
        re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
    re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
    final := re_leadclose_whtsp.ReplaceAllString(input, "")
    final = re_inside_whtsp.ReplaceAllString(final, " ")
        fmt.Println(final)

	sp := strings.Trim(input," ")
	split_str := strings.Split(sp," ")
		//fmt.Println("after:",split_str[0])
		//fmt.Println("after:",split_str[1])
		//fmt.Println("after:",split_str[2])
		//fmt.Println("after:",split_str[3])
		//fmt.Println("after:",split_str[4])



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

				/*	input:       "+1 (223) 456-7890",
					number:      "2234567890",
					areaCode:    "223",
					formatted:   "(223) 456-7890"*/
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

Number()
AreaCode()
Format()
}