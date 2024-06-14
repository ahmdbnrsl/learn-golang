package main

func main() {
    name := "Via Fitriana"
    
    if name == "Via" {
        println("Her firstName")
    } else if name == "Fitriana" {
        println("Her lastName")
    } else if name == "Via Fitriana" {
        println("I Love You yeeayyy, hehe maaf ya")
    } else {
        println("I just want her alone")
    }
    
    //short statement 
    if length := len(name); length > 5 {
        println(name)
    }
}