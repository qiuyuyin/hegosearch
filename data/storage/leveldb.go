package storage

import (
    "github.com/syndtr/goleveldb/leveldb"
    "github.com/syndtr/goleveldb/leveldb/opt"
    "log"
    "reflect"
)

type LevelDBStorage struct {
    DB *leveldb.DB
}

func NewLevelDBStorage(path string) *LevelDBStorage {
    kdb, err := leveldb.OpenFile(path, &opt.Options{})
    if err != nil {
        log.Fatalf("open leveldb err:%s", err)
    }
    return &LevelDBStorage{DB: kdb}
}

func (s *LevelDBStorage) Get(key []byte) (value []byte, err error) {
    return s.DB.Get(key, nil)
}

func (s *LevelDBStorage) Set(key []byte, value []byte) error {
    return s.DB.Put(key, value, nil)
}

func (s *LevelDBStorage) Count() uint64 {
    iter := s.DB.NewIterator(nil, nil)
    iterPointer := reflect.ValueOf(iter)
    iterValue := reflect.ValueOf(iterPointer.Elem())
    return iterValue.FieldByName("seq").Uint()
}

func (s *LevelDBStorage) Close() {
    s.DB.Close()
}
