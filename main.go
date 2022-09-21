package main

import (
    "fmt"
    "log"
    "example.com/password"
)

func main() {
    // getting pwd
    pwd := password.Get_pwd("Meth")
    fmt.Println(pwd)
    // adding pwd
    err := password.Add_pwd("Test", "pa55word")
    if err != nil {
        log.Fatal(err)
    }
    // testing for our pwd
    v := password.Get_pwd("Test")
    fmt.Printf("%s\n", v)
}
