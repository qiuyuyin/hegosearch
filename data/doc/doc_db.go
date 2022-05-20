package doc

import (
    "encoding/binary"
    "github.com/cockroachdb/pebble"
    "hegosearch/data/model"
    "hegosearch/data/storage"
    "hegosearch/util"
    "sync/atomic"
)

const DocId = "doc_index"

type DocDriver struct {
    Storage  storage.Storage
    CurIndex uint64
}

// DocDataInit when dbkv init , we open the database
func NewDocDriver(store storage.Storage) *DocDriver {
    doc := DocDriver{Storage: store}
    value, err := store.Get([]byte(DocId))
    if err != nil {
        if err == pebble.ErrNotFound {
            doc.CurIndex = 0
        }
    } else {
        docId := binary.BigEndian.Uint64(value)
        doc.CurIndex = docId
    }
    return &doc
}

// InsertIntoDocDB This function is to put the doc into the db
func (docDriver *DocDriver) InsertIntoDocDB(value *model.Document) (uint64, error) {
    valueBytes := util.Encoder(value)
    id := docDriver.CurIndex
    keyBytes := util.IntToByte(id)
    docDriver.IncDocIndex()
    err := docDriver.Storage.Set(keyBytes, valueBytes)
    if err != nil {
        return id, err
    }
    return id, nil
}

// FindFromDocDB Through the key to find the doc
func (docDriver *DocDriver) FindFromDocDB(key uint64) (*model.Document, error) {
    doc := new(model.Document)
    keyBytes := util.IntToByte(key)
    value, err := docDriver.Storage.Get(keyBytes)
    if err != nil {
        return nil, err
    }
    util.Decoder(value, &doc)
    return doc, nil
}

// IncDocIndex through the key to find the doc
func (docDriver *DocDriver) IncDocIndex() {
    atomic.AddUint64(&docDriver.CurIndex, 1)
}

func (docDriver *DocDriver) CountDoc() uint64 {
    return docDriver.Storage.Count()
}

func (docDriver *DocDriver) CloseDocDB() {
    docDriver.Storage.Close()
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
