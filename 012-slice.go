package main

import "fmt"

func main() {
    sholat := [5]string{
        "shubuh",
        "dzuhur",
        "ashar",
        "maghrib",
        "isya",
    }
    
    slice := sholat[:]
    
    fmt.Println(slice)
    println(slice[0])
    //slice length 
    println(len(slice))
    
    bolong := sholat[1:4]
    fmt.Println(bolong)
    //capacity 
    println(cap(bolong))
}