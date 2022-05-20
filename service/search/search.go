package search

import (
    "hegosearch/data"
    "log"
    "math"
)

type Search struct {
    indexDB  *data.IndexDB
    docDB    *data.DocDB
    DocItems []*DocItem
    DocMap   map[uint64]uint64
    ScoreMap map[uint64]float64
}

type DocItem struct {
    Id    uint64
    Count uint64
}

func SearchInit(indexDB *data.IndexDB, docDB *data.DocDB) *Search {
    search := Search{
        indexDB:  indexDB,
        docDB:    docDB,
        DocMap:   make(map[uint64]uint64),
        ScoreMap: make(map[uint64]float64),
    }
    return &search
}

func (search *Search) SearchKey(key string) (map[uint64]float64, error) {
    ids, err := search.indexDB.FindFromIndexDB(key)
    if err != nil {
        return nil, err
    }
    search.processIds(ids)
    // tf(t in d) = √frequency 词频
    // idf(t) = 1 + log ( numDocs / (docFreq + 1)) 逆向文档频率
    // norm(d) = 1 / √numTerms 字段长度归一值
    // 开始计算分数
    search.processScores()
    return search.ScoreMap, nil
}

func (search *Search) processScores() {
    count := search.docDB.CountDocDB()
    idf := math.Log10(float64(count) / float64(len(search.DocMap)))
    for index, item := range search.DocMap {
        tf := math.Sqrt(float64(item))
        doc, err := search.docDB.FindFromDocDB(index)
        if err != nil {
            log.Printf("Process Score Error %s", err)
            continue
        }
        textLen := len(doc.Text)
        norm := math.Pow(math.Sqrt(float64(textLen)), -1)
        keyScore := norm * tf * idf
        search.ScoreMap[index] = keyScore
    }
}

func (search *Search) processIds(ids []uint64) {
    for i := range ids {
        search.DocMap[ids[i]]++
    }
    items := make([]*DocItem, len(search.DocMap))
    for key, value := range search.DocMap {
        items = append(items, &DocItem{
            Id:    key,
            Count: value,
        })
    }
    search.DocItems = items
}
