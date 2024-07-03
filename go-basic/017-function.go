package main

import "fmt"

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

//variadic function
func sumAll(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total 
}

//Function value
func sayGoodBye(name string) string {
    return "Good bye " + name
}

//function as parameter
func plus(numbers ...int) int {
    hasil := 0
    for _, number := range numbers {
        hasil += number
    }
    return hasil
}
func kali(numbers ...int) int {
    hasil := 1
    for _, number := range numbers {
        hasil *= number
    }
    return hasil
}

type Math func(...int) int // type declaration

func math(pluss, kalis Math, numbers ...int) (int, int) {
    fmt.Println(numbers)
    return pluss(numbers...), kalis(numbers...)
}

//anonymous function
type Blacklist func(string) bool

func register(name string, blacklist Blacklist) {
    if blacklist(name) {
        println("hehe love you")
    } else {
        println("who are you?")
    }
}

//recursive function
func factorial(value int) int {
    if value == 1 {
        println(1)
        return 1
    } else {
        result := value * factorial(value - 1)
        println(result)
        return result
    }
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
    //Named return value
    firstName, lastName := Name()
    println(firstName, lastName)
    //variadic function
    println(sumAll(1, 2, 3, 4, 5))
    //convert slice to vararg parameter 
    isSlice := []int{
        20,
        50,
        50,
    }
    hasil := sumAll(isSlice...) // like spread operator in javascript:v
    println(hasil)
    //Function value
    goodBye := sayGoodBye
    println(goodBye("Beni"))
    //Function as parameter
    println(math(plus, kali, 10, 10))
    //anonymous function 
    register("Via", func(name string) bool {
        return name == "Via"
    })
    //recursive function
    factorial(10)
    //clousure
    count := 0
    increment := func() {
        println(count)
        count++
    }
    
    increment()
    increment()
    println(count)
}