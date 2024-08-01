package test 

import "testing"  
import "fmt"
import "io"
import "net/http"
import "net/http/httptest"

func SetCookie(writer http.ResponseWriter, request *http.Request) {
    cookie := new(http.Cookie)
    name := request.URL.Query().Get("name")
    if name != "" {
        cookie.Name = "Auth-ctx"
        cookie.Value = name
        cookie.Path = "/"
        http.SetCookie(writer, cookie)
        fmt.Fprint(writer, "cookie successfully generated")
    } else {
        fmt.Fprint(writer, "Cookie failed to generate")
    }
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
    cookie, err := request.Cookie("Auth-ctx")
    if err != nil {
        writer.WriteHeader(401)
        fmt.Fprint(writer, "Unauthorized")
    } else {
        writer.WriteHeader(200)
        fmt.Fprint(writer, cookie.Value)
    }
}

func Test12Cookie(t *testing.T) {
    router := http.NewServeMux()
    router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
        fmt.Fprint(writer, "Hello World!")
    })
    router.HandleFunc("/login", SetCookie)
    router.HandleFunc("/dashboard", GetCookie)
    
    server := http.Server{
        Addr : "localhost:8080",
        Handler : router,
    }
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}

func Test13CookieSetTest(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080/?name=via", nil)
    recorder := httptest.NewRecorder()
    
    SetCookie(recorder, request)
    
    cookies := recorder.Result().Cookies()
    
    for _, cookie := range cookies {
        fmt.Printf("Cookie : \n %s : %s \n", cookie.Name, cookie.Value)
    }
}

func Test14CookieGetTest(t *testing.T) {
    request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
    cookie := new(http.Cookie)
    cookie.Name = "Auth-ctx"
    cookie.Value = "Fitriana"
    cookie.Path = "/"
    request.AddCookie(cookie)
    
    recorder := httptest.NewRecorder()
    
    GetCookie(recorder, request)
    
    response := recorder.Result()
    body, _ := io.ReadAll(response.Body)
    fmt.Println(string(body))
}