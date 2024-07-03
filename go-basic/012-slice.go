package main

import "fmt"

func main() {
    sholat := [5]string{
        "shubuh",
        "dzuhur",
        "ashar",
        "maghrib",
        "isya",
    }
    
    slice := sholat[:]
    
    fmt.Println(slice)
    println(slice[0])
    
    bolong := sholat[1:4]
    fmt.Println(bolong) //[dzuhur, ashar, maghrib]
    bolong[0] = "Shubuh"
    fmt.Println(sholat) //[shubuh, shubuh, ashar, maghrib, isya]
    
    //function slice 
    
    //append
    days := [...]string{
        "senin",
        "Selasa",
        "Rabu",
        "kamis",
        "jumat",
        "sabtu",
        "minggu",
    }
    
    daySlice1 := days[5:]
    daySlice1[0] = "Sabtu Baru"
    daySlice1[1] = "Minggu Baru"
    fmt.Println(days) //[senin, Selasa, Rabu, kamis, jumat, Sabtu Baru, Minggu Baru]
    
    daySlice2 := append(daySlice1, "Libur Baru")
    daySlice2[0] = "Ups"
    fmt.Println(daySlice2) // [Ups, Sabtu Baru, Minggu Baru]
    fmt.Println(days) // [senin, Selasa, Rabu, kamis, jumat, Sabtu Baru, Minggu Baru]
    
    daySlice3 := days[1:4]
    fmt.Println(daySlice3)
    daySlice4 := append(daySlice3, "Jumu'ah")
    fmt.Println(daySlice4)
    fmt.Println(days) // [senin, Selasa, Rabu, kamis, Jumu'ah, Sabtu Baru, Minggu Baru]
    
    //make
    
    newSlice := make([]int8, 2, 20)
    fmt.Println(newSlice)
    
    newSlice[0] = 20
    newSlice[1] = 30
    fmt.Println(newSlice)
    newSlice2 := append(newSlice, 40)
    fmt.Println(newSlice2)
    newSlice2[0] = 10
    fmt.Println(newSlice)
    
    //copy
    
    fromSlice := days[:]
    toSlice := make([]string, len(fromSlice), cap(fromSlice))
    
    copy(toSlice, fromSlice)
    fmt.Println(toSlice)
    
    //different between array and slice
    arr := [...]int{
        1,
        2,
        3,
    }
    slicess := []int{
        1, 
        2, 
        3,
    }
    fmt.Println(arr)
    fmt.Println(slicess)
}