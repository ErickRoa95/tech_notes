package main

import "fmt"

type order struct{
	ProductCode int
	Quantity float64
	Status orderStatus
}

type invalidOrder struct {
	order order 
	err error
}

// String provides a string representation of the order
func (o order) String() string {
	return fmt.Sprintf("ProductCode: %d, Quantity: %.2f, Status: %s", 
		o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(status orderStatus) string {
	switch status {
	case none:
		return "None"
	case new:
		return "New"
	case received:
		return "Received"
	case reserved:
		return "Reserved"
	case filled:
		return "Filled"
	default:
		return "Unknown"
	}
}


type	 orderStatus int

// Enumeration of order statuses
const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

var orders []order