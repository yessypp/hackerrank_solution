package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
    var counter int

    for i := len(q)-1; i > 0 ; i-- {
        if q[i-1] > q[i] {
            a := q[i-1] - q[i]
            if a > 2 {
                counter = -1
                break
            } else {
                counter += int(a)
            }
        }
        // fmt.Println(i, q[i], counter)
    }

    if counter < 0 {
        fmt.Println("Too chaotic")
    } else {
        fmt.Println(counter)
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
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
