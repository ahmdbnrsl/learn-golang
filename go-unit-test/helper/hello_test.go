package helper 

import "testing"

func TestHello1(t *testing.T) {
    result := SayHello("Via")
    
    if result != "Hello Via" {
        t.Fail()
    }
    println("end")
}

func TestHello2(t *testing.T) {
    result := SayHello("Fitriana")
    
    if result != "Hello Fitriana" {
        t.FailNow()
    }
    println("end")
}