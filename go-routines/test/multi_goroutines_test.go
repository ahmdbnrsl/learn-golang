package test  

import "fmt"
import "time"
import "testing"

func PrintNum(nums int) {
    fmt.Println("Count ke : ",  nums)
}

func TestPrintNum(t *testing.T) {
    for i := 0; i < 100000; i++ {
        go PrintNum(i)
    }
    time.Sleep(20 * time.Second)
}