package main

import "fmt"

type person struct { // initiate like a dictionary in python
	firstName string
	lastName  string
	//contact contactInfo // we add contact
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	//alex := person{"Alex", "Anderson"} // first way if we respect the order of person
	//alex:= person{firstName:"Alex",lastName:"Anderson"}// second way

	// var alex person // third way
	// alex.firstName="Alex"
	// alex.lastName="Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex) // %+v to show all fields

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jim.updateName("Jimmy")
	jimPointer := &jim             // pointer style
	jimPointer.updateName("Jimmy") // pointer style
	jim.print()
}

// func (p person) updateName (newFirstName string) { // to update Name
// 	p.firstName=newFirstName
// }

// update Name function pointer style
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() { // to print
	fmt.Printf("%+v", p)
}
