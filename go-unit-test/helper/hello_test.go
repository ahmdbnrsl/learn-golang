package helper 

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "runtime"
)

func TestHello1(t *testing.T) {
    result := SayHello("Via")
    
    assert.Equal(t, "Hello Via", result, "unexpected result")
    println("end")
}

func TestHello2(t *testing.T) {
    result := SayHello("Fitriana")
    
    assert.NotEqual(t, "Hello Fitriana", result, "unexpected result")
    println("end")
}

func TestSkip(t *testing.T) {
    if runtime.GOOS == "android" {
        t.Skip("Your operating system is not supported!")
    }
    result := SayHello("Via")
    assert.Equal(t, "Hello Via", result, "unexpected result")
}