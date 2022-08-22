package main

import (
    "fmt"
    //"log"
    "example.com/password"
)

func main() {
    // TODO: Implement Arrowhead.
    pwd := password.Get_pwd("Meth")
    fmt.Println(pwd)
}
