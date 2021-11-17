package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cleanmachine1/capitalise"
)

/* Todo:

[x]Save as command-name.md

[x]implement a system for amount of command examples, either ask before or just when reaching end of each
[x]include an 8 cap on the amounts

[x]Use the correct syntax, eg. periods and backticks

[x]Create and append, deleting the old one with user's permission
*/

// Page creator for TLDR-pages, simplifying the syntax

func main() {
	fmt.Println("Enter the name of the program/command:")
	scanner := bufio.NewScanner(os.Stdin)                    // Create a text scanner
	scanner.Scan()                                           // Scan for the title
	title1 := scanner.Text()                                 // Collect this run of the scan and save
	title1 = strings.TrimSuffix(title1, " ")                 // Removes commonly applied extra space when entering values
	pagename := strings.ReplaceAll(title1, " ", "-") + ".md" // for creating the file nam
	if _, err := os.Stat(pagename); err == nil {
		fmt.Printf("file %q already exists, overwrite it? (y/N)", pagename)
		scanner.Scan()
		choice := scanner.Text()
		if choice == "y" || choice == "yes" || choice == "Yes" {
			os.Remove(pagename)
			os.Create(pagename)
		} else {
			fmt.Println("Exiting")
			os.Exit(1)
		}

	}
	title1 = "# " + title1 // For when writing to page

	fmt.Println("Enter a description for the program/command:")
	scanner.Scan()
	desc := scanner.Text()
	desc = strings.TrimSuffix(desc, " ")

	desc = "> " + desc + "."

	fmt.Println("Enter a more information link:")
	scanner.Scan()
	link := scanner.Text()
	link = strings.TrimSuffix(link, " ")
	link = "> More information: " + link + "."

	file, err := os.OpenFile(pagename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Close the file at the final command
	file.WriteString(title1 + "\n")
	file.WriteString("\n" + desc)
	file.WriteString("\n" + link + "\n")
	var i int
	for i = 1; i <= 8; i++ { // commands part of the page - allows 8
		fmt.Println("MAX 8 commands, enter nothing for exiting!\n")
		fmt.Println("Enter a description for a command example:")
		scanner.Scan()
		command_desc := scanner.Text()
		command_desc = strings.TrimSuffix(command_desc, " ") // Remove blankspace which the user could enter
		command_desc = capitalise.First(command_desc)
		if i < 8 {
			if command_desc == "" {
				fmt.Println("Saving and exiting")
				i = 9
			}

		}

		command_desc = "- " + command_desc + ":"

		file.WriteString("\n" + command_desc + "\n")

		fmt.Println("Now enter the command:")
		scanner.Scan()
		command := scanner.Text()
		command = strings.TrimSuffix(command, " ")

		command = "`" + command + "`"
		file.WriteString("\n" + command + "\n")

	}
}
