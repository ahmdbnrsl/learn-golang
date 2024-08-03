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
type Page struct {
    Title string
    Name string
}
func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.ParseFiles("./template/name.gohtml"))
    // USING MAP
    /*t.ExecuteTemplate(writer, "name.gohtml", map[string]any{
        Title: "Via",
        Name: "Fitriana",
    })*/
    t.ExecuteTemplate(writer, "name.gohtml", Page{
        Title: "Via",
        Name: "Fitriana",
    })
}

func Test21TemplateDataStruct(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateDataStruct(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}
// type Person struct {
//     Name string
// }
func TemplateIfStatement(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.ParseFiles("./template/if.gohtml"))
    person := struct{
        Name string
    }{"Via Fitriana"}
    t.ExecuteTemplate(writer, "if.gohtml", person)
}

func Test22TemplateIfStatement(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateIfStatement(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

//===== OPERATOR PERBANDINGAN======
/*
eq artinya arg1 == arg2
ne artinya arg1 !== arg2
lt  artinya arg1 < arg2
le artinya arg1 <= arg2
gt artinya arg1 > arg2
ge artinya arg1 >= arg2
*/

func TemplateIfComparison(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.ParseFiles("./template/perbandingan.gohtml"))
    nilai := struct{
        Value int
    }{90}
    t.ExecuteTemplate(writer, "perbandingan.gohtml", nilai)
}

func Test23TemplateIfComparison(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateIfComparison(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

/* MENGAPA OPERATOR NYA DI DEPAN, KARENA ITU SAMA HALNYA MEMANGGIL FUNCTION 

MISAL ge arg1 80 maka ge(arg1, 80)
*/

func TemplateRange(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.ParseFiles("./template/range.gohtml"))
    fav := struct{
        Favourite []string
    }{[]string{"Via", "Fitirana", "Via Fitriana"}}
    t.ExecuteTemplate(writer, "range.gohtml", fav)
}
func Test24TemplateRange(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateRange(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}

func TemplateWith(writer http.ResponseWriter, request *http.Request) {
    t := template.Must(template.ParseFiles("./template/with.gohtml"))
    type Addr struct{ Street, District string }
    person := struct{
        Name string
        Address Addr
    }{"Via Fitriana", Addr{"jln surga", "Sidareja"}}
    t.ExecuteTemplate(writer, "with.gohtml", person)
}

func Test25TemplateWith(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080", nil)
    recorder := httptest.NewRecorder()
    
    TemplateWith(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}