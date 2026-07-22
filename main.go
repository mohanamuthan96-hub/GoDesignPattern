package main

import "fmt"

//Adapter Design pattern - Extended Implementation

//Target interface
type mobile interface {
	chargeAppleMobile()
}

//Concrete Prototype
type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Apple Mobile is charging")
}

//adaptee
type android struct{}

func (a *android) chargeAndroidMobile() {
	fmt.Printf("Charging android mobile")
}

//adapter - extend the functionality
type androidadapter struct {
	android *android
}

func (ad *androidadapter) chargeAppleMobile() {
	ad.android.chargeAndroidMobile()
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
	android := &android{}
	androidadapter := &androidadapter{android: android}
	client.chargeMobile(androidadapter)
}
