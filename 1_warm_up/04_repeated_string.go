// Repeated String
// problem: https://www.hackerrank.com/challenges/repeated-string/problem

package main
import "fmt"
import "strings"

func main() {
    var data string
    fmt.Scanln(&data)

    var in int
    fmt.Scanln(&in)

    count := strings.Count(data, "a")
    count2 := 0

    l := len(data)
    pengkali, sisa, res := 0, 0, 0
    sisa = in % l
    if in > l {
        pengkali = in/l
    }

    for i, _ := range data {
        if i < sisa {
            if data[i] == 'a' {
                count2++
            }
        }
    }

    res = (pengkali * count) + count2
    fmt.Println(res)
}
