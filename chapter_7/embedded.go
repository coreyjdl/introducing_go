package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Greet() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

type Android struct {
	Person
	Model string
}


func main() {
	p := &Person{Name: "Alice"}
	p.Greet()

	a := &Android{Person: Person{Name: "Robo"}, Model: "X1000"}
	a.Greet()

	a.Person.Greet()
}