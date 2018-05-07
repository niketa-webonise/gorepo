package main

import (
	"fmt"
	
)

func main() {
	/*args := os.Args
	if len(args) > 1 {
		fmt.Println(args[1])
	} else {
		fmt.Println("hey !!!!!")
	}*/
	langs := []string{"Go","LANG","GO!!!"}
	fmt.Println(langs[0])
	fmt.Println(langs[1])
	fmt.Println(langs[2])
}
