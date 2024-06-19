package main

func end() {
    println("App has ended")
}
func run(error bool) {
    defer end()
    if error {
        panic("ERROR")
    }
}

func main() {
    run(true)
}