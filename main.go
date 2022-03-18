package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	BookID         string
	BookName       string
	BookPageNumber int
	BookStock      int
	BookPrice      float64
	BookStockCode  string
	BookISBN       string
	BookAuthorID   int
	BookAuthorName string
}

func Books() []*Book {

	book1 := NewBook()
	book1 = book1.SetID("1")
	book1 = book1.SetBookName("Guns, Germs, and Steel")
	book1 = book1.SetBookPageNumber(528)
	book1 = book1.SetBookStock(10)
	book1 = book1.SetBookPrice(18.95)
	book1 = book1.SetBookStockCode("GXGXAXSX")
	book1 = book1.SetBookISBN("978-0393354324")
	book1 = book1.SetBookAuthorID(123)
	book1 = book1.SetBookAuthorName("Jared Diamond")

	book2 := NewBook()
	book2 = book2.SetID("2")
	book2 = book2.SetBookName("Escape From Freedom")
	book2 = book2.SetBookPageNumber(309)
	book2 = book2.SetBookStock(8)
	book2 = book2.SetBookPrice(17.99)
	book2 = book2.SetBookStockCode("EXFXFX")
	book2 = book2.SetBookISBN("978-0805031492")
	book2 = book2.SetBookAuthorID(98)
	book2 = book2.SetBookAuthorName("Erich Fromm")

	book3 := NewBook()
	book3 = book3.SetID("3")
	book3 = book3.SetBookName("The Grapes of Wrath")
	book3 = book3.SetBookPageNumber(464)
	book3 = book3.SetBookStock(5)
	book3 = book3.SetBookPrice(18.00)
	book3 = book3.SetBookStockCode("TXGXOXWX")
	book3 = book3.SetBookISBN("978-0143039433")
	book3 = book3.SetBookAuthorID(231)
	book3 = book3.SetBookAuthorName("John Steinbeck")

	book4 := NewBook()
	book4 = book4.SetID("4")
	book4 = book4.SetBookName("Blindness")
	book4 = book4.SetBookPageNumber(352)
	book4 = book4.SetBookStock(9)
	book4 = book4.SetBookPrice(15.99)
	book4 = book4.SetBookStockCode("BX")
	book4 = book4.SetBookISBN("978-0156007757")
	book4 = book4.SetBookAuthorID(151)
	book4 = book4.SetBookAuthorName("Jose Saramago")

	book5 := NewBook()
	book5 = book5.SetID("5")
	book5 = book5.SetBookName("Crime and Punishment")
	book5 = book5.SetBookPageNumber(565)
	book5 = book5.SetBookStock(11)
	book5 = book5.SetBookPrice(17.95)
	book5 = book5.SetBookStockCode("CXAXPX")
	book5 = book5.SetBookISBN("978-0679734505")
	book5 = book5.SetBookAuthorID(321)
	book5 = book5.SetBookAuthorName("Fyodor Dostoyevsky")

	book6 := NewBook()
	book6 = book6.SetID("6")
	book6 = book6.SetBookName("Letter to a Child Never Born")
	book6 = book6.SetBookPageNumber(116)
	book6 = book6.SetBookStock(14)
	book6 = book6.SetBookPrice(24.00)
	book6 = book6.SetBookStockCode("LXTXAXCXNXBX")
	book6 = book6.SetBookISBN("978-0385134859")
	book6 = book6.SetBookAuthorID(412)
	book6 = book6.SetBookAuthorName("Oriana Fallaci")

	books := []*Book{book1, book2, book3, book4, book5, book6}

	return books
}

var usageMsg = `Commands:
	list			: The 'list' command should be used to access whole list of the book names and book IDs in the program as an output. 
	search <bookName>	: The 'search' command should be used to gives the name of that book as output if the list contains the name of the that book 	
	get <bookID>		: The 'get' command should be used to gets the book name as output if the list contains the book
	delete <bookID>		: The 'delete' command should be used to delete the book with given book ID and see the new list of the books as an output
	buy <bookID> <quantity>	: The 'buy' command should be used to buy the books that given book ID and quantity of books then returns the total price and the 						  latest status list of the books as an output
`

