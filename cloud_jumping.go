// Cloud Jumping
// problem: https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem


package main
import "fmt"

func main() {
    var in int
    fmt.Scanln(&in)

    var data []int
    data = make([]int, in)

    jump := 0
    hitung := 0
    for i := 0; i < in; i++ {
        fmt.Scanf("%d", &data[i])
        
        if i > 0 {
            if data[i] == 0 {
                if i == in -1 {
                    jump++
                } else if hitung == 0 {
                    hitung++
                } else if hitung > 0 {
                    jump++  
                    hitung = 0
                } 
            } else if data [i] == 1 && hitung == 0 {
                hitung++
            } else {
                jump++
            }
        }

    }

    fmt.Println(jump)
}


