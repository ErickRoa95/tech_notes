// go:build mainChannel
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

var rawOrders = []string{
	//type json
	`{"ProductCode": 1001, "Quantity": -1, "status": 1}`,
	`{"ProductCode": 1002, "Quantity": 10, "status": 1}`,
	`{"ProductCode": 1003, "Quantity": -15, "status": 1}`,
	`{"ProductCode": 1004, "Quantity": -20, "status": 1}`,
	`{"ProductCode": 1005, "Quantity": 25, "status": 1}`,
	`{"ProductCode": 1006, "Quantity": 25, "status": 1}`,
	`{"ProductCode": 1007, "Quantity": 25, "status": 1}`,
	`{"ProductCode": 1008, "Quantity": -25, "status": 1}`,
	`{"ProductCode": 1009, "Quantity": 25, "status": 1}`,
	`{"ProductCode": 1010, "Quantity": 25, "status": 1}`,
	`{"ProductCode": 1011, "Quantity": 25, "status": 1}`,
	`{"ProductCode": 1012, "Quantity": 25, "status": 1}`,
}

// Pattern used for this app: Single Producer - Multiple consumer.

func main() {
	// This is a placeholder for the main function.
	var wg sync.WaitGroup

	// Encapsulation patter for validate and receivedOrder functions. 
	receivedOrderCh :=  receivedOrders()
	validOrderCh, invalidOrderCh := validateOrders(receivedOrderCh)
	reserverdOrderCh := reserveInventory(validOrderCh)
	
	wg.Add(1) 
	go func(invalidOrderCh <-chan invalidOrder){
		for order := range invalidOrderCh{
			fmt.Printf("Invalid order received: %v | Issue: %v\n", order.order, order.err)
		}
		wg.Done()
	}(invalidOrderCh)

	// Will add 3 workers for Consuming information from channel 
	const workers = 3
	wg.Add(3)
	for i:=0 ; i < workers ; i ++ {
		go func(reservedInventoryCh <- chan order, worker int){
			for order := range reservedInventoryCh{
				fmt.Printf("Worker %d | Inventory reserverd for: %v\n", worker, order)
			}
			wg.Done()
		}(reserverdOrderCh, i)
	}

	wg.Wait()
}

// <-chan received-only channel
// chan<- send-only channel
func validateOrders(in <-chan order) (<-chan order, <-chan invalidOrder){
	//order := <-in
	out := make(chan order)
	errCh := make(chan invalidOrder, 1)
	go func (){
		for order := range in{
			if order.Quantity <= 0 {
				// Error condition
				errCh <- invalidOrder{
					order: order, 
					err: errors.New("not valid Quantity for order"),
				}
			}else{
				out <- order	 
			}
		}
		// Receibing all information from the buffer in, we close both channels to avoid deadlock. 
		close(out)
		close(errCh)
	}()

	return out, errCh
}

// chan <- send-only channel
//  will declare that the channel out will only send values, not receiving. 
func receivedOrders() <-chan order {
	out := make(chan order)
	go func(){
		for _, rawOrder := range rawOrders {
			var newOrder order
			err := json.Unmarshal([]byte(rawOrder), &newOrder	)
			if err != nil {
				log.Print(err)
				continue
			}
			out <- newOrder
		}
		// after sentind all rawOrders close channel, to indicate 
		// system that this channel wont received additional informations. 
		close(out)
		}() 

		return out
}

func reserveInventory(in <-chan order) <-chan order{
	out := make(chan order)

	go func(){
		for o := range in {
			o.Status = reserved
			out <- o
		}
		close(out)
	}()

	return out
}