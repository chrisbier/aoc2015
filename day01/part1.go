package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("input1")
    if err != nil {
        fmt.Println("Error opening file!!!")
    }
    defer file.Close()

    // declare chunk size
    const maxSz = 1
    floor_level := 0

    // create buffer
    b := make([]byte, maxSz)

    for {
        // read content to buffer
        readTotal, err := file.Read(b)
        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }

        if string(b[:1]) == "(" {
            floor_level += 1
        } else if string(b[:1]) == ")" {
            floor_level -= 1
        }

        // fmt.Println(floor_level, string(b[:readTotal])) // print content from buffer
        fmt.Println(readTotal) // print content from buffer
    }
    fmt.Println(floor_level)
}
