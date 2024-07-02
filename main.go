package main

import (
	"fmt"

)

/*
	Description:
	Create a simple phonebook application that stores and retrieves phone numbers using a map. The application should allow adding, removing, and searching for contacts by name.

	Tasks:

	Implement a menu-driven interface with options to add, remove, and search contacts.
	Use a map where keys are contact names and values are phone numbers.
	Ensure the program can handle the following operations:
	Add a new contact
	Remove an existing contact
	Search for a contact by name
	Display all contacts
*/



func main() {

	fmt.Println("1 - To add a contact")
	fmt.Println("2 - To remove an existing contact")
	fmt.Println("3 - To search a contact by name")
	fmt.Println("4 - To display all contacts")

	var contacts map[string]string
	
	var option int
	fmt.Println("Enter a number from the menu:")
	fmt.Scan(&option)

	switch option {
	case 1: 

		var name map[string]string
		fmt.Println("Enter the contacts name:")
		fmt.Scan(&name)

		var number map[int]int
		fmt.Println("Enter the contacts number:")
		fmt.Scan(&number)



	case 2: 

		fmt.Println("Enter the name of the existing contact to remove:")	
		var remove string 
		fmt.Scan(&remove)

		delete(contacts,remove)
	
	case 3: 
	
		fmt.Println("Enter the name of the contact you want to search:")
		var search string 
		fmt.Scan(&search)



		





	}
	

}