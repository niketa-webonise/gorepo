package main

import "fmt"

type company struct {
name string
address string
}

func main() {
fmt.Println(company{name:"ABC,DEF",address:"XYZ address,DEF address"})
fmt.Println(company{name:"PTC"})

company_var := company{name:"webonise",address:"bhavdhan"}
fmt.Println("accessing company name:",company_var.name)
fmt.Println("accessing company address:",company_var.address)

}
