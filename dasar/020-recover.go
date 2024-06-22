package main

func end() {
    println("App has ended")
    message := recover()
    println("Something went wrong", message)
}
func run(error bool) {
    defer end()
    if error {
        panic("ERROR")
    }
}

func main() {
    run(true)
    println("Via Fitriana")
}