package main

import "fmt"

func main(){
	ch := make(chan string,5)

	names := [5]string {"caro", "eliot", "erick","ivan", "tony"}

	for _,name := range names {
		ch <- name
	}
	close(ch) // close ch after all messages where send. 

	for _,msg := range <-ch {
		fmt.Println(msg)
	}
}