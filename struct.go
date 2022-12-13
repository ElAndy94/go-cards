package main

import "fmt"

type contactInfo struct {
	email    string
	postCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func structs() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:    "andrewpeliza",
			postCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// p.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
