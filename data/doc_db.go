package data

import (
    "encoding/binary"
    "github.com/cockroachdb/pebble"
    "github.com/cockroachdb/pebble/vfs"
    "hegosearch/data/model"
    "hegosearch/util"
    "reflect"
    "sync/atomic"
)

const DOC_ID = "doc_index"

type DocDB struct {
    DocDB    *pebble.DB
    CurIndex uint64
}

// when dbkv init , we open the database
func DocDataInit(dbpath string) *DocDB {
    kdb, err := pebble.Open(dbpath, &pebble.Options{FS: vfs.NewMem()})
    doc_db := DocDB{}
    if err != nil {
        panic(err)
    }
    doc_db.DocDB = kdb
    value, closer, err := doc_db.DocDB.Get([]byte(DOC_ID))
    if err != nil {
        if err == pebble.ErrNotFound {
            doc_db.CurIndex = 0
        }
    } else {
        doc_id := binary.BigEndian.Uint64(value)
        doc_db.CurIndex = doc_id
        closer.Close()
    }
    return &doc_db
}

// this function is to put the doc into the db
func (docDb *DocDB) InsertIntoDocDB(value *model.Document) (uint64, error) {
    valueBytes := util.Encoder(value)
    id := docDb.CurIndex
    keyBytes := util.IntToByte(id)
    docDb.IncDocIndex()
    err := docDb.DocDB.Set(keyBytes, valueBytes, nil)
    if err != nil {
        return id, err
    }
    return id, nil
}

// through the key to find the doc
func (docDb *DocDB) FindFromDocDB(key uint64) (*model.Document, error) {
    doc := new(model.Document)
    keyBytes := util.IntToByte(key)
    value, closer, err := docDb.DocDB.Get(keyBytes)
    if err != nil {
        return nil, err
    }
    util.Decoder(value, &doc)
    closer.Close()
    return doc, nil
}

// through the key to find the doc
func (docDb *DocDB) IncDocIndex() {
    atomic.AddUint64(&docDb.CurIndex, 1)
    docDb.DocDB.NewIndexedBatch()
}

func (docDb *DocDB) CountDocDB() uint64 {
    iter := docDb.DocDB.NewIter(nil)
    iterValue := reflect.ValueOf(*iter)
    num := iterValue.FieldByName("seqNum").Uint()
    return num
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
