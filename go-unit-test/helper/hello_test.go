package helper 

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "runtime"
)

func TestMain(m *testing.M) {
    println("Starting test....")
    m.Run()
    println("Testing finish....")
}

func TestHello1(t *testing.T) {
    result := SayHello("Via")
    
    assert.Equal(t, "Hello Via", result, "unexpected result")
    println("end")
}

func TestHello2(t *testing.T) {
    result := SayHello("Fitriana")
    
    assert.NotEqual(t, "Hello Via", result, "unexpected result")
    println("end")
}

func TestHello5(t *testing.T) {
    t.Run("Via", func(t *testing.T) {
        result := SayHello("Fitriana")
        assert.NotEqual(t, "Hello Via", result, "unexpected result")
        println("end")
    })
}

func TestHelloTable(t *testing.T) {
    tests := []struct {
        name string
        request string
        expected string
    } {
        {
            name: "SayHello(Via)",
            request: "Via",
            expected: "Hello Via",
        },
        {
            name: "SayHello(Fitriana)",
            request: "Fitriana",
            expected: "Hello Fitriana",
        },
    }
    
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := SayHello(test.request)
            assert.Equal(t, test.expected, result, "unexpected result")
            println("end")
        })
    }
    
}

func TestSkip(t *testing.T) {
    if runtime.GOOS == "android" {
        t.Skip("Your operating system is not supported!")
    }
    result := SayHello("Via")
    assert.Equal(t, "Hello Via", result, "unexpected result")
}