package test  

import "fmt"
import "time"
import "testing"

func TestGoRoutines(t *testing.T) {
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
    
    time.Sleep(5 * time.Second)
}