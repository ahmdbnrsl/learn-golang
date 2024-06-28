package main 

import (
    "fmt"
    "container/ring"
    "strconv"
)


func main() {
    data := ring.New(5)
    for i := 0; i < data.Len(); i++ {
        data.Value = "Value " + strconv.FormatInt(int64(i + 1), 10)
        data = data.Next()
    }
    
    data.Do(func(value any) {
        fmt.Println(value)
    })
    
    
}