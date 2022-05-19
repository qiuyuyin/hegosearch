package util

import (
    "bytes"
    "encoding/binary"
    "encoding/gob"
    "log"
)

func IntToByte(num uint64) []byte {
    bytes := make([]byte, 8)
    binary.BigEndian.PutUint64(bytes, num)
    return bytes
}

func Decoder(data []byte, v interface{}) {
    if data == nil {
        return
    }
    buffer := bytes.NewBuffer(data)
    decoder := gob.NewDecoder(buffer)
    err := decoder.Decode(v)
    if err != nil {
        log.Printf("decode error:%s", err)
    }
    return
}

func Encoder(data interface{}) []byte {
    if data == nil {
        return nil
    }
    buffer := new(bytes.Buffer)
    encoder := gob.NewEncoder(buffer)
    err := encoder.Encode(data)
    if err != nil {
        return nil
    }
    return buffer.Bytes()
}
