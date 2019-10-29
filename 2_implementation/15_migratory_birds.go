package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {
    var res [6]int32

    for _, v := range arr {
        if v == int32(1) {
            res[1]++
        } else if v == int32(2) {
            res[2]++
        } else if v == int32(3) {
            res[3]++
        } else if v == int32(4) {
            res[4]++
        } else if v == int32(5) {
            res[5]++
        }
    }

    bigfreq := int32(0)
    smallnum := int32(0)
    //get the bigger count
    for i:=1; i<=5; i++ {
        if res[i] > bigfreq {
            bigfreq = res[i]
            smallnum = int32(i)
        } 
    }

    return smallnum
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(arrCount); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := migratoryBirds(arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
