package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// alex.firstName = "Alex"
	// alex.lastName = "Mitra"

	jim := person{
		firstName: "Jim",
		lastName:  "Carry",
		contactInfo: contactInfo{
			email:   "jim@carry.com",
			zipcode: 9000,
		},
	}

	jim.updateName("Tim")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
