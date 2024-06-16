package main

func say() {
    println("Beni")
}

//function parameter 
func sayName(name string, age int8) {
    println(name, age)
}

//function return value
func sum(num1 int, num2 int) int {
    return num1 + num2
}

//return multiple value
func operator(num1 int, num2 int) (int, int) {
    sums := num1 + num2
    dec := num1 - num2
    return sums, dec
}

//Named return value
func Name() (firstName, lastName string) {
    firstName = "Via"
    lastName = "Fitriana"
    return firstName, lastName
}

func main() {
    say()
    sayName("Via Fitriana", 20)
    println(sum(10, 20))
    //multiple return value
    sums, dec := operator(20, 10)
    println(sums, dec)
    //to ignore value
    _, kurang := operator(20, 1)
    println(kurang)
    tambah, _ := operator(30, 20)
    println(tambah)
    
    firstName, lastName := Name()
    println(firstName, lastName)
}