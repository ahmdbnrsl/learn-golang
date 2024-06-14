package main

import "fmt"

func main() {
    person:= map[string]string{
        "name": "beni",
        "address": "Cilacap",
    }
    fmt.Println(person)
    fmt.Println(person["name"])
    
    //function
    //len
    fmt.Println(len(person))
    //get key data
    fmt.Println(person["address"])
    //set key data
    person["address"] = "Sidareja"
    fmt.Println(person["address"])
    //add data
    person["age"] = "19"
    fmt.Println(person["age"])
    //create new map
    newMap := make(map[string]string)
    newMap["key"] = "value"
    fmt.Println(newMap["key"])
    //delete key from map
    delete(person, "age")
    fmt.Println(person)
}