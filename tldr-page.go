package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cleanmachine1/capitalise"
)

const ( // Usage for changing the color of text
	colorRed   = "\033[31m"
	colorWhite = "\033[37m"
)

func checkempty(input string) { // Function used to check if whether a string entered is empty/whitespace
	input = strings.ReplaceAll(input, " ", "") // if the string given is "      ", its still blank
	if input == "" {
		fmt.Println("Exiting, invalid input!")
		os.Exit(1)
	}
}
func main() {
	fmt.Println("Enter the name of the program/command:")

	scanner := bufio.NewScanner(os.Stdin) // Create a text scanner
	scanner.Scan()                        // Scan for the title
	title1 := scanner.Text()              // Collect this run of the scan and save

	checkempty(title1) // Check if title1 is whitespace/blank

	title1 = strings.TrimSuffix(title1, " ")                 // Removes commonly applied extra space when entering values
	pagename := strings.ReplaceAll(title1, " ", "-") + ".md" // for creating the file name
	// If the command entered is (for example) git push, the white space will become - so therefore git-push.md

	if _, err := os.Stat(pagename); err == nil { // Check if page exists before trying to overwrite it
		fmt.Printf("file %q already exists, overwrite it? (y/N)", pagename)
		scanner.Scan()
		choice := scanner.Text()
		if choice == "y" || choice == "yes" || choice == "Yes" {
			os.Remove(pagename) // Delete the file
			os.Create(pagename) // Recreate it, blank and empty
		} else { // If the user input is no, then exit rather than continuing
			fmt.Println("Exiting")
			os.Exit(1)
		}

	}
	title1 = "# " + title1 // For when writing to page, for TLDR syntax

	fmt.Println("Enter a description for the program/command:")
	scanner.Scan()
	desc := scanner.Text()
	desc = strings.TrimSuffix(desc, " ") // Remove blank space error
	checkempty(desc)
	desc = "> " + capitalise.First(desc) + "."

	fmt.Println("Enter a more information link:")
	scanner.Scan()
	link := scanner.Text()
	link = strings.TrimSuffix(link, " ")
	checkempty(link)
	link = "> More information: <" + link + ">." // Formating

	file, err := os.OpenFile(pagename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // Open the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Close the file as the final command

	file.WriteString(title1 + "\n" + "\n" + desc + "\n" + link + "\n") // Write the title, desc, and link

	var i int // Assign empty variable
	fmt.Println(string(colorRed), "MAX 8 commands, enter nothing for saving and exiting!", string(colorWhite))
	for i = 1; i <= 8; i++ { // commands part of the page - allows 8
		fmt.Println(" 1. Enter a description for a command example:")
		scanner.Scan()
		command_desc := scanner.Text()
		command_desc = strings.TrimSuffix(command_desc, " ") // Remove blankspace which the user could enter
		command_desc = capitalise.First(command_desc)

		if command_desc == "" { // Break to end
			break
		}

		command_desc = "- " + capitalise.First(command_desc) + ":"

		fmt.Println(" 2. Now enter the corresponding command:") // Part 2
		scanner.Scan()
		command := scanner.Text()
		command = strings.TrimSuffix(command, " ")

		if command == "" { // Break to end
			break
		}
		command = "`" + command + "`"

		file.WriteString("\n" + command_desc + "\n" + "\n" + command + "\n") // Write to file

	}
	fmt.Println("Saving and exiting.")
	fmt.Println("\nIf you want to contribute this page to TLDR, please use the following link:")
	fmt.Println("https://github.com/tldr-pages/tldr")
}
