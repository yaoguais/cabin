package main

import (
    "fmt"
)

func main() {
    num := []int{2, 7, 11, 15}
    fmt.Printf("%v\n", twoSum(num, 9))
    fmt.Printf("%v\n", twoSum(num, 30))
}

func twoSum(nums []int, target int) []int {
    l := len(nums)
    for i := 0; i < l - 1; i++ {
        for j := i + 1; j < l; j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }

    return nil
}
