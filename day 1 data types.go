package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var in int
    var do float64
    var st string
    
    // Read and save an integer, double, and String to your variables.
    fmt.Scanf("%d\n", &in)
    fmt.Scanf("%f\n", &do)
    for scanner.Scan() {
        st = scanner.Text()
    }

    // Print the sum of both integer variables on a new line.
    fmt.Println(i + uint64(in))
    
    // Print the sum of the double variables on a new line.
    sum := d + do
    fmt.Printf("%.1f\n", sum)
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Printf("%s%s", s, st)

}