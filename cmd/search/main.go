package main

import (
    "flag"
    "fmt"
    "hegosearch/data/doc"
    "hegosearch/data/index"
    "hegosearch/data/storage"
    "hegosearch/data/tokenize"
    "hegosearch/server"
    "hegosearch/service/search"
    "net/http"
    "os"
    "time"
)

func main() {
    path := flag.String("p", "data", "this is the path of data")
    flag.Parse()
    _, err := os.Stat(*path)
    if os.IsNotExist(err) {
        fmt.Println("path error")
        return
    }
    _, err = os.Stat(*path + "/db/index")
    if os.IsNotExist(err) {
        fmt.Println("path error")
        return
    }
    _, err = os.Stat(*path + "/db/doc")

    if os.IsNotExist(err) {
        fmt.Println("path error")
        return
    }
    docDB := doc.NewDocDriver(
        storage.NewLevelDBStorage(*path + "/db/doc"),
    )
    indexDB := index.NewIndexDriver(
        storage.NewLevelDBStorage(*path + "/db/index"),
    )
    token := tokenize.NewToken()
    newSearch := search.NewSearch(indexDB, docDB, token)
    router := server.Router(server.NewSearchSever(newSearch))
    s := &http.Server{
        Addr:           "127.0.0.1:8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
