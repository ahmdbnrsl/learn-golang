package test  

import "html/template"
import "net/http"
import "net/http/httptest"
import "io"
import "testing"  
import "fmt"
import "embed"


func TemplateFileString(writer http.ResponseWriter, request *http.Request) {
    templateText := `<html><body>{{.}}</body></html>`
    t, err := template.New("SIMPLE").Parse(templateText)
    if err != nil {
        panic(err)
    }
    t.ExecuteTemplate(writer, "SIMPLE", "HELLO WORLD!!")
}

func Test19TemplateString(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateFileString(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}


//go:embed template/*.gohtml
var templates embed.FS
func TemplateFromFile(writer http.ResponseWriter, request *http.Request) {
   // t, err := template.ParseFiles("./template/simple.gohtml")
   //template from directory 
    //t, err := template.ParseGlob("./template/*.gohtml")
    //USING GO EMBED
    t, err := template.ParseFS(templates, "template/*.gohtml")
    if err != nil {
        panic(err)
    }
    t.ExecuteTemplate(writer, "simple.gohtml", "Via Fitriana")
}

func Test20TemplateFiles(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateFromFile(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}