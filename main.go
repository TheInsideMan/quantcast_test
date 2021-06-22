package main

import (
    "fmt"
    "flag"
    "os"
    "encoding/csv"
    "time"
    "controllers"
)

// I could build a test to check if there is a '.csv' file in the root

func main() {
    // retrieve the command line arguments
    dateCommandStringValue := flag.String("d", "", "help message for dateCommandValue")
    fileCommandValue := flag.String("f", "", "help message for file")
    flag.Parse()
    
    // convert the d argument string into a date object
    // `layout` indicates to golang how it should parse the string into time obj
    timeArgument,err := TimeConvert("2006-01-02",*dateCommandStringValue)
    if err != nil {
        return
    }
    csvLines,err := OpenCsvFile(*fileCommandValue)
    if err != nil {
        fmt.Println(err)
        return
    }

    summedCookieMap := GetCookieMap(csvLines,timeArgument )

    // print out all the remaining cookie winners!
    for k, _ := range summedCookieMap {
        fmt.Println(k)
        // fmt.Println(k +" "+ strconv.Itoa(v))
    }
}