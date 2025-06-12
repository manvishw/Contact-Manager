package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("CLI Contact Manager")
	for {

		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")
		fmt.Println("Enter your Choice:")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Add Contact")
		case "2":
			fmt.Println("View Contacts")
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
