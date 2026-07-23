package main

import "fmt"

type human struct {
	name   string
	age    int
	height float64
	color  string
}

//Empty Constructor
func newhuman() human {
	return human{}
}

//Populated Constructor
func newHumanWithField(name string, age int, color string, height float64) human {
	return human{age: age, name: name, color: color, height: height}
}

//Builders
func (human human) withName(name string) human {
	human.name = name
	return human
}

func (human human) withAge(age int) human {
	human.age = age
	return human
}

func (human human) withColor(name string) human {
	human.color = name
	return human
}

func (human human) withHeight(height float64) human {
	human.height = height
	return human
}

//reset
func (human human) reset() human {
	human = newhuman()
	return human
}

func giant() human {
	return newhuman().withHeight(3.2).withColor("Green")
}
func main() {

	human := newhuman().withName("Naresh").withHeight(1.5).withAge(23).withColor("Darkbrown")
	you := newHumanWithField("MOhan", 29, "Brown", 1.6)
	fmt.Printf("%+v\n", human)
	fmt.Printf("%+v\n", you)
	youngGiant := giant().withName("Young Giant").withAge(2)
	oldGiant := giant().withName("Old Giant").withAge(52)
	fmt.Printf("%+v\n", youngGiant)
	fmt.Printf("%+v\n", oldGiant)
}
