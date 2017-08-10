package main

import (
    "fmt"
)

func main() {
    var s string
    s = "abcabcbb"
    fmt.Printf("len(%s) = %d\n", s, lengthOfLongestSubstring(s))
    s = "bbbbb"
    fmt.Printf("len(%s) = %d\n", s, lengthOfLongestSubstring(s))
    s = "pwwkew"
    fmt.Printf("len(%s) = %d\n", s, lengthOfLongestSubstring(s))
    s = "abba"
    fmt.Printf("len(%s) = %d\n", s, lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
    max := 0
    cs := make(map[byte]int)
    bs := []byte(s)
    l := len(s)
    p := -1
    for i := 0; i < l; i++ {
        t := 0
        if tp, ok := cs[bs[i]]; ok && tp > p {
            p = tp
        }
        t = i - p
        cs[bs[i]] = i
        if t > max {
            max = t
        }
    }

    return max
}
