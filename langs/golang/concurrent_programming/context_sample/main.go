package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main(){
	// context.BACKGROUND  is main execution background. 
	// context.WithCancel(parent context) -> Parent context is the thread/goroutine.
	//ctx, cancel := context.WithCancel(context.Background())
	// context.WithTimeout (provide time for the context to get cancelled. )
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel() // required to free resources when canel after timeout.

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context){
		for range time.Tick(500 * time.Millisecond){
			if ctx.Err() != nil{
				log.Println(ctx.Err())
				return
			}
			fmt.Println("tick!")
		}
		wg.Done()
	}(ctx)

	// Code only used for context.WithCancel()
	// time.Sleep(5 * time.Second)
	// cancel() // call cancellation of context. 

	wg.Wait()
}