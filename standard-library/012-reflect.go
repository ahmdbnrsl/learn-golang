package main 

import (
    "fmt"
    "reflect"
)

type Sample struct{
    Name string
}

/*STRUCT TAG*/
type Person struct{
    Name string `required:"true" max:"18"`
    Address string `required:"true" max:"18"`
}

func IsValid(data any) (result bool) {
    t := reflect.TypeOf(data)
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        if field.Tag.Get("required") == "true" {
            result = reflect.ValueOf(data).Field(i).Interface() != ""
            if result == false {
                return result
            }
        }
    }
    return result
}

func main() {
    sample := Sample{"Via Fitrianaa"}
    sampleType := reflect.TypeOf(sample)
    structField := sampleType.Field(0)
    
    fmt.Println(structField)
    //GET STRUCT TAG
    person := Person{"Via Fitrianaa", ""}
    personType := reflect.TypeOf(person)
    personField := personType.Field(0);
    required := personField.Tag.Get("required")
    fmt.Println(required)
    
    fmt.Println(IsValid(person))
}