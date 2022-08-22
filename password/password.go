package password

import (
    "log"
    "fmt"
    "time"
    bolt "go.etcd.io/bbolt"
)

func Get_pwd(email string) string {
    if err := create_bucket(); err != nil {
        log.Fatal(err)
    }
    return email + "test"
}

func create_db() *bolt.DB {
    db, err := bolt.Open("password/password.db", 0600, &bolt.Options{
        Timeout: 1 * time.Second,
    })
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func create_bucket() error {
    db := create_db()
    err := db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucket([]byte("Arowhead"))
        fmt.Printf("%v\n", b)
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        return nil
    })
    return err
}


