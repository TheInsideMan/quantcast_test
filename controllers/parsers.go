package controllers

import (
    "os"
    "encoding/csv"
    "time"
)


func TimeConvert(layout string,strTime string) (time.Time, error) {
    // layout := "2006-01-02" 
    timeArgument, err := time.Parse(layout, strTime)
    if err != nil {
        return timeArgument, err
    }
    return timeArgument, err
}

func OpenCsvFile(fileCommandValue string) ([][]string,error) {
    // retrieve the csv file
    var emptyArray [][]string
    csvFile, err := os.Open(fileCommandValue)
    if err != nil {
        return emptyArray, err
    }
    defer csvFile.Close()
    // load csv contents into memory
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        return emptyArray, err
    }
    return csvLines,err
}

func GetCookieMap(csvLines [][]string,timeArgument time.Time) map[string]int {
    // create a golang map for storing the cookie objects and how many times they occour in csv
    summedCookieMap := make(map[string]int)
    for _, line := range csvLines {
        // convert the csv row date to a time object 
        csvRowDay,err := TimeConvert("2006-01-02T15:04:05+00:00",line[1])
        if err != nil {
            continue
        }
        
        // skip any dates that do not match calendar day `d` from terminal command
        if csvRowDay.Day() != timeArgument.Day() {
            continue   
        }
        // either insert or +1 the value of cookie in map
        summedCookieMap[line[0]] = summedCookieMap[line[0]]+1
    }
    // remove from `summedCookieMap` the least occuring cookies on the given day
    previousKey := ""
    previousValue := 1
    highestValue := 1
    // check if summedCookieMap is empty
    for k, v := range summedCookieMap {
        if v < highestValue {
            delete(summedCookieMap, k);
        } else {
            highestValue = v
            if previousValue < v {
                delete(summedCookieMap, previousKey);
            }
            previousKey = k
            previousValue = v
        }
    }
    return summedCookieMap
}