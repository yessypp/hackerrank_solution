package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
    var a, b, c int

    for _, v := range arr {
        if v > 0 {
            a++
        } else if v < 0 {
            b++
        } else {
            c++
        }
    }

    a2 := float64(a) / float64(len(arr))
    b2 := float64(b) / float64(len(arr))
    c2 := float64(c) / float64(len(arr))
    
    fmt.Printf("%.5f\n", a2)
    fmt.Printf("%.5f\n", b2)
    fmt.Printf("%.5f\n", c2)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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
