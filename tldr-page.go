package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cleanmachine1/capitalise"
)

const ( // Usage for program's version and ANSI codes for text formatting
	VERSION     = "v1.4"
	COLORWHITE  = "\033[37m"
	COLORBLUE   = "\033[36m"
	COLORYELLOW = "\033[33m"
	UNDERLINE   = "\033[4m"
	NORMAL      = "\033[0m"
)

func Check_Empty(input string) { // Function used to check if whether a string entered is empty/whitespace
	input = strings.ReplaceAll(input, " ", "") // if the string given is "      ", its still blank
	if input == "" {
		fmt.Println("Exiting, invalid input!")
		os.Exit(1)
	}
}

func Reader() string { // Function for collecting user input easier as a string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func Remove_Punctuation(input string, choice int) string { // Function to fix errors regarding syntax
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
	double_desc_flag := flag.Bool("2", false, "Use 2 lines in the description") // format: parameter, default, description
	version_flag := flag.Bool("v", false, "Display version")
	// -h comes with using the flag package using the descriptions of each flag
	flag.Parse()

	if *version_flag {
		fmt.Println("TLDR-page version:", VERSION)
		os.Exit(0)
	}

	fmt.Println("Enter the name of the program/command:")
	title := Reader()                    // Uses bufio in a function to limit repeated code
	Check_Empty(title)                   // Check if title is whitespace/blank
	title = Remove_Punctuation(title, 1) // Removes the punctuation which the user could enter

	pagename := strings.ReplaceAll(title, " ", "-") + ".md" // for creating the file name

	// If the command entered is (for example) git push, the white space will become - so therefore git-push.md

	if _, err := os.Stat(pagename); err == nil { // Check if page exists before trying to overwrite it
		fmt.Printf(string(COLORYELLOW)+"file %q already exists, overwrite it? (y/N) "+string(COLORWHITE), pagename)
		choice := Reader()

		if choice == "y" || choice == "yes" || choice == "Yes" {
			os.Remove(pagename) // Delete the file, to be created later
		} else { // If the user input is no, then exit rather than continuing
			fmt.Println("Exiting")
			os.Exit(1)
		}

	}
	title = "# " + title // For when writing to page, for TLDR syntax

	fmt.Println("Enter a description for the program/command:")
	desc1 := Reader()
	Check_Empty(desc1)
	desc1 = Remove_Punctuation(desc1, 1)

	final_desc := "> " + capitalise.First(desc1) + "."

	if *double_desc_flag { // If the flag is raised then:
		fmt.Println("Enter a second description for the program/command:")
		desc2 := Reader()
		Check_Empty(desc2)
		desc2 = Remove_Punctuation(desc2, 1)

		// Change the desc variable to include desc2
		final_desc = final_desc + "\n> " + capitalise.First(desc2) + "."
	}

	fmt.Println("Enter a more information link:")
	link := Reader()
	Check_Empty(link)
	link = "> More information: <" + link + ">." // Formating

	file, err := os.OpenFile(pagename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // Open the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Close the file as the final command

	file.WriteString(title + "\n" + "\n" + final_desc + "\n" + link + "\n") // Write the title, desc, and link

	var i int // Assign empty variable
	fmt.Println(string(COLORBLUE) + "MAX 8 commands, to exit and save early, enter no input at any stage." + string(COLORWHITE))
	for i = 1; i <= 8; i++ { // commands part of the page - allows 8
		fmt.Printf("Command example %d/8\n", i)

		fmt.Println(string(COLORBLUE) + string(UNDERLINE) + "Part 1." + string(NORMAL) + " Enter a description for a command example:")
		command_desc := Reader()
		command_desc = Remove_Punctuation(command_desc, 1)

		command_desc = capitalise.First(command_desc)

		if command_desc == "" { // Break to end if empty
			// Wont use Check_Empty() since we don't want to exit
			break
		}

		command_desc = "- " + capitalise.First(command_desc) + ":"

		fmt.Println(string(COLORBLUE) + string(UNDERLINE) + "Part 2." + string(NORMAL) + " Now enter the corresponding command:") // Part 2
		command := Reader()
		command = Remove_Punctuation(command, 0) // with 0 because we need the backticks to be removed since users may enter them

		if command == "" { // Break to end
			break
		}

		command = "`" + command + "`"

		file.WriteString("\n" + command_desc + "\n" + "\n" + command + "\n") // Write to file
		fmt.Printf("\n")
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Saving to " + string(COLORBLUE) + path + "/" + pagename + " " + string(COLORWHITE) + "and exiting.\n\n")

	fmt.Println("Would you like to open the page in your default text editor?")
	further_edits_choice := Reader()
	if further_edits_choice == "y" || further_edits_choice == "yes" || further_edits_choice == "Yes" {
		command_string := "$EDITOR " + pagename
		cmd := exec.Command(`bash`, `-c`, command_string) // execute the command
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	fmt.Println("If you want to contribute this page to TLDR, please follow the instructions\nfrom the following link:")
	fmt.Print(string(COLORBLUE) + "https://github.com/tldr-pages/tldr#how-do-i-contribute\n" + string(COLORWHITE))
}
