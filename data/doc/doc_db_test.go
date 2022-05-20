package doc_test

import (
    "fmt"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "hegosearch/data/doc"
    "hegosearch/data/model"
    "hegosearch/data/storage"
    "log"
    "os"
)

var _ = Describe("DocDb", func() {
    var doc_db *doc.DocDriver
    document := &model.Document{
        Url:  "https://baidu.com",
        Text: "这是一个简单的操作",
    }
    BeforeEach(func() {
        levelDBStorage := storage.NewLevelDBStorage("tmp")
        doc_db = doc.NewDocDriver(levelDBStorage)
    })
    Describe("doc_db", func() {
        Context("test insert", func() {
            It("insert", func() {
                id, err := doc_db.InsertIntoDocDB(document)
                fmt.Printf("insert success, get the id: %d\n", id)
                if err != nil {
                    log.Fatalf("insert error")
                }
            })
        })
        Context("test find", func() {
            It("find", func() {
                id, err := doc_db.InsertIntoDocDB(document)
                if err != nil {
                    log.Fatalf("insert error")
                }
                doc, err := doc_db.FindFromDocDB(id)
                Expect(doc.Url).To(Equal("https://baidu.com"))
                Expect(doc.Text).To(Equal("这是一个简单的操作"))
                if err != nil {
                    log.Fatalf("find error")
                }
                fmt.Printf("find success : %v\n", doc)
            })
        })
    })
    AfterEach(func() {
        doc_db.CloseDocDB()
        err := os.RemoveAll("tmp")
        if err != nil {
            fmt.Println("delete error")
        }
    })
})
