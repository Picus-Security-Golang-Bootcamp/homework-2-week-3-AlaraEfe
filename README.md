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
  
  ***

# Concurrency In Golang

"So if you look around in the world at large, what you see is a lot of independently executing things" said Rob Pike, the designer of Golang in 2012 Google I/O conference.

In the real world everything moves quickly and acts independently at the same time. Therefore, if we have a program that runs the function step by step, we could not be capable of dealing with real world problems. Because, the real world problems are not basic, on the contrary these are quite complex problems. So, we need something else to deal with these complex real life problems. This approach is the subtext of Concurrency. In this post we will take a quick look at Concurrency in Golang.  

![Dining philosophers](https://media.geeksforgeeks.org/wp-content/uploads/20210522124013/dfp-300x220.png)

> The dining philosophers problem is an example problem often used in concurrent algorithm design to illustrate synchronization issues and techniques for resolving them. You can also check dining philosophers problem to get more understanding of concurrency topic.  

## Concurrency

According to Rob Pike, concurrency is the composition of independently executing computations. In other terms, concurrency is the ability for functions to run independent of each other. Concurrency is the ability of a program to run multiple tasks independently during overlapping periods.  In a concurrent program, several tasks can run at the same time in no particular order

## Goroutines

Goroutine is the first important key topic of concurrency in Golang. A goroutine is a function that is capable of running concurrently with other functions. When you create a function as a goroutine, it has been treated as an independent unit of work that gets scheduled and then executed on an available logical processor. Basically, goroutine is a function that executes simultaneously with other goroutines in a program. We can consider goroutines as the main function. Actually, Golang’s main function is also a goroutine.  

   * Goroutines is an independently executing functions, launch by a go statement
   * Goroutines provides us speed, efficiency and also cost efficiency

Let's see with a quick example and create a goroutine. 

```
package main

import "fmt"

func hello(){
    fmt.Println("Hello Goroutine") 
}

func main(){
    go hello()
    time.Sleep(100 * time.Millisecond)
}
```

The output of this code shown below

```
Hello Goroutine
```

We said 'Hello' to a goroutine with this code. That is how we create goroutines. However, if we did not put the "time.Sleep(100 * time.Millisecond)" row, the goroutine did not have enough time to execute before the main and the program could not return "Hello Goroutine". That is because the program’s execution begins by initializing the main package and then calling the function 'main'. When the'main' function calls returns, the program exits. It does not wait for other goroutines to complete. In other terms, when the main function completes its execution, it doesn’t wait for other goroutines to complete their operations. That is why it is highly significant that the main function should wait for the other goroutines to complete their operations. However, don't worry. Go has a solution for that condition other than timer. Which is "WaitGroups" 

### WaitGroups

In simple terms, WaitGroups allow us to block until all goroutines within that waitgroup complete their execution.

So, Let's see how we can implement the wait group in the previous example. 

```
package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	fmt.Println("Hello Goroutine, it is good to see you here:)")
	wg.Done() /////// notifies the waitgroup that it finished
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1) /////// adds an entry to the waitgroup counter
	go hello(&wg)
	wg.Wait() ////// blocks execution until the goroutine finishes
	fmt.Println("Main ended")
}
```

The output of this code shown below

```
Hello Goroutine, it is good to see you here:)
Main ended
```
You can possibly think Timer also can be used instead of WaitGroup. However, we applied a timer to a simple operation. Let's think we have more complex and long operations like in real life. In that case, it is clear that using a timer would not be the best practice.  

* WaitGroups better and faster, because we don’t have to wait a fixed amount of time. 

We created goroutine. As we learned Goroutine provides us independence however, what could be the missing thing here. Things happen independently in real life, but at the same time these things communicate between each other. So, actually if we could not provide the communication between these independent actions, then we could not synchronize and there would be no flow of nothing. That could have been a major mistake for the development of many of the things we have today.

Fortunately, Go has 'Channels' to provide communication between Goroutines.

## Channels 

Channels are the second important key topic of concurrency in Golang. Channels provide the communication of goroutines as we mentioned before. In simple terms, a channel is a pipe that allows a goroutine to either put or take the data. To illustrate this, one goroutine sends data to a channel, other goroutines receive that data from the other side of this channel. There are 2 types of channels called buffered and unbuffered. There are some differences between them.

Let's see how we create a channel 
 
``` 
    // the channel can be in any type, in this example we make a 'int' channel
	ch := make(chan int) // unbuffered channel  
	ch := make(chan int, 2) // buffered channel
```

There is a given capacity to hold data in the buffered channel. But an unbuffered channel does not have the capacity to hold more than one data. That means, only one piece of data fits through the unbuffered channel at a time.

Writing and reading from an unbuffered channel is blocking operation. When one goroutine sends data, it is blocked until other goroutines receive data from the channel. It is the same for the receiving part. When one goroutine tries to receive data from the channel, it is blocked until data sent to the channel. It is kind of the same for the buffered channel. Sender goroutine is blocked when capacity is full until other goroutines fetch the data from the channel.

Now, Let's look at how we use channels to put and take the data in Go.

``` 
	ch <-     // put the data in channel
	   <- ch // transfer the data from channel to somewhere
```
We saw how to create the channel and put and get data of channel syntax. Now, Let's see with a quick example how channels work in Golang.

``` 
package main

import "fmt"

// Create gouroutine function
func printChannelData(c chan string) {
	fmt.Println("Data in channel is:", <-c)
}

func main() {

	fmt.Println("Main started")
	//create channel of string
	c := make(chan string)
	// call to goroutine
	go printChannelData(c)
	// put the data in channel
	c <- "data is coming to channel"
	fmt.Println("Main ended")
}
```

The output of this code shown below

```
Main started
Data in channel is: data is coming to channel
Main ended
```
We created a channel and put this channel data, then received that data from this channel and printed it. 

So far we have briefly learned about Concurrency in Golang and its concepts such as goroutines, waitgroups and channels. I hope you found this useful and it's been a good start for you on concurrency in Golang.

Thank you for your time and reading this post.

If you have any questions, please do not hesitate to contact with me. 

*** 

### References

[Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)

[Concurrency patterns in Golang: WaitGroups and Goroutines](https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/) 

[Concurrency and Channels in Go](https://medium.com/trendyol-tech/concurrency-and-channels-in-go-bbc4dea75286)

[Getting Started With Golang Channels! Here’s Everything You Need to Know](https://www.velotio.com/engineering-blog/understanding-golang-channels#:~:text=So%2C%20what%20are%20the%20channels,put%20or%20read%20the%20data.)

[Common pitfalls when using goroutines](https://reshefsharvit.medium.com/common-pitfalls-and-cases-when-using-goroutines-15107237d4f5)
