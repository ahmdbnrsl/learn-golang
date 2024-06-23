package main 

import (
    "errors"
    "fmt"
)

var empetyNameError, tooShortNameError = errors.New("Name should not empety"), errors.New("name length must more than 3")

func getKhodam(name string) (string, error) {
    if name == "" {
        return "", empetyNameError
    } else if len(name) < 3 {
        return "", tooShortNameError
    } else {
        return "Kampas kopling", nil
    }
}

func main() {
    khodam, err := getKhodam("beni")
    if err != nil {
        if errors.Is(err, empetyNameError) {
            fmt.Println("Name is empety")
        } else if errors.Is(err, tooShortNameError) {
            fmt.Println("Name too short")
        } else {
            fmt.Println("gatau yg error yg mana")
        }
    } else {
        fmt.Println(khodam)
    }
}