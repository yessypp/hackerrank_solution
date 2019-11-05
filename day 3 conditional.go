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

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int(NTemp)

    check(N)
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

func check(a int) {
    if math.Mod(float64(a), float64(2)) != float64(0) { //odd
        fmt.Println("Weird")
    } else {
        if a >= 2 && a <= 5 {
            fmt.Println("Not Weird")
        } else if a>=6 && a <= 20 {
            fmt.Println("Weird")
        } else if a >= 20 {
            fmt.Println("Not Weird")
        }
    } 
}


