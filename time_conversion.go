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
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
    s1 := strings.Split(s, ":")

    hour := s1[0]
    min := s1[1]

    a := []rune(s1[2])
    sec := string(a[0:2])
    sec2 := string(a[2:4])

    malam := false
    if string(sec2) == "PM" {
        malam = true
        fmt.Println("malam = ", malam)
    }

    hourInt, _ := strconv.Atoi(hour)

    if malam {
        hourInt += 12
        
        if hourInt == 24 {
            hourInt = 0
        }
    } else if !malam {
        if hourInt == 12 {
            hourInt = 0
        }
    }

    var hour2 string
    if hourInt < 10 {
        hour2 = strconv.Itoa(hourInt)
        hour2 = "0" + hour2
    } else {
        hour2 = strconv.Itoa(hourInt)
    }

    result := hour2 + ":" + min + ":" + sec
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

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
