package main 

import (
    "fmt"
    "container/list"
)

func main() {
    data := list.New()
    data.PushBack("Via")
    data.PushBack("Fitriana")
    data.PushBack("Hehe")
    
    for i := data.Front(); i != nil; i = i.Next() {
        fmt.Println(i.Value)
    }
}