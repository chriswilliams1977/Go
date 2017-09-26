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

	jim.print()
	jim.updateName("chris");
	jim.print()
}

//reciever means you can call this function on any type of person so person.print
func (p person) print()  {
	fmt.Printf("%+v",p)
}

func (p person) updateName(newFirstName string)  {
	p.firstName = newFirstName
}
