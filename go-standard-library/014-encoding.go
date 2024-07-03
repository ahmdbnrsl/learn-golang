package main 

import (
    "encoding/base64"
    "encoding/csv"
    "strings"
    "io"
    "fmt"
    "os"
)

func main() {
    // BASE 64
    var encoded = base64.StdEncoding.EncodeToString([]byte("Via Fitriana"))
    fmt.Println(encoded)
    
    var decoded, err = base64.StdEncoding.DecodeString(encoded)
    
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(string(decoded))
    }
    // CSV READER
    csvString := "Via,Fitriana,Beni\n" +
    "budi,wahyu,arnaf\n" +
    "chindy,tyan,aguss"
    
    reader := csv.NewReader(strings.NewReader(csvString))
    
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        
        fmt.Println(record)
    }
    // CSV WRITER
    writer := csv.NewWriter(os.Stdout)
    _ = writer.Write([]string{"Via", "Fitriana"})
    _ = writer.Write([]string{"Beni", "Rusli"})
    
    writer.Flush()
}