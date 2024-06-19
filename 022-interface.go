package main

import "fmt"

type Human interface {
    GetName() string
}

func SayHello(human Human) {
    println("Hello", human.GetName())
}

type Person struct{
    Name string
}
type Animal struct{
    Name string
}

func (person Person) GetName() string {
    return person.Name
}
func (animal Animal) GetName() string {
    return animal.Name
}

func main() {
    beni := Person{Name: "Beni"}
    fmt.Println(beni.GetName())
    gorilla := Animal{Name: "Gorilla"}
    SayHello(beni)
    SayHello(gorilla)
}