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

// Complete the utopianTree function below.
func utopianTree(n int32) int32 {
    var result int32
    result = 1
    
    for i:=int32(0); i < n; i++ {
        if math.Mod(float64(i), float64(2)) == 0 {
            result = result * 2        
        } else {
            result = result + 1
        }
    }
    
    if n < 1 {
        return 1
    }
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        result := utopianTree(n)

        fmt.Fprintf(writer, "%d\n", result)
    }

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
