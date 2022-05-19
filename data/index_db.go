package data

import (
    "github.com/cockroachdb/pebble"
    "hegosearch/util"
)

type IndexDB struct {
    IndexDB *pebble.DB
}

// when dbkv init , we open the database
func IndexDataInit(dbpath string) *IndexDB {
    kdb, err := pebble.Open(dbpath, &pebble.Options{})
    index_db := IndexDB{}
    if err != nil {
        panic(err)
    }
    index_db.IndexDB = kdb
    return &index_db
}

func (indexDb *IndexDB) InsertIntoDocDB(key string, id uint64) error {
    ids := []uint64{id}
    encoder := util.Encoder(ids)
    err := indexDb.IndexDB.Set([]byte(key), encoder, nil)
    if err != nil {
        return err
    }
    return nil
}

// put the id into the word kvdb
func (indexDb *IndexDB) InsertIndexIntoWord(id uint64, key string) error {
    if len(key) == 0 {
        return nil
    }
    value, closer, err := indexDb.IndexDB.Get([]byte(key))
    if err != nil {
        if err == pebble.ErrNotFound {
            return indexDb.InsertIntoDocDB(key, id)
        } else {
            return err
        }
    }
    ids := make([]uint64, 0)
    util.Decoder(value, &ids)
    closer.Close()
    ids = append(ids, id)
    encoder := util.Encoder(ids)
    err = indexDb.IndexDB.Set([]byte(key), encoder, nil)
    if err != nil {
        return err
    }
    return nil
}

func (indexDB *IndexDB) FindFromIndexDB(key string) ([]uint64, error) {
    value, closer, err := indexDB.IndexDB.Get([]byte(key))
    if err != nil {
        return nil, err
    }
    ids := make([]uint64, 0)
    util.Decoder(value, &ids)
    closer.Close()
    return ids, nil
}
