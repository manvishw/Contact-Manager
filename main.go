package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

var contacts = make(map[string]Contact)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("CLI Contact Manager")
	for {
		fmt.Println("------------------------------")
		fmt.Println("|     1. Add Contact         |")
		fmt.Println("|     2. View Contacts       |")
		fmt.Println("|     3. Search Contact      |")
		fmt.Println("|     4. Delete Contact      |")
		fmt.Println("|     5. Exit                |")
		fmt.Println("------------------------------")

		fmt.Println("Enter your Choice:")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addContact(scanner)
		case "2":
			viewAllContacts(scanner)
		case "3":
			fmt.Println("Search Contact")
		case "4":
			fmt.Println("Delete Contact")
		case "5":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Please Select between 1 to 5")
		}

	}
}

func addContact(scanner *bufio.Scanner) {
	fmt.Println("\nEnter Name")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("\nEnter Phone Number")
	scanner.Scan()
	phone := scanner.Text()
	fmt.Println("\nEnter Your Email")
	scanner.Scan()
	email := scanner.Text()

	newContact := Contact{
		name, phone, email,
	}

	contacts[name] = newContact

	fmt.Println("\nContact Added Successfully : ", newContact, "\n\n")
}

func viewAllContacts(scanner *bufio.Scanner) {
	fmt.Println("---  ")
	fmt.Println("\n\nAvailabe Contacts\n")
	fmt.Println("---  ")
	for _, contact := range contacts {
		fmt.Println("Name => ", contact.Name, " Phone Number => ", contact.Phone, " Email => ", contact.Email)
	}
	fmt.Println("--- \n\n")

}
