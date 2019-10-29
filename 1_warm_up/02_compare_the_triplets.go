package main

import (
    "fmt"
)

func main() {
    var arr1, arr2 []int 
    var temp int

    for i:=0; i<3; i++ {
        fmt.Scanf("%d", &temp)
        arr1 = append(arr1, temp)
    }
    
    for i:=0; i<3; i++ {
        fmt.Scanf("%d", &temp)
        arr2 = append(arr2, temp)
    }

    var res [2]int 
    res[0] = 0
    res[1] = 0
    for i:=0; i<3; i++ {
        if arr1[i] > arr2[i] {
            res[0] += 1
        } else if arr1[i] < arr2[i] {
            res[1] += 1
        }
    }

    fmt.Println(res[0], res[1])
}