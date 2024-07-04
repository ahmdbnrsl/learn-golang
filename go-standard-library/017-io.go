package main 

import (
    "bufio"
    "io"
    "fmt"
    "strings"
    "os"
)

func main () {
    //Read
    input := strings.NewReader("Ahmad Beni Rusli\nVia Fitriana\n")
    
    reader := bufio.NewReader(input)
    
    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            break
        }
        fmt.Println(string(line))
    }
    //Write
    writer := bufio.NewWriter(os.Stdout)
    _, _ = writer.WriteString("Hello World!\n")
    
    writer.Flush()
}