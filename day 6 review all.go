package main
import (
    "fmt"
    "strings"
)

func main() {
    var count int
    fmt.Scan(&count)

    for j:=0; j<count; j++ {
        var line string
        fmt.Scan(&line)

        var left []string
        var right []string

        for i:=0; i<len(line); i++ {
            if i % 2 == 0 {
                left = append(left, string(line[i]))
            } else {
                right = append(right, string(line[i]))
            }
        } 

        fmt.Printf("%s %s\n", strings.Join(left, ""), strings.Join(right, ""))
    }
}

