package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
    // var days []int
    days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    daysjulian := []int{29, 28, 31, 29, 31, 29, 31, 29, 29, 31, 29, 29}

    if year >= 1919 && year <= 2700 { //gregorian
        if year % 400 == 0 {
            days[1] = 29
        }
        if year % 4 == 0 && year % 100 != 0 {
            days[1] = 29
        }
    } else if year <= 1918 && year >= 1700 {
        days = daysjulian
        if year % 4 == 0 {
            days[1] = 29
        }
        if year == 1918 {
            days[1] = 14
        }
    }

    sisa := 256
    month := 0
    
    for i:=0; i<11; i++ {
        month++
        sisa -= days[i]

        // fmt.Println(month, sisa)
        if sisa <= days[i+1] {
            break
        }
    }

    if year == 1918 && month+1 == 2 {
        sisa += 14
    }

    return fmt.Sprintf("%02d.%02d.%d", sisa, month+1, year)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    year := int32(yearTemp)

    result := dayOfProgrammer(year)

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
