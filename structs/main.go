package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// kingkong := person{firstName: "King", lastName: "Kong"}

	// var queenkong person

	// queenkong.firstName = "Queen"
	// queenkong.lastName = "Kong"

	// fmt.Println(queenkong)

	// fmt.Println(kingkong.firstName)
	// fmt.Println(kingkong)

	jim := person{
		firstName: "Jim",
		lastName:  "Kingkong",
		contactInfo: contactInfo{
			email:   "jungle@yo.com",
			zipCode: 2223232,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("NewName")

	jim.print()

}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
