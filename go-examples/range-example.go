package main

import "fmt"

func main() {
array1:=[]int{2,3,4}
fmt.Println(array1)
sum:=0
for _, num:=range array1 {
	sum=sum+num
}
fmt.Println(sum)

for i,num:=range array1 {
if num==3{
fmt.Println("index:",i)}
}

kvp := map[string]string{"key1":"value1","key2":"value2","key3":"value3"}
for k,v:=range kvp {
fmt.Println(" %s -> %s \n",k,v)
}

for k:= range kvp {
fmt.Println("key:",k)
}


}
