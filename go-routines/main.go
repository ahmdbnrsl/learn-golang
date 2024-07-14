package main 

import "fmt"
import "time"

func main() {
    for i := 0; i < 100; i++ {
        go func(it int) {
            fmt.Println(it)
        }(i)
    }
    
    time.Sleep(4 * time.Second)
}