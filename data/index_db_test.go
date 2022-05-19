package data_test

import (
    "fmt"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "log"
    "os"

    "hegosearch/data"
)

var _ = Describe("IndexDb", func() {
    var indexDB *data.IndexDB
    BeforeEach(func() {
        indexDB = data.IndexDataInit("tmp")
        fmt.Println("err")
    })
    Describe("doc_db", func() {
        Context("test insert", func() {
            It("insert", func() {
                if err := indexDB.InsertIndexIntoWord(0, "伊离"); err != nil {
                    log.Fatalf("insert error")
                }
                if err := indexDB.InsertIndexIntoWord(2, "伊离"); err != nil {
                    log.Fatalf("insert error")
                }
                ids, err := indexDB.FindFromIndexDB("伊离")
                if err != nil {
                    log.Fatalf("find error")
                }
                Expect(ids[0]).To(Equal(uint64(0)))
                Expect(ids[1]).To(Equal(uint64(2)))
            })
        })

    })
    AfterEach(func() {
        indexDB.IndexDB.Close()
        err := os.RemoveAll("tmp")
        if err != nil {
            fmt.Println("delete error")
        }
        fmt.Println("after")
    })
})
