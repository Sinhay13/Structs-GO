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

	// first way to use pointer
	// jimPointer := &jim             // pointer style, "&" to give address to the memory localization
	// jimPointer.updateName("Jimmy") // pointer style
	// jim.print()

	//shortcut pointer second way
	jim.updateName("jimmy")
	jim.print()

}

// func (p person) updateName (newFirstName string) { // to update Name
// 	p.firstName=newFirstName
// }

// update Name function pointer style
func (pointerToPerson *person) updateName(newFirstName string) { //"*person" this is description it means we are working with a pointer to a person
	(*pointerToPerson).firstName = newFirstName // "*pointerToPerson" this is an operator - it means we want to manipulate the value the pointer is referencing
}

// In go we can not change a value by the classic way : because only the copy will be modified not the original it is for that we need to use pointer !
// For updating : int, float, string, bool, structs

func (p person) print() { // to print
	fmt.Printf("%+v", p)
}