func main() {

	books := Books()

	args := os.Args[1:]

	if len(args) == 0 {
		usageAndExit("Running without any command")
	}

	if args[0] == "list" {

		if len(args) == 1 {

			list(books)

		} else {
			usageAndExit("Only the 'list' command should be entered for seeing the list of the Book Archive")
		}

	} else if args[0] == "search" {

		if len(args) == 2 {

			search(args, books)

		} else {
			usageAndExit("Please specify the book name after the 'search' command as 'search <the name of the book>'")
		}

	} else if args[0] == "get" {

		if len(args) == 2 {

			get(args[1], books)

		} else {
			usageAndExit("Please specify the book ID after the 'get' command as 'get <bookID>'")
		}

	} else if args[0] == "delete" {

		if len(args) == 2 {

			deleteIndex := delete(args[1], books)
			if deleteIndex > -1 {
				s := fmt.Sprintf("The Book with ID '%s' deleted\n\n", books[deleteIndex].BookID)
				print(s)
				books = append(books[:deleteIndex], books[deleteIndex+1:]...)
				list(books)
			} else {
				fmt.Println("The Book ID cannot be found")
			}
		} else {

			usageAndExit("Please specify the book ID after the 'delete' command as 'delete <bookID>'")

		}
	} else if args[0] == "buy" {

		if len(args) == 3 {

			argsQuantityInt, _ := strconv.Atoi(args[2])
			index := 0
			index = buy(args[1], argsQuantityInt, books, index)
			if index > -1 {
				s := fmt.Sprintf("Book ID: %s\n Book Name: %s\n Book Page Count: %d\n Book Stock Count: %d\n Book Price: %.2f\n Book Stock Code: %s\n Book ISBN: %s\n Book Author ID: %d\n Book Author Name: %s\n",
					books[index].BookID, books[index].BookName, books[index].BookPageNumber, books[index].BookStock, books[index].BookPrice, books[index].BookStockCode, books[index].BookISBN, books[index].BookAuthorID, books[index].BookAuthorName)
				print(s)
			} else if index == -1 {
				fmt.Println("The Book ID cannot be found")
			} else {
				fmt.Println("There is not enough stock of that book")

			}
		} else {
			usageAndExit("Please specify the book ID and  the count of books after the 'buy' command as 'buy <bookID> <bookCount>'")
		}

	} else {
		usageAndExit("Unknown Command")
	}

}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) SetID(bookID string) *Book {
	b.BookID = bookID
	return b
}

func (b *Book) SetBookName(bookName string) *Book {
	b.BookName = bookName
	return b
}

func (b *Book) SetBookPageNumber(bookPageNumber int) *Book {
	b.BookPageNumber = bookPageNumber
	return b
}

func (b *Book) SetBookStock(bookStock int) *Book {
	b.BookStock = bookStock
	return b
}

func (b *Book) SetBookPrice(bookPrice float64) *Book {
	b.BookPrice = bookPrice
	return b
}

func (b *Book) SetBookStockCode(bookStockCode string) *Book {
	b.BookStockCode = bookStockCode
	return b
}

func (b *Book) SetBookISBN(bookISBN string) *Book {
	b.BookISBN = bookISBN
	return b
}

func (b *Book) SetBookAuthorID(bookAuthorID int) *Book {
	b.BookAuthorID = bookAuthorID
	return b
}

func (b *Book) SetBookAuthorName(bookAuthorName string) *Book {
	b.BookAuthorName = bookAuthorName
	return b
}

func list(books []*Book) {

	print(fmt.Sprintf("%s\n\n", "The list of the Book Archive shown below;"))

	for i := 0; i < len(books); i++ {
		s := fmt.Sprintf("%s-%s\n", books[i].BookID, books[i].BookName)
		print(s)

	}

}

func search(args []string, books []*Book) {

	var userBookNameInput string
	userBookNameInput = strings.Join(args[1:], " ")

	var bookExistInList bool
	bookExistInList = false

	for i := 0; i < len(books); i++ {

		if strings.ToLower(userBookNameInput) == strings.ToLower(books[i].BookName) || strings.Contains(strings.ToLower(books[i].BookName), strings.ToLower(args[1])) {
			s := fmt.Sprintf("%s-'%s'\n", books[i].BookID, books[i].BookName)
			print(s)
			bookExistInList = true
		}

	}

	if bookExistInList == false {
		fmt.Println("The book does not exist in the list")
	}

}

func get(argsBookID string, books []*Book) {

	for i := 0; i < len(books); i++ {

		if books[i].BookID == argsBookID {
			print(fmt.Sprintf("%s-%s", books[i].BookID, books[i].BookName))
			return
		}

	}
	fmt.Println("The Book ID cannot be found")

}

func delete(argsBookID string, books []*Book) int {

	delIndex := -1

	for i := 0; i < len(books); i++ {

		if books[i].BookID == argsBookID {

			delIndex = i

			break

		}

	}
	return delIndex
}

func buy(argsBookID string, argsQuantity int, books []*Book, index int) int {

	newIndex := -1

	for i := 0; i < len(books); i++ {

		if books[i].BookID == argsBookID {

			if books[i].BookStock >= argsQuantity {

				books[i].SetBookStock(books[i].BookStock - argsQuantity)
				newIndex = i
				s := fmt.Sprintf("Total cost of book or books: %.2f\n\n", float64(argsQuantity)*books[i].BookPrice)
				print(s)

				break

			} else {

				newIndex = -2
				break

			}

		}

	}
	return newIndex
}

func usageAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n\n")
	fmt.Fprintf(os.Stderr, usageMsg)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)

}
