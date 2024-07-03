package main 

import (
    "fmt"
    "flag"
)

type Person struct{
    Name string
}

func main() {
    host := flag.String("host", "localhost", "Put your database host")
    username := flag.String("username", "root", "Put your database username")
    password := flag.String("password", "root", "Put your database password")
    
    flag.Parse()
    
    if *username != "Via" || *password != "viaacntik" {
        fmt.Println("Login gagal...")
    } else {
        fmt.Println("login berhasil...")
    }
    
    
    fmt.Println(*host, *username, *password)
    
    
    /*person := Person{"Via"}
    person2 := &person
    person3 := &person
    person2.Name = "Viaaaa"
    *person3 = Person{"Fitriana"}
    person.Name = "Beni"
    
    fmt.Println(person, *person2, person3)*/
}