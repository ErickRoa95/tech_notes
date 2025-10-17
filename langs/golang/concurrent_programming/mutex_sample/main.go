package main

import (
	"fmt"
	"sync"
)

var o sync.Once

func main(){
	AppendTo1000()
	AppendTo1000()
}

func AppendTo1000(){
	// Will only run once, no matter how many times the func is called. 
	o.Do(func(){
		s := []int{}
		var wg sync.WaitGroup
		var m sync.Mutex // to protect shared memory

		const iterations = 1000
		wg.Add(iterations)
		for i:= 0 ; i <iterations ; i++{
			go func(){
				m.Lock()
				defer m.Unlock()
				s= append(s, 1) // Mutex is used to avoid multiple goroutines access the same variable at once.
				wg.Done()
			}()
		}
		wg.Wait()

		fmt.Println(len(s))
	})

}