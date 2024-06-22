package helper

var version = "1.0.0" //cannot access in another package but can access in same package
var App = "golang"  //can access in another package

func goodBye(name string) string {
    return "goodbye " + name
}

func SayHello(name string) string {
    return "hello " + name
}