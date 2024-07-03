package main 

import (
    "fmt"
    "os"
)

func main() {
    var args []string = os.Args
    for _, arg := range args {
        fmt.Println(arg)
    }
    
    //hostname
    hostname, err := os.Hostname()
    if err == nil {
        fmt.Println(hostname)
    } else {
        fmt.Println(err.Error())
    }
}