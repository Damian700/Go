package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo //defines a property with the same name as type
}

func main() {
	var pio person // ---> defines zero values
	//pio := person{"Pio","Di Pietrelcina"} assigned in order
	//pio := person{firstName: "Pio", lastName: "Di Pietrelcina"}
	pio.firstName = "Pio"
	pio.lastName = "Di Pietrelcina"
	pio.contactInfo.zipCode = 82020
	fmt.Printf("%+v", pio)
	juan := person{
		firstName: "Juan",
		lastName:  "MVG",
		contactInfo: contactInfo{
			email:   "juan.mvg@gmail.com",
			zipCode: 1001,
		},
	}
	fmt.Printf("%+v", juan)

	/* 	pioPointer := &pio
	   	pioPointer.updateName("Saint Pio")
	   	pio.print() */
	pio.updateName("Saint Pio")
	pio.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Println(
		p.firstName,
		p.lastName,
		p.contactInfo.zipCode,
	)
}
