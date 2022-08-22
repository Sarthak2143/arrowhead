package main

import (
    "fmt"
    "log"
    "example.com/password"
)

func main() {
    // TODO: Implement Arrowhead.
    pwd := password.Get_pwd("Meth")
    fmt.Println(pwd)
    err := password.Put_pwd("Test", "pa55word")
    if err != nil {
        log.Fatal(err)
    }

    v, _ := password.Get_pwd("Test")
    fmt.Printf("%s\n", v)
}
