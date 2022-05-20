package index

import (
    "github.com/syndtr/goleveldb/leveldb"
    "hegosearch/data/storage"
    "hegosearch/util"
)

type IndexDriver struct {
    IndexStorage storage.Storage
}

// when dbkv init , we open the database
func NewIndexDriver(store storage.Storage) *IndexDriver {
    return &IndexDriver{IndexStorage: store}
}

func (index *IndexDriver) InsertIntoDocDB(key string, id uint64) error {
    ids := []uint64{id}
    encoder := util.Encoder(ids)
    err := index.IndexStorage.Set([]byte(key), encoder)
    if err != nil {
        return err
    }
    return nil
}

// put the id into the word kvdb
func (index *IndexDriver) InsertIndexIntoWord(id uint64, key string) error {
    if len(key) == 0 {
        return nil
    }
    value, err := index.IndexStorage.Get([]byte(key))
    if err != nil {
        if err == leveldb.ErrNotFound {
            return index.InsertIntoDocDB(key, id)
        } else {
            return err
        }
    }
    ids := make([]uint64, 0)
    util.Decoder(value, &ids)
    ids = append(ids, id)
    encoder := util.Encoder(ids)
    err = index.IndexStorage.Set([]byte(key), encoder)
    if err != nil {
        return err
    }
    return nil
}

func (index *IndexDriver) InsertIdsIntoIndexDB(ids []uint64, key string) error {
    encoder := util.Encoder(ids)
    err := index.IndexStorage.Set([]byte(key), encoder)
    if err != nil {
        return err
    }
    return nil
}

func (index *IndexDriver) FindFromIndexDB(key string) ([]uint64, error) {
    value, err := index.IndexStorage.Get([]byte(key))
    if err != nil {
        return nil, err
    }
    ids := make([]uint64, 0)
    util.Decoder(value, &ids)
    return ids, nil
}

func (index *IndexDriver) CloseIndexDB() {
    index.IndexStorage.Close()
}
