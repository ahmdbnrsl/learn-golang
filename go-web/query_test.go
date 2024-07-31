package test 

import "testing"  
import "fmt"
import "io"
import "net/http"
import "net/http/httptest"
import "strings"

func SayHello(writer http.ResponseWriter, request *http.Request) {
    name := request.URL.Query().Get("name")
    if name == "" {
        fmt.Fprint(writer, "Hello")
    } else {
        fmt.Fprintf(writer, "Hello %s", name)
    }
}

func Test5Query(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080/?name=Beni", nil)
    recorder := httptest.NewRecorder()
    
    SayHello(recorder, request)
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

func MultipleParamQuery(writer http.ResponseWriter, request *http.Request) {
    firstName := request.URL.Query().Get("firstname")
    lastName := request.URL.Query().Get("lastname")
    fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func Test6MultipleQueryParam(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080/?firstname=Beni&lastname=Rusli", nil)
    recorder := httptest.NewRecorder()
    
    MultipleParamQuery(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

func MultipleValueQuery(writer http.ResponseWriter, request *http.Request){
    queryValue := request.URL.Query()
    name := queryValue["name"]
    fmt.Fprintln(writer, strings.Join(name, ","))
}

func Test7MultipleValueQuery(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080/?name=Beni&name=Rusli", nil)
    recorder := httptest.NewRecorder()
    
    MultipleValueQuery(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}