package main

import "fmt"

//is error
/*func createName(firstName, lastName string) string {
    if firstName == "" && lastName == "" {
        return nil
    } else {
        return firstName + " " + lastName
    }
}*/

func NewMap(name string) map[string]string {
    if name == "" {
        return nil
    } else {
        return map[string]string {
            "name": name,
        }
    }
}

func main () {
    nama := NewMap("")
    
    if nama == nil {
        println("Data kosong")
    } else {
        fmt.Println(nama)
    }
}