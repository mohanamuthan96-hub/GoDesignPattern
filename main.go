package main

import (
	"fmt"
	"sync"
)

type configure struct {
	//configuration
}

var configureinstance *configure
var lock = &sync.Mutex{}
var counter int

func getConfigureInstances() *configure {
	if configureinstance == nil {
		(*lock).Lock()
		defer lock.Unlock()
		if configureinstance == nil {
			fmt.Println("Creating a instance of singleton")
			counter = counter + 1
			configureinstance = &configure{}
		} else {
			fmt.Println("Instance already created, counter is", counter)
		}
	} else {
		fmt.Println("Instance already created, counter is", counter)
	}
	return configureinstance
}
func main() {
	for range 50 {
		go getConfigureInstances()
	}
	fmt.Scanln()
}
