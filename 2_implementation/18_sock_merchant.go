// Sock Merchant
// problem: https://www.hackerrank.com/challenges/sock-merchant/problem

package main
import "fmt"

func main() {
    var in, total int
    fmt.Scanf("%d\n", &in)

	// array
    var data []int
    data = make([]int, in)

	// map
    pair := map[int]int{}

    for i := 0; i < in; i++ {
        fmt.Scanf("%d", &data[i])

		// searching inside a map - apakah di dalam map pair ada key berupa data[i] ?
        if _, ok := pair[data[i]]; !ok{
           pair[data[i]] = 1 
        } else {
            pair[data[i]]++
        }

        if pair[data[i]] %2 == 0 {
            total++
        }
    } 

    fmt.Println(total)
}
