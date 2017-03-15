package main

import "fmt"

func main() {
    var a string = "initial"
    fmt.Println(a) // "initial"

    var b, c int = 1, 2
    fmt.Println(b, c) // 1 2

    var d = true
    fmt.Println(d) // true

    var e int
    fmt.Println(e) // 0

    var f string
    fmt.Println(f) //

    var g bool
    fmt.Println(g) // false

    h := "short"
    fmt.Println(h) // short
}