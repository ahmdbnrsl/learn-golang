package main 

import (
    "regexp"
    "fmt"
)

func main() {
    var regex = regexp.MustCompile(`v([a-z, A-Z, 0-9])a`)
    
    fmt.Println(regex.MatchString("via"))
    fmt.Println(regex.MatchString("vaa"))
    fmt.Println(regex.MatchString("vIa"))
    
    fmt.Println(regex.FindAllString("via vua vaa voa vea vta", 20))
}