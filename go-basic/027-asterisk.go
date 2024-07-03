package main

import "fmt"

type Person struct{
    Name, Address string
}

func main() {
    person1 := Person{"Via Fitriana", "Cilacap"}
    pers1 := Person{"Ben ben", "Cilacap"}
    person2 := &person1
    pers2 := &pers1
    pers3 := &pers1
    
    person2.Name = "Via"
    pers2.Name = "Fitriana"
    
    
    person2 = &Person{"Beni", "Cilacap"}
    *pers2 = Person{"Benu", "Cilacap"}
    
    fmt.Println(person1)
    fmt.Println(person2)
    fmt.Println(pers1)
    fmt.Println(pers3)
    //asterisk operator
    
}