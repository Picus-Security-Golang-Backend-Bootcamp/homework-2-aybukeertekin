package library

import (
	"strings"
)

var Books []*Book = make([]*Book, len(BookNames)) //book list

//This functions prints book by checking their availability
func PrintBookList(bookList []*Book) {
	for index := 0; index < len(bookList); index++ {
		if IsAvailable(index + 1) {
			PrintBook(bookList[index])
		}
	}
}

//This function searchs book list and look if a book exists
//with the given book name, isbn, or stock code and prints them.
//Prints error message if no book exists.
func SearchBooks(searchWord string) {
	aBookFound := false
	searchWord = strings.ToLower(searchWord)
	for index := 0; index < len(Books); index++ {
		if IsAvailable(index + 1) {
			bookName := Books[index].BookName
			found := false
			if strings.Contains(strings.ToLower(bookName), searchWord) {
				PrintBook(Books[index])
				aBookFound = true
				found = true
			}
			authorName := Books[index].Author.Name
			if !found && strings.Contains(strings.ToLower(authorName), searchWord) {
				PrintBook(Books[index])
				aBookFound = true
				found = true
			}
			stockCode := Books[index].StockCode
			if !found && strings.Contains(strings.ToLower(stockCode), searchWord) {
				PrintBook(Books[index])
				aBookFound = true
			}
		}
	}
	if !aBookFound {
		PrintLineMessage("\tNo books are found with given argument.")
	}
}

//This function prints book list.
func ListBooks() {
	PrintBookList(Books)
}
