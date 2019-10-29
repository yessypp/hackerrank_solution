package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int {
    panjang := int32(len(arr))
    // var arr2 []int32
    var tempPos int32
    var tempVal int32
    flag := false
    counter := 0

    for {
        i := int32(1)
        for i=0; i<panjang; i++ {
            flag = false
            fmt.Println(i, arr[i])

            if arr[i] == i+1 {
                fmt.Println("sama")
                continue
            }

            if arr[i] == tempPos && flag {
                arr[i] = tempVal
                fmt.Println("tukarrr")
                fmt.Println(arr)
            }

            if !flag && arr[i] != i+1 {
                tempVal = arr[i]
                tempPos = i+1
                arr[i] = i+1
                flag = true
                fmt.Println("beda", tempVal, tempPos)
                fmt.Println(arr)
            }
        }

        counter++
        fmt.Println("counter = ", counter)

        //counter all are set?
        if i == panjang && flag == false {
            break
        }
    }

    return counter

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

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

    res := minimumSwaps(arr)

    fmt.Fprintf(writer, "%d\n", res)

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
