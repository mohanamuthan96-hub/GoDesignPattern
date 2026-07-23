package main

import "fmt"

//State Design Pattern

type tvState interface {
	state()
}

//Concrete State implementation

type on struct{}

func (o *on) state() {
	fmt.Println("TV is on!")
}

type off struct{}

func (o *off) state() {
	fmt.Println("TV is off!")
}

type stateContext struct {
	currenttvstate tvState
}

func getcontext() *stateContext {
	return &stateContext{
		currenttvstate: &off{},
	}
}

func (sc *stateContext) setState(state tvState) {
	sc.currenttvstate = state
}

func (sc *stateContext) getState() {
	sc.currenttvstate.state()

}

//clinet
func main() {
	tvContext := getcontext()
	tvContext.getState()
	tvContext.setState(&on{})
	tvContext.getState()
}
