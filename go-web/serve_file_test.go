package test  

import "net/http"
import "testing"
import "fmt"
import _"embed"

func ServeFile(writer http.ResponseWriter, request *http.Request) {
    if request.URL.Query().Get("name") != "" {
        http.ServeFile(writer, request, "./resources/ok.html")
    } else {
        http.ServeFile(writer, request, "./resources/notfound.html")
    }
}

func Test17ServeFile(t *testing.T) {
    server := http.Server{
        Addr : "localhost:8080",
        Handler : http.HandlerFunc(ServeFile),
    }
    server.ListenAndServe()
}

// USING GO EMBED 

//go:embed resources/ok.html
var ok string

//go:embed resources/notfound.html
var notFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
    if request.URL.Query().Get("name") != "" {
        fmt.Fprint(writer, ok)
    } else {
        fmt.Fprint(writer, notFound)
    }
}

func Test18ServeFile(t *testing.T) {
    server := http.Server{
        Addr : "localhost:8080",
        Handler : http.HandlerFunc(ServeFileEmbed),
    }
    server.ListenAndServe()
}