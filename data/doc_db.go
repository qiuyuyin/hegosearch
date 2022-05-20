package data

import (
    "encoding/binary"
    "github.com/cockroachdb/pebble"
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
    "hegosearch/data/model"
    "hegosearch/util"
    "sync/atomic"
)

const DocId = "doc_index"

type DocDB struct {
    DocDB    *leveldb.DB
    CurIndex uint64
}

// when dbkv init , we open the database
func DocDataInit(dbpath string) *DocDB {
    kdb, err := leveldb.OpenFile(dbpath, &opt.Options{})
    docDb := DocDB{}
    if err != nil {
        panic(err)
    }
    docDb.DocDB = kdb
    value, err := docDb.DocDB.Get([]byte(DocId), nil)
    if err != nil {
        if err == pebble.ErrNotFound {
            docDb.CurIndex = 0
        }
    } else {
        docId := binary.BigEndian.Uint64(value)
        docDb.CurIndex = docId
    }
    return &docDb
}

// this function is to put the doc into the db
func (docDb *DocDB) InsertIntoDocDB(value *model.Document) (uint64, error) {
    valueBytes := util.Encoder(value)
    id := docDb.CurIndex
    keyBytes := util.IntToByte(id)
    docDb.IncDocIndex()
    err := docDb.DocDB.Put(keyBytes, valueBytes, nil)
    if err != nil {
        return id, err
    }
    return id, nil
}

// through the key to find the doc
func (docDb *DocDB) FindFromDocDB(key uint64) (*model.Document, error) {
    doc := new(model.Document)
    keyBytes := util.IntToByte(key)
    value, err := docDb.DocDB.Get(keyBytes, nil)
    if err != nil {
        return nil, err
    }
    util.Decoder(value, &doc)
    return doc, nil
}

// through the key to find the doc
func (docDb *DocDB) IncDocIndex() {
    atomic.AddUint64(&docDb.CurIndex, 1)
}

func (docDb *DocDB) CountDocDB() uint64 {
    iter := docDb.DocDB.NewIterator(nil, nil)
    count := 0
    for iter.Next() {
        count++
    }
    return uint64(count)
}

func (docDb *DocDB) CloseDocDB() {
    docDb.DocDB.Close()
}

// put the id into the word kvdb
//func InsertIndexIntoWord(id uint64, key string) error {
//    value, closer, err := db.Get([]byte(key))
//    if err != nil {
//        return err
//    }
//    ids := make([]uint64, 0)
//    util.Decoder(value, &ids)
//    closer.Close()
//    ids = append(ids, id)
//    encoder := util.Encoder(ids)
//    err = db.Set([]byte(key), encoder, nil)
//    if err != nil {
//        return err
//    }
//    return nil
//}
