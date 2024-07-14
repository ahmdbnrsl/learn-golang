package test

import "fmt"
import "testing"
import "time"

func TestChannel(t *testing.T) {
    
    channel := make(chan string)
    defer close(channel)
    
    go func() {
        time.Sleep(2 * time.Second)
        channel <- "Ahmad Beni Rusli"
        fmt.Println("Finishing...")
    }()
    
    data := <- channel
    fmt.Println(data)
    
    // channel <- "Beni"
//     data := <- channel
//     fmt.Println(<- channel)
    
    
}

func TestChannel2(t *testing.T) {
    channel := make(chan int)
    
    go func() {
        channel <- 10  //Memasukan data ke dalam channel
    }()
    go func() {
        fmt.Println(<- channel)
    }()
    time.Sleep(3 * time.Second)
}

func pow(params int) chan int {
    channel := make(chan int)
    go func(it int) {
        channel <- it * it
    }(params)
    return channel
}

func TestAsyncFunction(t *testing.T) {
    a := pow(20)
    b := pow(10)
    fmt.Println(<- a, <-b)
}