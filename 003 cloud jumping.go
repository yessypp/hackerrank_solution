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
                    // fmt.Println("di sini ", data[1], i, in)
            if data[i] == 0 {
                if i == in -1 {
                    jump++
                    // fmt.Println("last one")
                } else if hitung == 0 {
                    hitung++
                } else if hitung > 0 {
                    jump++  //--------------
                    hitung = 0
                } else {
                    // jump++
                    // hitung = 0
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

10
0 0 1 0 0 0 0 1 0 0

7
0 0 1 0 0 1 0

2
0 0


