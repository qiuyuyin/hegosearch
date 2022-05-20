package storage

// Storage To reduce coupling, multiple database operations are supported
// We use the Storage interface , if you want to use ,please New DB and put into this interface
type Storage interface {
    Get(key []byte) (value []byte, err error)
    Set(key []byte, value []byte) error
    Count() uint64
    Close()
}
