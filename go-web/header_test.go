package test 

import "testing"  
import "fmt"
import "io"
import "net/http"
import "net/http/httptest"

func GetHeader(writer http.ResponseWriter, request *http.Request) {
    contentType := request.Header.Get("Content-type")
    fmt.Fprint(writer, contentType)
}

func Test8Header(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
    request.Header.Add("Content-type", "application/json")
    recorder := httptest.NewRecorder()
    
    GetHeader(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

func SendHeader(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Add("X-Powered-By", "Ahmad Beni Rusli")
    fmt.Fprint(writer, "Ok")
}

func Test9SendHeader(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    response := httptest.NewRecorder()
    
    SendHeader(response, request)
    
    poweredBy := response.Header().Get("X-Powered-By")
    fmt.Println(poweredBy)
}