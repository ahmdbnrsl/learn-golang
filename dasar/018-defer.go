package main

func logging() {
    println("Finish")
}

func run() {
    defer logging()
    println("running...")
}
func main() {
    run()
}