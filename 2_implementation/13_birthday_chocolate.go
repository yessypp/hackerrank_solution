package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the birthday function below.
func birthday(s []int32, d int32, m int32) int32 {
    var count int32
    length := int32(len(s))

    if len(s) > 1 {
        for i:= 0; int32(i) <= length-m; i++ {
            fmt.Println("new i = ", i)
            sum := s[i]
            limit := int32(i) + m

            for j:=int32(i+1); int32(j) < limit; j++ {
                fmt.Println("another j = ", j)
                fmt.Println("> ", s[j])
                sum += s[j]
            }

            if sum == d {
                count++
            }
        }
    } else if s[0] == d {
        count++
    }

    return count

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var s []int32

    for i := 0; i < int(n); i++ {
        sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
        checkError(err)
        sItem := int32(sItemTemp)
        s = append(s, sItem)
    }

    dm := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    dTemp, err := strconv.ParseInt(dm[0], 10, 64)
    checkError(err)
    d := int32(dTemp)

    mTemp, err := strconv.ParseInt(dm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    result := birthday(s, d, m)

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
