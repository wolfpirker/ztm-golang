//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Statistics struct {
	enteredCmnds int
	linesOfTextEntered int
}

//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
func reactToCommand(cmd string, stat *Statistics) bool{
	cmd = strings.ToUpper(cmd)
	switch cmd {
	case "HELLO":
		stat.enteredCmnds+=1
		stat.linesOfTextEntered+=1
		fmt.Println("Hello there!")
	case "BYE":
		stat.enteredCmnds+=1
		stat.linesOfTextEntered+=1
		fmt.Println("Bye, have a nice day ;)")
	case "Q":
		fmt.Println("Entered commands ", stat.enteredCmnds)
		fmt.Println("Entered lines ", stat.linesOfTextEntered)
		return true
	default:
		fmt.Println("invalid command")
		stat.linesOfTextEntered+=1
	}
	return false
}

func main() {
	//  Create an interactive command line application that supports arbitrary
	//  commands. When the user enters a command, the program will respond
	//  with a message. The program should keep track of how many commands
	//  have been ran, and how many lines of text was entered by the user.

	var stats Statistics

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if reactToCommand(scanner.Text(), &stats){
			return
		}
	}
	
}
