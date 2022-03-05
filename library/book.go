package library

type Book struct {
	ID         int
	BookName   string
	StockCode  string
	ISBN       int64
	PageCount  int
	Price      float32
	StockCount int
	Author
	IsDeleted bool
}

type Author struct {
	Name string
}

//constructor
func (b *Book) Init(id int, bookName string, stockCode string, isbn int64, pageCount int, price float32, stockCount int, authorName string) {
	b.ID = id
	b.BookName = bookName
	b.StockCode = stockCode
	b.ISBN = isbn
	b.PageCount = pageCount
	b.Price = price
	b.StockCount = stockCount
	b.IsDeleted = false
	b.Author.Name = authorName
}

//book library is initialized
//for each book book-list-index + 1 = ID
func CreateBookLibrary() {
	bookCount := len(BookNames)
	for index := 0; index < bookCount; index++ {
		b := new(Book)
		id := index + 1
		isbn := generateISBN()
		price := generatePrice(15, 100)
		pageCount := generateIntegerWithMaxValue(1000)
		stockCount := generateIntegerWithMaxValue(5) + 1
		stockCode := generateStockCode()
		b.Init(id, BookNames[index], stockCode, isbn, int(pageCount), price, int(stockCount), Authors[index])
		Books[index] = b
	}
}

//This function checks if a book exists
//with the given id and is not deleted.
func IsAvailable(id int) bool {
	if id > len(Books) {
		return false
	} else if Books[id-1].IsDeleted == true {
		return false
	}
	return true
}

//This function checks if the book exists,
//then check stock count
//then decreases stock
func BuyBooks(id int, count int) error {
	if !IsAvailable(id) {
		return ErrBookNotFoundWithId
	} else if Books[id-1].StockCount < count {
		return ErrStockIsNotEnough
	} else {
		Books[id-1].StockCount -= count
	}
	return nil
}

//Finds the book with given id and deletes it
func DeleteBook(id int) error {
	if id <= len(Books) {
		book := Books[id-1]
		book.Delete()
		return nil
	}
	return ErrBookNotFoundWithId
}

//Deletes book if not deleted already
func (book *Book) Delete() error {
	//Check if the book is already deleted or not
	if book.IsDeleted {
		return ErrBookNotFoundWithId
	} else {
		book.IsDeleted = true
		return nil
	}
}
