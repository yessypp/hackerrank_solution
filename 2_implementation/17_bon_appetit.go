package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the bonAppetit function below.
func bonAppetit(bill []int32, k int32, b int32) {
    var cost int32

    for key, v := range bill {
        if int32(key) != k {
            cost += v
        }
    }

    sisa := b - (cost/2)
    if sisa == 0 {
        fmt.Println("Bon Appetit")
    } else {
        fmt.Println(sisa)
    }

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nk := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    billTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var bill []int32

    for i := 0; i < int(n); i++ {
        billItemTemp, err := strconv.ParseInt(billTemp[i], 10, 64)
        checkError(err)
        billItem := int32(billItemTemp)
        bill = append(bill, billItem)
    }

    bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    b := int32(bTemp)

    bonAppetit(bill, k, b)
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
