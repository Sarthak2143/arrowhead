package password

import (
    "log"
    "fmt"
    "time"
    bolt "go.etcd.io/bbolt"
)

func Get_pwd(email string) ([]uint8, error) {
    db := open_db()
    var v []uint8
    err := db.View(func (tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Arrowhead"))
        v = b.Get([]byte(email))
        return nil
    })
    db.Close()
    return v, err
}

func Put_pwd(email, password string) error {
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
    db, err := bolt.Open("password/password.db", 0600, &bolt.Options{
        Timeout: 1 * time.Second,
    })
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func create_bucket() error {
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


