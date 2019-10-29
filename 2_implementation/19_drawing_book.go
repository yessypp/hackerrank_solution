package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the pageCount function below.
 */
func pageCount(n int32, p int32) int32 {
    var countdepan, countbel int32

    for i:=0; int32(i)<=n; i++ {
        countdepan++
        if p == 1 {
            break
        }
        if int32(i) == p || int32(i+1) == p {
            break
        }
        i++
    }

    for j:=n; int32(j)>0; j-- {
        countbel++
        if j%2 != 0 {
            j--
        }

        // fmt.Println(">> j", j)
        // fmt.Println("> n", n)

        if j == p || j+1 == p{
            // fmt.Println("> break")
            break
        }
    }

    if countdepan > 0 { countdepan-- }
    if countbel > 0 { countbel-- }

    // fmt.Println(countdepan, countbel)
    if countdepan < countbel {
        return countdepan
    } 
    return countbel
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

    pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    p := int32(pTemp)

    result := pageCount(n, p)

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
