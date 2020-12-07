package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}
type person struct {
	firstName string
	lastName string
	contactInfo // contactInfo contactInfo
}

func main() {
	// hieu := person{firstName: "Nguyen",lastName: "Hieu"}
	//-- print only values --
	// fmt.Println(hieu) 
	//-- print fields and values --
	// fmt.Printf("%+v",hieu) 
	hieu := person{
		firstName: "Hieu",
		lastName: "Nguyen",
		contactInfo: contactInfo {
			email: "hieutk5@gmail.com",
			zipcode: 550000, // must have comma end of line
		}
	}
	// fmt.Printf("%+v", hieu)
	hieuPointer := &hieu
	hieuPointer.updateName("an")
	hieu.print()
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
