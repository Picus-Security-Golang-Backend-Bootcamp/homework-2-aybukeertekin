# homework-2-aybukeertekin

Simple Library Application in Go Language 

* Library is a collection of Book Structs. 
* Book implements Deletable interface. 

Commands are: 

Usage:

* go run main.go [arguments]

The commands are:

* list
* search [book-name or ISBN or stock-code]
* buy [ID count]
* delete [ID]

Explanations:

- List command lists all available (not deleted) books in the library.
- Search command search a book with given book-name, author-name, ISBN number or stock code.
- Buy command decreases stock count of a book by given count. Returns error if the book is not found or stock is not sufficient for buy operation.
- Delete command deletes a book with given ID. Returns error if the book is not found. 
