package main 

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now.Local())
    
    utc := time.Date(2004, time.November, 10, 23, 0, 0, 0, time.UTC)
    fmt.Println(utc.Local())
    
    parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
    fmt.Println(parse)
    
    //DURATION
    var duration time.Duration = 100 * time.Second
    var duration2 time.Duration = 10 * time.Millisecond
    var duration3 time.Duration = duration - duration2
    
    
    fmt.Println(duration3)
}