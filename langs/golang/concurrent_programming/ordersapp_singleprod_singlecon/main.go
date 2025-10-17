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
}


// This app was done using the pattern: Single Producer - Single Consumer
func main() {
	// This is a placeholder for the main function.
	var wg sync.WaitGroup

	// Marshall all orders to the order struct. 
	receivedOrderCh :=  receivedOrders()
	// Get valid orders and invalid orders.
	validOrderCh, invalidOrderCh := validateOrders(receivedOrderCh)
	// update orders to Reserve status.
	reserverdOrderCh := reserveInventory(validOrderCh)
	
	wg.Add(2) 
	go func(invalidOrderCh <-chan invalidOrder){
		for order := range invalidOrderCh{
			fmt.Printf("Invalid order received: %v | Issue: %v\n", order.order, order.err)
		}
		wg.Done()
	}(invalidOrderCh)

	go func(reservedInventoryCh <- chan order){
		for order := range reservedInventoryCh{
			fmt.Printf("Inventoryu reserverd for: %v\n", order)
		}
		wg.Done()
	}(reserverdOrderCh)

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