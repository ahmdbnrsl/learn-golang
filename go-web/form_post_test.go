package test 

import "testing"  
import "fmt"
import "io"
import "net/http"
import "net/http/httptest"
import "strings"

func FormPost(writer http.ResponseWriter, request *http.Request) {
    form := request.ParseForm()
    if form != nil {
        panic(form)
    }
    
    firstName := request.PostForm.Get("firstname")
    lastName := request.PostForm.Get("lastname")
    
    fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func Test10PostForm(t *testing.T) {
    requestBody := strings.NewReader("firstname=Ahmad&lastname=Beni")
    request := httptest.NewRequest("POST", "http://localhost:8080", requestBody)
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    recorder := httptest.NewRecorder()
    
    FormPost(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}