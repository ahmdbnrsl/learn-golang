package test  

import "fmt"
import "time"
import "testing"

func SayHello(name string) {
    fmt.Println("Hello " + name)
}

func TestSayHello(t *testing.T) {
    go SayHello("Xaviera")
    fmt.Println("I Love You")
    
    time.Sleep(1 * time.Second)
}