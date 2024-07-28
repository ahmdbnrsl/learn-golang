package test  

import "net/http"
import "testing"
import "fmt"

func Test2ServeMux(t *testing.T) {
    handler := http.NewServeMux()
    handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer, "Hello World!")
    })
    
    handler.HandleFunc("/admin", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer, "Hallo admin")
    })
    
    handler.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer, "Hallo bung")
    })
    
    handler.HandleFunc("/images/thumb", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprint(writer, "Hallo thumb.")
    })
    
    server := http.Server{
        Addr : "localhost:8080",
        Handler : handler,
    }
    
    err := server.ListenAndServe();
    
    if err != nil {
        panic(err)
    }
}