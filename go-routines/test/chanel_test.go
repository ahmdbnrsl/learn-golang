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