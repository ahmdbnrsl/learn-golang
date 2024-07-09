package helper

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestHello3(t *testing.T) {
    result := SayHello("Via")
    
    require.Equal(t, "Hello Via", result, "unexpected result")
    println("end")
}

func TestHello4(t *testing.T) {
    result := SayHello("Fitriana")
    
    require.NotEqual(t, "Hello Fitriana", result, "unexpected result")
    println("end")
}