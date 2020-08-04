package main

import (
    "fmt"
)
import . "./package_test"

func main() {
    var sum int
    sum = Sum_test(2, 3)
    fmt.Printf("sum: %d; Total_sum: %d\n", sum, Total_sum)
}
