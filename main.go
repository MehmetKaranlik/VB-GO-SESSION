package main

import "VB-GO-SESSION/person"

func main() {
	processPersonInterface(person.Member{})
	processPersonInterface(person.Person{})

}

func processPersonInterface(p person.PersonInterface) {
	p.GetName()
}
