package main

import("fmt"
"strings")

func main() {

	input_str := "9 8 7\n5 3 2\n6 6 7"

	rows := strings.Split(input_str,"\n")

	//fmt.Println("rows:",rows)//["9 8 7" "5 3 2" "6 6 7"]

	sl := make([][]string,len(rows))//[[] [] []]
	//fmt.Println("slice:",sl)

	for i,row := range rows {

		row_array := strings.Split(row," ")
		//fmt.Println("after spliting by spaces",row_array)

		sl[i] = row_array



		//fmt.Println("slice:",sl)//[[9 8 7][5 3 2][6 6 7]]
	}

	fmt.Println("row output")


	for outerindex,row_separated_value := range sl {

	for innerindex,_ := range row_separated_value {
	if innerindex+1 == len(row_separated_value) {
	fmt.Printf("%s",sl[outerindex][innerindex])
	} else {
	fmt.Printf("%s,",sl[outerindex][innerindex])
	}
	//fmt.Printf("%s,",sl[outerindex][index])
	}
	fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("column output")

	for outerindex,col_separated_value := range sl {

		for innerindex,_ := range col_separated_value {

			if innerindex+1 == len(col_separated_value) {

				fmt.Printf("%s",sl[innerindex][outerindex])
			} else {

				fmt.Printf("%s,",sl[innerindex][outerindex])
			}
		}
		fmt.Println()
	}

	
}