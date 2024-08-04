package test  

import "html/template"
import "net/http"
import "net/http/httptest"
import "io"
import "testing"  
import "fmt"
import "strings"
import _"embed"


func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.ParseFiles("./template/header.gohtml", "./template/footer.gohtml", "./template/layout.gohtml"))
    person := struct{
        Title, Name string
    }{"Via", "Fitriana"}
    t.ExecuteTemplate(writer, "layout.gohtml", person)
}

func Test26TemplateLayout(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateLayout(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

type Person struct{
    Name string
}
func (person Person) SayHello(name string) string {
    return "Hello " + name + " My name is " + person.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.ParseFiles("./template/function.gohtml"))
    t.ExecuteTemplate(writer, "function.gohtml", Person{"Via Fitriana"})
}

func Test27TemplateFunction(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateFunction(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

func TemplateGlobalFunction(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
    t.ExecuteTemplate(writer, "FUNCTION", struct{Name string}{"Beni"})
}

func Test28TemplateGlobalFunction(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateGlobalFunction(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

type Human struct{
    Name string
}
func (human Human) Upper(name string) string {
    return strings.ToUpper(name)
}
func TemplateCreateGlobalFunction(
    writer http.ResponseWriter,
    request *http.Request,
) {
    human := Human{"Via"}
    t := template.New("FUNCTION")
    t = t.Funcs(map[string]any{
        "upper" : func(name string)string {
            return strings.ToUpper(name)
        },
    })
    t = template.Must(t.Parse(`{{upper .Name}}`))
    t.ExecuteTemplate(writer, "FUNCTION", human)
}

func Test29TemplateDefineGlobalFunction(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateCreateGlobalFunction(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

func TemplatePipelinesFunction(writer http.ResponseWriter, request *http.Request) {
    human := Human{"Via"}
    t := template.New("FUNCTION")
    t = t.Funcs(map[string]any{
        "upper" : func(name string)string {
            return strings.ToUpper(name)
        },
        "sayHello" : func(name string)string {
            return "Hello " + name
        },
    })
    t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))
    t.ExecuteTemplate(writer, "FUNCTION", human)
}

func Test30TemplatePipelinesFunc(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplatePipelinesFunction(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}