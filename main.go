package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-aybukeertekin/library"
)

func init() {
	//Create books
	library.CreateBookLibrary()
}

func main() {
	args := os.Args
	command, error := library.CheckUserArguments(args)
	//If user arguments are correct, apply commands.
	if error == nil {
		switch command {
		case library.SearchCommand:
			//If there are more than one word, handle
			wordsList := make([]string, len(args)-2)
			for i := 2; i < len(args); i++ {
				wordsList[i-2] = args[i]
			}
			words := strings.Join(wordsList, " ")
			library.PrintLineMessage("BOOKS:")
			library.SearchBooks(words)
		case library.ListCommand:
			library.PrintLineMessage("BOOKS:")
			library.ListBooks()
		case library.BuyCommand:
			//Arguments are already checked, no error handling needed here
			id, _ := strconv.Atoi(args[2])
			count, _ := strconv.Atoi(args[3])

			//Business logic error handling
			err := library.BuyBooks(id, count)
			if err != nil {
				library.PrintLineMessage(err.Error())
			} else {
				library.PrintLineMessage("Buy operation is successful!")
			}
		case library.DeleteCommand:
			//Arguments are already checked, no error handling needed here
			id, _ := strconv.Atoi(args[2])

			//Business logic error handling
			err := library.DeleteBook(id)
			if err != nil {
				library.PrintLineMessage(err.Error())
			} else {
				library.PrintLineMessage("Delete operation is successful!")
			}
		}
	}
}
