package main

import "fmt"

//Adapter Design pattern

//Target interface
type mobile interface {
	chargeAppleMobile()
}

//Concrete Prototype
type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Apple Mobile is charging")
}

//client
type client struct{}

func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleMobile()
}
func main() {
	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)
}
