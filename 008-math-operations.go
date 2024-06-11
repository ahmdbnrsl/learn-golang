package main

func main() {
    a := 10
    b := 20
    
    c := a + b
    d := c * (a + b) / a
    println(c)
    println(d)
    
    //Augmented assignment 
    
    a += 20
    println(a)
    a -= 5
    println(a)
    a *= 10
    println(a)
    a /= 2
    println(a)
    a %= 3
    println(a)
}