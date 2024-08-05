package test  

import "testing"
import "net/http"
import "fmt"

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprint(writer, "Hello World!")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func Test32Redirect(t *testing.T) {
    router := http.NewServeMux()
    router.HandleFunc("/", RedirectFrom)
    router.HandleFunc("/redirec-to", RedirectTo)
    
    server := http.Server{
        Addr: "localhost:8080",
        Handler: router,
    }
    
    server.ListenAndServe()
}