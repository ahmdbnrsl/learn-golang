package main

import "fmt"

type Human struct{
    Name string
    Age int8
}

func (human *Human) Married() {
    human.Name = "Mrs. " + human.Name
}

func main() {
    beni := &Human{"Via Fitriana", 20}
    beni.Married()
    
    fmt.Println(beni.Name)
}