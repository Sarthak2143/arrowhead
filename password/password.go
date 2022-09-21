package password

import (
    "log"
    "fmt"
    "time"
    bolt "go.etcd.io/bbolt"
)

func Get_pwd(email string) []uint8 {
    // opening db
    db := open_db()
    var v []uint8
    // viewing an element from our bucket
    err := db.View(func (tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Arrowhead"))
        v = b.Get([]byte(email))
        return nil
    })
    if err != nil {
        // checking for errors
        log.Fatal(err)
    }
    // always close db
    db.Close()
    // returning password
    return v
}

func Add_pwd(email, password string) error {
    // function to add passwords in our db
    db := open_db()
    err := db.Update(func (tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Arrowhead"))
        err := b.Put([]byte(email), []byte(password))
        return err
    })
    db.Close()
    return err
}

func open_db() *bolt.DB {
    // helper function to return an instance of db
    db, err := bolt.Open("password/password.db", 0600, &bolt.Options{
        Timeout: 1 * time.Second,
    })
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func create_bucket() error {
    // helper function to create a bucket
    db := open_db()
    err := db.Update(func(tx *bolt.Tx) error {
        b,err := tx.CreateBucket([]byte("Arrowhead"))
        fmt.Printf("%v\n", b)
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        return nil
    })
    return err
}
