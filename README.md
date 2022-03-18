## Homework 2| Week 3

This program contains a list of books.

This program has five tasks.

* 'list' command: To show whole list of the books in the program as an output.
* 'search <bookName>' command:  If the list contains the name of the book or one single word of the book name received as input, the program gives the name of that book as an output.
* 'get <bookID>' commmand:  If the list contains the book ID received as input, the program gives the name of that book as an output.  
* 'delete <bookID>' command:    If the list contains the book ID received as input, the program deletes the book and return the new list as an output.
* 'buy <bookID> <quantity>' command:    If the list contains the book ID received as input and have enough stock, the program returns returns the total price and the latest status list of the books as an output 

The 'list', 'search', 'get', 'delete' or 'buy' commands should be used to implement these five tasks.

## Usage 

### 'list' command

``` 
go run main.go list
``` 
* The 'list' command should be used to access whole list of the books in the program as an output.


### 'search' command

``` 
go run main.go search <bookName> 
``` 
* The 'search' command should be used to gives the name of that book as output if the list contains the name of the that book

### get command

```
go run main.go get <bookID>
go run main.go get 5
```
* The 'get' command should be used to gets the book name as output if the list contains the book

### delete command

```
go run main.go delete <bookID>
go run main.go delete 5
```
* The 'delete' command should be used to delete the book with given book ID and see the new list of the books as an output

### buy command

```
go run main.go buy <bookID> <quantity>
go run main.go buy 5 2
```
* The 'buy' command should be used to buy the books that given book ID and quantity of books then returns the total price and the latest status list of the books as an output

> Note: The program returns usage messages when the commands are not entered accurately. User could follow that messages to use that program. 
