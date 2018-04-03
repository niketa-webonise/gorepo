package main
import (
 		"strings" 
 		"fmt"
 		)

func main() {
row:="   7   8  3   "
S:=strings.Split(strings.TrimSpace(row)," ")
	fmt.Println(S[0])
	
}
