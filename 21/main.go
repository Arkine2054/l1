package main

import "fmt"

type iPhone struct {
}

func (p iPhone) lighting() string {
	return fmt.Sprintf("Charging>>>")
}

type USB interface {
	canCharge() string
}

type lightingAdapter struct {
	adapter *iPhone
}

func (a lightingAdapter) canCharge() string {
	return a.adapter.lighting()
}

func main() {

	cable := &iPhone{}

	var adapter USB = &lightingAdapter{adapter: cable}

	fmt.Printf("Your iPhone: %v", adapter.canCharge())
}
