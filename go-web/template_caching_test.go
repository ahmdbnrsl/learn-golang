package test  

import "html/template"
import "net/http"
import "net/http/httptest"
import "io"
import "testing"  
import "fmt"
import "embed"

//go:embed template/*.gohtml
var templatess embed.FS
var myTemplates = template.Must(template.ParseFS(templatess, "template/*.gohtml"))
func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
    myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello World!")
}

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
   // myTemplates.ExecuteTemplate(writer, "post.gohtml", struct{Body string}{"<h1>Beni</h1>"})
   //disable auto escape
   /*myTemplates.ExecuteTemplate(writer, "post.gohtml", struct{Style, Body any}{template.CSS(`
   * {
       color: blue;
   }
   `), template.HTML("<h1>Beni</h1>")})*/
   // XSS EXAMPLE
   myTemplates.ExecuteTemplate(writer, "post.gohtml", struct{Style, Body any}{template.CSS(`
   * {
       color: blue;
   }
   `), template.HTML(request.URL.Query().Get("name"))})
}

func Test31Server(t *testing.T) {
    request := httptest.NewRequest("GET", `http://localhost:8080?name=<a%20href="https://wa.me/status">click</a>`, nil)
    recorder := httptest.NewRecorder()
    
    //TemplateCaching(recorder, request)
    TemplateAutoEscape(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}