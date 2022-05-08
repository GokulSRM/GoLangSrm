package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
)

func main() {
    flag.Parse()
    s := flag.Arg(0)

    i, err := strconv.Atoi(s)
    if err != nil {

        fmt.Println(err)
        os.Exit(2)
    }
    fmt.Println(s, i)
}