package library

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const SearchCommand = "search"
const ListCommand = "list"
const BuyCommand = "buy"
const DeleteCommand = "delete"

var ErrBookNotFoundWithId = errors.New("A book cannot be found with the given id!")
var ErrStockIsNotEnough = errors.New("Requested number of books is not available in the stock!")

func generateISBN() (number int64) {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}

func generatePrice(min float32, max float32) (number float32) {
	rand.Seed(time.Now().UnixNano())
	result := (rand.Float32() * (max - min)) + min
	result = float32(math.Floor(float64(result*100))) / 100
	return result
}

func generateIntegerWithMaxValue(max int32) (number int32) {
	rand.Seed(time.Now().UnixNano())
	result := rand.Int31n(max)
	return result
}

func generateStockCode() (code string) {
	rand.Seed(time.Now().UnixNano())
	stockCodeVar1 := [...]string{"ABC", "XYZ", "DEF", "GHI", "JKL"}
	stockCodeVar2 := [...]string{"K", "L", "M", "N"}
	stockCode := fmt.Sprintf("%s %s %v", stockCodeVar1[generateIntegerWithMaxValue(5)], stockCodeVar2[generateIntegerWithMaxValue(4)], generateIntegerWithMaxValue(10000))
	return stockCode
}

//Line message is used in many place
//Prints a line message.
func PrintLineMessage(message string) {
	fmt.Println(message)
}

func PrintBook(book *Book) {
	fmt.Printf("\tID:\t\t%v\n\tISBN:\t\t%v\n\tBook Name: \t%s\n\tAuthor:\t\t%s\n\tPage Count:\t%v\n\tPrice:\t\t%v\n\tStock Code:\t%s\n\tStock Count:\t%v\t\n\n", book.ID, book.ISBN, book.BookName, book.Author.Name, book.PageCount, book.Price, book.StockCode, book.StockCount)
}
