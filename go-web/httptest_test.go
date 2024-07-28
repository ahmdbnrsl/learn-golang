package test  

import "testing"
import "net/http"
import "net/http/httptest"
import "fmt"
import "io"

func SayHelloHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprint(writer, "Hello World")
}

func Test4HttpTest(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
    recorder := httptest.NewRecorder()
    
    SayHelloHandler(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    
    fmt.Println(response.StatusCode)
    fmt.Println(response.Status)
    fmt.Println(string(body))
}