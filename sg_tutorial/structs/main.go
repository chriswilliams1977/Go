package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//you can also specify field name like contact: contactInfo
	contactInfo
}

func main() {
	//chris := person{firstName: "Chris", lastName: "Williams"}
	//fmt.Println(chris.firstName)
	//fmt.Println(chris.lastName)

	//var chris person
	//chris.firstName = "Chris"
	//chris.lastName = "Williams"
	//fmt.Println(chris)
	//fmt.Printf("%+v",chris)

	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo : contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	//& is an operator
	//States give me access to the memory address jim is at (jimPointer is pointing at the memory address of jim
	jimPointer := &jim
	jimPointer.updateName("chris");
	jim.print()
}

//reciever means you can call this function on any type of person so person.print
func (p person) print()  {
	fmt.Printf("%+v",p)
}

//* in front of type ir different to * in front of a pointer
//whenever you have * then type you are saying you are looking for a pointer with type person
//function can only be called from a receiver of type of pointer to a person
func (pointerToPerson *person) updateName(newFirstName string)  {
	//* operator here get value at that memory address, get me value of pointerToPerson
	(*pointerToPerson).firstName = newFirstName
}
