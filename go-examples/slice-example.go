package main

import "fmt"

func main() {

slice_var:=make([]string,3)
fmt.Println("slice variable",slice_var)

slice_var[0]="b"
slice_var[1]="a"
slice_var[2]="10"
fmt.Println(slice_var)
fmt.Println("length:",len(slice_var));
slice_var=append(slice_var,"20","30","c","d")
fmt.Println(slice_var);

copy_var:=make([]string,len(slice_var))
copy(copy_var,slice_var)
fmt.Println("copy of slice variable",copy_var)

slicing:=slice_var[:3]
fmt.Println(slicing)
}
