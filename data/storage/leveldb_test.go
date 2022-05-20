package storage

import (
    "bytes"
    "os"
    "testing"
)

func TestLevelDBStorage_Get(t *testing.T) {
    levelDB := NewLevelDBStorage("temp")
    defer os.RemoveAll("temp")
    var store Storage = levelDB
    defer store.Close()
    err := store.Set([]byte("hello"), []byte("world"))
    if err != nil {
        t.Errorf("err")
    }
    value, err := store.Get([]byte("hello"))
    if err != nil {
        t.Errorf("err")
    }
    if !bytes.Equal(value, []byte("world")) {
        t.Errorf("not equal")
    }

}
