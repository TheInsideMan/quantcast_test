package main

import (
    "fmt"
    "flag"
    "testing"
    controllers "QuantCast/controllers"
)

func TestEcho(t *testing.T) {
    // Test happy path
    err := echo([]string{"bin-name", "hello", "world!"})
    require.NoError(t, err)
}

func TestEchoErrorNoArgs(t *testing.T) {
    // Test empty arguments
    err := echo([]string{})
    require.Error(t, err)
}

// I could build a test to check if there is a '.csv' file in the root

func main() {
    // retrieve the command line arguments
    dateCommandStringValue := flag.String("d", "", "help message for dateCommandValue")
    fileCommandValue := flag.String("f", "", "help message for file")
    flag.Parse()
    
    // Open the CSV file specified in command and return as a array slice
    csvLines,err := controllers.OpenCsvFile(*fileCommandValue)
    if err != nil {
        return
    }
    
    // convert the date string into a time object
    // `layout` indicates to golang how it should parse the string into a time obj
    specifiedDay,err := controllers.TimeConvert("2006-01-02", *dateCommandStringValue)
    if err != nil {
        return
    }

    // create a map for storing only the most active cookies for the specific day
    summedCookieMap := controllers.GetCookieMap(csvLines,specifiedDay )

    // print out the most active cookies!
    for k, _ := range summedCookieMap {
        fmt.Println(k)
        // print with the number of occourances as well
        // fmt.Println(k +" - "+ strconv.Itoa(v))
    }
}