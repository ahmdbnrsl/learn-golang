package main

func main() {
    bool1 := true
    bool2 := false
    bool3 := true
    bool4 := false
    
    //&& (and)
    
    println(bool1 && bool3) //true
    println(bool2 && bool4) //false
    println(bool1 && bool2) //false
    println(bool4 && bool3) //false 
    
    //|| (or)
    
    println(bool1 || bool3) //true
    println(bool2 || bool4) //false
    println(bool1 || bool2) //true
    println(bool4 || bool3) //true 
    
    //example
    
    const(
        nilaiAbsensi int8 = 90
        nilaiAkhir int8 = 80
    )
    lulusNilaiAkhir := nilaiAkhir >= 80
    lulusAbsensi := nilaiAbsensi >= 80
    
    lulus := lulusNilaiAkhir && lulusAbsensi
    
    println(lulus)
}