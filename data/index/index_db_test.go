package index_test

import (
    "fmt"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "hegosearch/data/index"
    "hegosearch/data/storage"
    "log"
    "os"
)

var _ = Describe("IndexDb", func() {
    var indexDB *index.IndexDriver
    BeforeEach(func() {
        levelDBStorage := storage.NewLevelDBStorage("tmp")
        indexDB = index.NewIndexDriver(levelDBStorage)
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
        indexDB.CloseIndexDB()
        err := os.RemoveAll("tmp")
        if err != nil {
            fmt.Println("delete error")
        }
    })
})
