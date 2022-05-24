package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "hegosearch/data/doc"
    "hegosearch/data/index"
    "hegosearch/data/model"
    "hegosearch/data/storage"
    "hegosearch/data/tokenize"
    "io"
    "log"
    "os"
    "time"
)

// import the csv data from dataset

func main() {
    path := flag.String("p", "data", "this is the path of data")
    filePath := flag.String("f", "data/dataset/wukong_100m_1.csv", "this is the path of data")
    token := tokenize.NewToken()
    csvfile, err := os.Open(*filePath)
    if err != nil {
        panic(err)
    }
    defer csvfile.Close()

    csvReader := csv.NewReader(csvfile)
    docDB := doc.NewDocDriver(
        storage.NewLevelDBStorage(*path + "/db/doc"),
    )
    indexDB := index.NewIndexDriver(
        storage.NewLevelDBStorage(*path + "/db/index"),
    )
    defer docDB.Storage.Close()
    defer indexDB.IndexStorage.Close()
    _, err = csvReader.Read()
    if err == io.EOF {
        log.Fatalf("read first error")
    }
    // start the time and count
    start := time.Now()
    count := 0
    // new map to contain the index result
    wordMap := make(map[string][]uint64)
    for {
        count++
        record, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        // fmt.Printf("image: %s,text: %s \n", record[0], record[1])
        if len(record[0]) == 0 || len(record[1]) == 0 {
            continue
        }
        doc := model.Document{
            Url:  record[0],
            Text: record[1],
        }
        id, err := docDB.InsertIntoDocDB(&doc)
        if err != nil {
            panic(err)
        }
        words := token.PartWord(doc.Text)

        for i := range words {
            if ids, ok := wordMap[words[i]]; ok {
                wordMap[words[i]] = append(ids, id)
            } else {
                wordMap[words[i]] = []uint64{id}
            }
        }
        // if count print the result
        if count%10000 == 0 {
            cost := time.Since(start)
            fmt.Println("import cost:", cost.Seconds(), "s", " || and the count: ", count)
            start = time.Now()
        }
    }
    for k, v := range wordMap {
        err := indexDB.InsertIdsIntoIndexDB(v, k)
        if err != nil {
            panic(err)
        }
    }
    fmt.Println("import count:", count)
}
