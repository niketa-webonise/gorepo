package main

import "fmt"
import "math"

type calculation interface{
	area() float32
	perimeter() float32
}

type circle struct{
       radius float32
}

type rect struct{
       length,breath float32
}

func (c circle) area() float32  {
	return math.Pi*c.radius*c.radius
}

func (c circle) perimeter() float32 {
        return 2*math.Pi*c.radius
}

func (r rect) area() float32 {
return r.length*r.breath
}


func (r rect) perimeter() float32  {
return 2*(r.length+r.breath)
}


func docal(c calculation) {
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.perimeter())
}

func main() {
	r:=rect{length:10,breath:20}
	c:=circle{radius:5}

	docal(r)
	docal(c)
}

