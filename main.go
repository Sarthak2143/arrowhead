package main

import (
    "fmt"
    "log"
    "example.com/password"
    "github.com/boltdb/bolt"
    //"time"
)

func main() {
    fmt.Println(password.Get_pwd())
    // opening db
    db, err := bolt.Open("data.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    // creating the db
    db.Update(func (tx *bolt.Tx) error {
        _, err := tx.CreateBucket([]byte("Passwords"))
        if err != nil {
            return fmt.Errorf("create bucket %s",err)
        }
        return nil
    })
    // updating db
    db.Update(func (tx *bolt.Tx) error {
        b := tx.Bucket([]byte("MyBucket"))
        err := b.Put([]byte("Username"), []byte("Password"))
        return err
    })
    // getting value from db
    db.View(func (tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Passwords"))
        v := b.Get([]byte("Username"))
        fmt.Printf("%s", v)
        return nil
    })
    // closing db
    defer db.Close()
}
