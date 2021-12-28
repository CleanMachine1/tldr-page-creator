package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cleanmachine1/capitalise"
)

const ( // Usage for changing the color of text
	version     = "v1.3"
	colorWhite  = "\033[37m"
	colorBlue   = "\033[36m"
	colorYellow = "\033[33m"
)

func checkempty(input string) { // Function used to check if whether a string entered is empty/whitespace
	input = strings.ReplaceAll(input, " ", "") // if the string given is "      ", its still blank
	if input == "" {
		fmt.Println("Exiting, invalid input!")
		os.Exit(1)
	}
}

func reader() string { // Function for collecting user input easier as a string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func remove_punctuation(input string, choice int) string { // Function to fix errors regarding syntax
	switch choice {
	case 1:
		temp := strings.Trim(input, ".:-># ") // without backticks
		return temp

	default:
		temp := strings.Trim(input, ".:`-># ")
		return temp

	}

	// This sanitizes the users in attempt to match TLDR's syntax
}

func main() {
	doubledescflag := flag.Bool("2", false, "Use 2 lines in the description") // format: parameter, default, description
	versionflag := flag.Bool("v", false, "Display version")
	// -h comes with using the flag package using the descriptions of each flag
	flag.Parse()

	if *versionflag {
		fmt.Println("TLDR-page version:", version)
		os.Exit(0)
	}

	fmt.Println("Enter the name of the program/command:")
	title1 := reader()                     // Uses bufio in a function to limit repeated code
	checkempty(title1)                     // Check if title1 is whitespace/blank
	title1 = remove_punctuation(title1, 1) // Removes the punctuation which the user could enter

	pagename := strings.ReplaceAll(title1, " ", "-") + ".md" // for creating the file name

	// If the command entered is (for example) git push, the white space will become - so therefore git-push.md

	if _, err := os.Stat(pagename); err == nil { // Check if page exists before trying to overwrite it
		fmt.Print(string(colorYellow))
		fmt.Printf("file %q already exists, overwrite it? (y/N) ", pagename)
		fmt.Print(string(colorWhite))
		choice := reader()

		if choice == "y" || choice == "yes" || choice == "Yes" {
			os.Remove(pagename) // Delete the file, to be created later
		} else { // If the user input is no, then exit rather than continuing
			fmt.Println("Exiting")
			os.Exit(1)
		}

	}
	title1 = "# " + title1 // For when writing to page, for TLDR syntax

	fmt.Println("Enter a description for the program/command:")
	desc1 := reader()
	checkempty(desc1)
	desc1 = remove_punctuation(desc1, 1)

	desc := "> " + capitalise.First(desc1) + "."

	if *doubledescflag { // If the flag is raised then:
		fmt.Println("Enter a second description for the program/command:")
		desc2 := reader()
		checkempty(desc2)
		desc2 = remove_punctuation(desc2, 1)

		// Change the desc variable to include desc2
		desc = desc + "\n> " + capitalise.First(desc2) + "."
	}

	fmt.Println("Enter a more information link:")
	link := reader()
	checkempty(link)
	link = "> More information: <" + link + ">." // Formating

	file, err := os.OpenFile(pagename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // Open the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Close the file as the final command

	file.WriteString(title1 + "\n" + "\n" + desc + "\n" + link + "\n") // Write the title, desc, and link

	var i int // Assign empty variable
	fmt.Println(string(colorBlue), "MAX 8 commands, to exit and save early, enter no input.", string(colorWhite))
	for i = 1; i <= 8; i++ { // commands part of the page - allows 8
		fmt.Printf("Command %d/8\n", i)

		fmt.Println(" Part 1. Enter a description for a command example:")
		command_desc := reader()
		command_desc = remove_punctuation(command_desc, 1)

		command_desc = capitalise.First(command_desc)

		if command_desc == "" { // Break to end if empty
			// Wont use checkempty() since we don't want to exit
			break
		}

		command_desc = "- " + capitalise.First(command_desc) + ":"

		fmt.Println(" Part 2. Now enter the corresponding command:") // Part 2
		command := reader()
		command = remove_punctuation(command, 0) // with 0 because we need the backticks to be removed since users may enter them

		if command == "" { // Break to end
			break
		}

		command = "`" + command + "`"

		file.WriteString("\n" + command_desc + "\n" + "\n" + command + "\n") // Write to file

	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Saving to ", string(colorBlue), path+"/"+pagename, string(colorWhite), " and exiting.\n")
	fmt.Println("If you want to contribute this page to TLDR, please follow the instructions\nfrom the following link:")
	fmt.Print(string(colorBlue), "https://github.com/tldr-pages/tldr#how-do-i-contribute\n", string(colorWhite))
}
