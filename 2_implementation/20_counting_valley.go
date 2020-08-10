// Counting Valley
// problem: https://www.hackerrank.com/challenges/counting-valleys/problem

package main
import "fmt"

func main() {
    var in int
    fmt.Scanln(&in) //read input without format

    var data string
    fmt.Scanln(&data)

    level, valley := 0, 0
    
    for i, _ := range data{     
    //looping from string into each char/alphabet
        if data[i] == 'U' {
            if level < 0 && level+1 == 0{
                valley++
            }
            level++
        } else {
            level--
        }
    }

    fmt.Println(valley)
}