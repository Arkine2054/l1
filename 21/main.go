package main

import "fmt"

type iPhone struct {
}

func (p iPhone) USB() string {
	return fmt.Sprintf("Charging>>>")
}

type lighting interface {
	canCharge() string
}

type usbAdapter struct {
	adapter *iPhone
}

func (a usbAdapter) canCharge() string {
	return a.adapter.USB()
}

func main() {

	cable := &iPhone{}

	var adapter lighting = &usbAdapter{adapter: cable}

	fmt.Printf("Your iPhone: %v", adapter.canCharge())
}
