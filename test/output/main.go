package main

import "fmt"

func main() {
    for i:=0; i<5000; i++ {
        fmt.Println("TEXT TEXT TEXT TEXT TEXT", i)
        fmt.Println("TEXT TEXT TEXT TEXT TEXT", i)
        fmt.Println("TEXT TEXT TEXT TEXT TEXT", i)
        fmt.Println("TEXT TEXT TEXT TEXT TEXT", i)
    }
    fmt.Println("DONE")
}
