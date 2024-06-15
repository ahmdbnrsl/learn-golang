package main

import "fmt"

func main() {
    count := 1
    
    for count <= 10 {
        println(count)
        count++
    }
    println(count)
    
    //statement
    for i := 1; i <= 10; i++ {
        println(i)
    }
    
    //looping collection
    
    names := []string{
        "Ahmad",
        "Beni",
        "Rusli",
    }
    
    for i := 0; i < len(names); i++ {
        println(names[i])
    }
    //use for range
    for i, name := range names {
        println(i + 1, name)
    }
    //unuse index variable
    for _, name := range names {
        println(name)
    }
    
    //break
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
        
        println(i)
    }
    
    //continue
    for i := 0; i < 10; i++ {
        if i % 2 == 0 {
            continue
        }
        
        println(i)
    }
    
    //data filtration 
    println("STUDI KASUS")
    
    nilai := []int8{
        80,
        90,
        70,
        80,
        70,
        90,
        90,
        60,
        70,
        50,
    }
    
    nilai90 := [20]int8{}
    for i := 0; i < len(nilai); i++ {
        if nilai[i] != 90 {
            continue
        }
        nilai90[i] = nilai[i]
    }
    fmt.Println(nilai90)
}