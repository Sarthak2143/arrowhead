package password

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
    "time"
)

func Get_pwd(email string) string {
    db := open_db() // calling open_db to return an instance of db
    create_bucket()
    var v []uint8
    db.View(func (tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Arrowhead"))
        v = b.Get([]byte(email))
        return nil
    })
    db.Close()
    return fmt.Sprintf("%s", v)
}

func open_db() *bolt.DB {    
    // opening db
    db, err := bolt.Open("password/pwd.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func Put_pwd(email, password string) error {
    db := open_db()
    create_bucket()
    // updating db
    var err error
    db.Update(func (tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Arrowhead"))
        err = b.Put([]byte(email), []byte(password))
        return err
    })
    db.Close()
    return err
}

func create_bucket() {
    db := open_db()
    //creating the db
    db.Update(func (tx *bolt.Tx) error {
        _, err := tx.CreateBucket([]byte("Arrowhead"))
        if err != nil {
            return fmt.Errorf("create bucket %s",err)
        }
        return nil
    })
    db.Close()
    return
}
