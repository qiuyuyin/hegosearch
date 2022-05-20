package data

import (
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
    "hegosearch/util"
)

type IndexDB struct {
    IndexDB *leveldb.DB
}

// when dbkv init , we open the database
func IndexDataInit(dbpath string) *IndexDB {
    kdb, err := leveldb.OpenFile(dbpath, &opt.Options{})
    index_db := IndexDB{}
    if err != nil {
        panic(err)
    }
    index_db.IndexDB = kdb
    return &index_db
}

func (indexDB *IndexDB) InsertIntoDocDB(key string, id uint64) error {
    ids := []uint64{id}
    encoder := util.Encoder(ids)
    err := indexDB.IndexDB.Put([]byte(key), encoder, nil)
    if err != nil {
        return err
    }
    return nil
}

// put the id into the word kvdb
func (indexDB *IndexDB) InsertIndexIntoWord(id uint64, key string) error {
    if len(key) == 0 {
        return nil
    }
    value, err := indexDB.IndexDB.Get([]byte(key), nil)
    if err != nil {
        if err == leveldb.ErrNotFound {
            return indexDB.InsertIntoDocDB(key, id)
        } else {
            return err
        }
    }
    ids := make([]uint64, 0)
    util.Decoder(value, &ids)
    ids = append(ids, id)
    encoder := util.Encoder(ids)
    err = indexDB.IndexDB.Put([]byte(key), encoder, nil)
    if err != nil {
        return err
    }
    return nil
}

func (indexDB *IndexDB) InsertIdsIntoIndexDB(ids []uint64, key string) error {
    encoder := util.Encoder(ids)
    err := indexDB.IndexDB.Put([]byte(key), encoder, nil)
    if err != nil {
        return err
    }
    return nil
}

func (indexDB *IndexDB) FindFromIndexDB(key string) ([]uint64, error) {
    value, err := indexDB.IndexDB.Get([]byte(key), nil)
    if err != nil {
        return nil, err
    }
    ids := make([]uint64, 0)
    util.Decoder(value, &ids)
    return ids, nil
}

func (indexDB *IndexDB) CloseIndexDB() {
    indexDB.IndexDB.Close()
}
