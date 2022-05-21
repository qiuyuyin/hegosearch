package main

import (
    "hegosearch/data/doc"
    "hegosearch/data/index"
    "hegosearch/data/storage"
    "hegosearch/data/tokenize"
    "hegosearch/server"
    "hegosearch/service/search"
    "net/http"
    "time"
)

func main() {
    docDB := doc.NewDocDriver(
        storage.NewLevelDBStorage("data/db/doc"),
    )
    indexDB := index.NewIndexDriver(
        storage.NewLevelDBStorage("data/db/index"),
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
