package main

import "log"

type animal interface {
	sound()
}

type cat struct {
}

func newcat() cat {
	return cat{}
}
func (cat cat) sound() {
	log.Print("Meow")
}

type dog struct {
	name string
}

func newDog(name string) *dog {
	return &dog{}
}
func (dog dog) sound() {
	log.Print("wow")
}

func farm(sound int) {
	var animal animal
	if sound > 30 {
		animal = newDog("Dogh")
	} else {
		animal = newcat()
	}
	animal.sound()
}
func main() {
	//Decibel greater 30 is cat else dog
	farm(54)
}
