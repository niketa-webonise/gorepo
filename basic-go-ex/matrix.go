package main

import("fmt"
"strings")

func main() {

	//input_str := "9 8 7\n5 3 2\n6 6 7"
	//input_str := "1 2\n10 20"
	//input_str := "9 7\n8 6"
	input_str := "9 8 7\n19 18 17"
	//input_str := "1 4 9\n16 25 36"
	//input_str := "1 2 3\n4 5 6\n7 8 9\n8 7 6"
	//input_str := "89 1903 3\n18 3 1\n9 4 800"
	//input_str := "1 2 3"
	//input_str := "1\n2\n3"
	//input_str := "\n3 4\n5 6"

	rows := strings.Split(input_str,"\n")

	fmt.Println("rows:",rows)//["9 8 7" "5 3 2" "6 6 7"]

	sl := make([][]string,len(rows))//[[] [] []]
	fmt.Println("slice:",sl)

	for i,row := range rows {

		row_array := strings.Split(row," ")
		fmt.Println("after spliting by spaces",row_array)

		sl[i] = row_array



		fmt.Println("slice:",sl)//[[9 8 7][5 3 2][6 6 7]]
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