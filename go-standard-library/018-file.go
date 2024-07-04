package main 

import (
    "os"
    "bufio"
    "io"
    "fmt"
)


func createNewFile(name, message string) error {
    file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)  
    if err != nil {
        return err
    }
    defer file.Close()
    file.WriteString(message)
    return nil
}

func readFile(name string) (string, error) {
    file, err := os.OpenFile(name, os.O_RDONLY, 0666)  
    if err != nil {
        return "", err
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    var message string
    for {
        line, _, err := reader.ReadLine()
        message += string(line) + "\n"
        if err == io.EOF {
            break
        }
    }
    
    return message, nil
}

func addToFile(name, message string) error {
    file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        return err
    }
    defer file.Close()
    file.WriteString("\n" + message)
    return nil
}

func main () {
    //createNewFile("ya.zig", "anu")
    let, _ := readFile("ya.zig")
    fmt.Println(let)
    addToFile("ya.zig", "iya itu")
}