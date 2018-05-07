package main

import "github.com/niketa/gorepo/interface-example/pets"

type FourLegged interface{
	Walk()
	Sit()
	Fetch()
}

func Demo(animal FourLegged){
	animal.Walk()
	animal.Sit()
	animal.Fetch()
}

func main(){

	dog := pets.Dog{"Fido", "Terrier"}
	cat := pets.Cat{"Fluffy", "Siamese"}

	Demo(dog)
	Demo(cat)
}