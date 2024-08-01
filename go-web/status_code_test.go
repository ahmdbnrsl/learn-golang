package test 

import "testing"  
import "fmt"
import "io"
import "net/http"
import "net/http/httptest"

func StatusCode(writer http.ResponseWriter, request *http.Request) {
    name := request.URL.Query().Get("name")
    if name != "" {
        writer.WriteHeader(200)
        fmt.Fprint(writer, name)
    } else {
        writer.WriteHeader(400)
        fmt.Fprint(writer, "kosong")
    }
}

func Test11StatusCode(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    StatusCode(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(response.StatusCode)
    fmt.Println(response.Status)
    fmt.Println(string(body))
}