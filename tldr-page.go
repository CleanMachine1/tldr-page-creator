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

/* Variables to track:
title1
pagename
desc
link
file
command
command_desc
path
i
*/

const ( // Usage for changing the color of text
	colorWhite = "\033[37m"
	colorBlue  = "\033[36m"
)

func removesuffix(input string) string { // Function used for removing trailing whitespace
	temp := strings.TrimSpace(input)
	return temp
}
func remove_punctuation(input string) string { // Function to fix errors regarding syntax
	temp := strings.Trim(input, ".:`-># ")
	/* This function achieves the ability for the user to be able enter punctuation.
	For example, if the user enters "> Version control system."
	This function will still allow the page to have the correct punctuation without duplicating it
	*/
	return temp
}
func reader() string { // Function for collecting user input easier
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}
func checkempty(input string) { // Function used to check if whether a string entered is empty/whitespace
	input = strings.ReplaceAll(input, " ", "") // if the string given is "      ", its still blank
	if input == "" {
		fmt.Println("Exiting, invalid input!")
		os.Exit(1)
	}
}
func main() {
	flagvar := flag.Bool("2", false, "Use 2 lines in the description")
	flag.Parse()

	fmt.Println("Enter the name of the program/command:")
	title1 := reader()                  // Uses bufio in a function to limit repeated code
	checkempty(title1)                  // Check if title1 is whitespace/blank
	title1 = removesuffix(title1)       // Remove blankspace which the user could enter
	title1 = remove_punctuation(title1) // Removes the punctuation which the user could enter

	pagename := strings.ReplaceAll(title1, " ", "-") + ".md" // for creating the file name

	// If the command entered is (for example) git push, the white space will become - so therefore git-push.md

	if _, err := os.Stat(pagename); err == nil { // Check if page exists before trying to overwrite it
		fmt.Printf("file %q already exists, overwrite it? (y/N)", pagename)
		choice := reader()
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
	desc1 := reader()
	checkempty(desc1)
	desc1 = removesuffix(desc1) // Remove blankspace which the user could enter
	desc1 = remove_punctuation(desc1)

	desc := "> " + capitalise.First(desc1) + "."

	if *flagvar { // If the flag is raised then:
		fmt.Println("Enter a second description for the program/command")
		desc2 := reader()
		checkempty(desc2)
		desc2 = removesuffix(desc2)
		desc2 = remove_punctuation(desc2)

		// Change the desc variable to include desc2
		desc = desc + "\n> " + capitalise.First(desc2) + "."
	}

	fmt.Println("Enter a more information link:")
	link := reader()
	link = removesuffix(link)
	checkempty(link)
	link = "> More information: <" + link + ">." // Formating

	file, err := os.OpenFile(pagename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // Open the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Close the file as the final command

	file.WriteString(title1 + "\n" + "\n" + desc + "\n" + link + "\n") // Write the title, desc, and link

	var i int // Assign empty variable
	fmt.Println(string(colorBlue), "MAX 8 commands, enter nothing for saving and exiting!", string(colorWhite))
	for i = 1; i <= 8; i++ { // commands part of the page - allows 8
		fmt.Printf("Command %d/8\n", i)

		fmt.Println(" Part 1. Enter a description for a command example:")
		command_desc := reader()
		command_desc = removesuffix(command_desc) // Remove blankspace which the user could enter
		command_desc = remove_punctuation(command_desc)

		command_desc = capitalise.First(command_desc)

		if command_desc == "" { // Break to end if empty
			break
		}

		command_desc = "- " + capitalise.First(command_desc) + ":"

		fmt.Println(" Part 2. Now enter the corresponding command:") // Part 2
		command := reader()
		command = removesuffix(command)
		command = remove_punctuation(command)

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
	fmt.Print("Saving to ", string(colorBlue))
	fmt.Print(path + "/" + pagename ,string(colorWhite))
	fmt.Print(" and exiting.\n")
	fmt.Println("If you want to contribute this page to TLDR, please follow the instructions ")
	fmt.Println("from the following link:")
	fmt.Print(string(colorBlue), "https://github.com/tldr-pages/tldr#how-do-i-contribute\n", string(colorWhite))
}
