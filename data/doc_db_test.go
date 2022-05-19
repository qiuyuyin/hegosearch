package data_test

import (
    "fmt"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "hegosearch/data"
    "hegosearch/data/model"
    "log"
    "os"
)

var _ = Describe("DocDb", func() {
    var doc_db *data.DocDB
    doc := &model.Document{
        Url:  "http://baidu.com",
        Text: "这是一个简单的操作",
    }
    BeforeEach(func() {
        doc_db = data.DocDataInit("tmp")
        fmt.Println("err")
    })
    Describe("doc_db", func() {
        Context("test insert", func() {
            It("insert", func() {
                id, err := doc_db.InsertIntoDocDB(doc)
                fmt.Printf("insert success, get the id: %d\n", id)
                if err != nil {
                    log.Fatalf("insert error")
                }
            })
        })
        Context("test find", func() {
            It("find", func() {
                id, err := doc_db.InsertIntoDocDB(doc)
                if err != nil {
                    log.Fatalf("insert error")
                }
                doc, err := doc_db.FindFromDocDB(id)
                Expect(doc.Url).To(Equal("http://baidu.com"))
                Expect(doc.Text).To(Equal("这是一个简单的操作"))
                if err != nil {
                    log.Fatalf("find error")
                }
                fmt.Printf("find success : %v\n", doc)
            })
        })
    })
    AfterEach(func() {
        doc_db.DocDB.Close()
        err := os.RemoveAll("tmp")
        if err != nil {
            fmt.Println("delete error")
        }
        fmt.Println("after")
    })
})
