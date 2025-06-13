package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var contacts = make(map[string]Contact)

func updateContacts() {
	file, err := os.Open("contacts.json")
	checkErr(err)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	checkErr(err)

}
func saveContacts() {
	file, err := os.Create("contacts.json")
	checkErr(err)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // optional: for pretty formatting
	err = encoder.Encode(contacts)
	checkErr(err)

	fmt.Println("Contacts is Saved")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error")
	}
}

func main() {
	updateContacts()
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
			searchContact(scanner)
		case "4":
			deleteContact(scanner)
		case "5":
			fmt.Println("Exiting...")
			saveContacts()
			return

		default:
			fmt.Println("Please Select between 1 to 5")
		}

	}

}

func addContact(scanner *bufio.Scanner) {
	fmt.Println("\nEnter Name")
	scanner.Scan()
	name := strings.ToLower(scanner.Text())
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

func searchContact(scanner *bufio.Scanner) {
	fmt.Println("---------")
	fmt.Println("Which Contact you want to search (Enter Name):")
	scanner.Scan()

	noOfContacts := 0

	searchedName := strings.ToLower(scanner.Text())

	for key, contact := range contacts {
		if strings.Contains(key, searchedName) {
			fmt.Println("Name => ", contact.Name, " Phone Number => ", contact.Phone, " Email => ", contact.Email)
			noOfContacts++
		}
	}

	if noOfContacts == 0 {
		fmt.Println("No Contact Found")
	}

}

func deleteContact(scanner *bufio.Scanner) {
	fmt.Println("Enter Full Contact Name to Delete")
	scanner.Scan()
	deleteContactName := strings.ToLower(scanner.Text())

	contactData, ok := contacts[deleteContactName]
	if ok {
		fmt.Println("Are you really want to Delete : ", contactData.Name)
		fmt.Println("Enter yes | no ")
		scanner.Scan()
		yesOrNo := scanner.Text()

		if yesOrNo == "yes" {
			delete(contacts, deleteContactName)
			fmt.Println("Contact is deleted Successfully!!")
			return
		} else {
			return
		}
	} else {
		fmt.Println("No contact found")
		return
	}
}
