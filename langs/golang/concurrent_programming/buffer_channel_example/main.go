package main

import "fmt"

func main(){
	ch := make(chan string) // Bidirectional channel. 
	go func (ch chan<- string)  { // chan<- send-only channel. 
		ch <- "message"
	}(ch)

	go func (ch <-chan string)  { // <-chan receive-only channel. 
		fmt.Println(<-ch)
	}(ch)
}

