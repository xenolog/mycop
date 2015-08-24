package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

const (
    URL = "http://realmeteo.ru/moscow/1/charts.json"
)

type RealMeteoType struct {
    Start    time.Duration //`json:start`
    Interval time.Duration //`json:interval`
    Data     struct {
        Temperature []float32
        Pressure    []float32
        Humidity    []float32
    }
}

var (
    Rmeteo *RealMeteoType
    err    error
)

func main() {
    var ts time.Duration
    resp, err := http.Get(URL)
    if err != nil {
        fmt.Printf("Error while downloading URL '%s':%v", URL, err)
        os.Exit(1)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error while processing response from URL '%s':%v", URL, err)
        os.Exit(1)
    }
    Rmeteo = new(RealMeteoType)
    err = json.Unmarshal(body, &Rmeteo)
    if err != nil {
        fmt.Printf("Error while json decoding from URL '%s':%v", URL, err)
        os.Exit(1)
    }
    // fmt.Printf("date, Temperature, 0, Pressure, 760\n")
    for i, _ := range Rmeteo.Data.Temperature {
        ts = Rmeteo.Start/1000 + time.Duration(i)*Rmeteo.Interval
        fmt.Printf("%10d %.2f 0 %.2f 760\n", ts, Rmeteo.Data.Temperature[i], Rmeteo.Data.Pressure[i])
    }
}
