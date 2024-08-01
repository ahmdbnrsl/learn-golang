package test 

import "testing"  
import _"fmt"
import "io/fs"
import "net/http"
import _"net/http/httptest"
import "embed"
import _"embed"

func Test15FileServer(t *testing.T) {
    dir := http.Dir("./resources")
    fileServer := http.FileServer(dir)
    router := http.NewServeMux()
    router.Handle("/static/", http.StripPrefix("/static", fileServer))
    
    server := http.Server{
        Addr : "localhost:8080",
        Handler : router,
    }
    
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}

//go:embed resources
var resources embed.FS

func Test16FileServerEmbed(t *testing.T) {
    //fileServer := http.FileServer(http.FS(resources))
    dir, _ := fs.Sub(resources, "resources")
    fileServer := http.FileServer(http.FS(dir))
    router := http.NewServeMux()
    router.Handle("/static/", http.StripPrefix("/static", fileServer))
    
    server := http.Server{
        Addr : "localhost:8080",
        Handler : router,
    }
    
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}