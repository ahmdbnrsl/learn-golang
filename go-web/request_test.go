package test  

import "net/http"
import "fmt"
import "testing"

func Test3Request(t *testing.T) {
    var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintln(writer, request.Method)
        if request.RequestURI == "/" {
            fmt.Fprintln(writer, "You now in a root page")
        } else if request.RequestURI == "/admin" {
            fmt.Fprintln(writer, "You now in a admin page")
        } else {
            fmt.Fprintln(writer, "Oops you're lost...")
        }
    }
    
    server := http.Server{
        Addr : "localhost:8080",
        Handler : handler,
    }
    
    err := server.ListenAndServe();
    
    if err != nil {
        panic(err)
    }
}