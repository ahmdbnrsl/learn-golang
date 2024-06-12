package main

func main() {
    
    //string
    
    str := "Via"
    str1 := "Via"
    str2 := "Fitriana"
    str3 := "a"
    str4 := "b"
    
    println(str == str2)//false
    println(str == str1)//true
    
    println(str != str2)//true 
    println(str != str1)//false
    
    println(str3 < str4)
    
    //number
    
    const(
        number1 int16 = 456
        number2 int16 = 456
        number3 int16 = 20
        number4 int16 = 458
    )
    
    println(number1 == number2)
    println(number1 != number3)
    println(number1 > number3)
    println(number3 < number1)
    println(number4 >= number2)
    println(number3 <= number4)
}