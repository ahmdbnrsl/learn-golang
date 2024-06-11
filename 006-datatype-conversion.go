package main

import "fmt"

func main() {
    var val1 int32 = 32768
    var val2 int64 = int64(val1)
    var val3 int16 = int16(val2)
    
    fmt.Println(val1)
    fmt.Println(val2)
    fmt.Println(val3)
    
    //convert byte to string
    
    const name string = "Via Fitriana"
    var e uint8 = name[0]
    var eString string = string(e)
    
    fmt.Println(e)
    fmt.Println(eString)
}