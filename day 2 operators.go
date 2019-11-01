package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
    var tip, tax, sum float64

    tip = meal_cost * float64(tip_percent) / 100
    tax = meal_cost * float64(tax_percent) / 100

    sum = meal_cost + tip + tax

    sisa := math.Mod(sum, 1)
    if sisa > 0.5 {
        sum++
    }
    fmt.Println(int64(sum))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
    checkError(err)

    tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tip_percent := int32(tip_percentTemp)

    tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tax_percent := int32(tax_percentTemp)

    solve(meal_cost, tip_percent, tax_percent)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
