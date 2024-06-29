package main 

import (
    "sort"
    "fmt"
)

type SortInterface interface{
    //number of element
    Len() int
    //index
    Less(i, j int) bool
    //swap the element with indexes i and j
    Swap(i, j int)
}

type User struct{
    Name string
    Age int
}

type UserSlice []User

func(userSlice UserSlice) Len() int {
    return len(userSlice)
}

func(userSlice UserSlice) Less(i, j int) bool {
    return userSlice[i].Age < userSlice[j].Age
}

func(userSlice UserSlice) Swap(i, j int) {
    userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
    var users UserSlice = []User{
        {"Via", 20},
        {"Chindy", 19},
        {"Cinta", 19},
        {"Nessa", 21},
    }
    
    sort.Sort(users)
    
    fmt.Println(users)
}