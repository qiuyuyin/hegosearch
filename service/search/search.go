package search

import (
    "hegosearch/data/doc"
    "hegosearch/data/index"
    "hegosearch/data/tokenize"
    "log"
    "math"
)

type Search struct {
    indexDB  *index.IndexDriver
    docDB    *doc.DocDriver
    Tokenize tokenize.Token
}

type DocItem struct {
    Id    uint64
    Count uint64
}

func NewSearch(indexDB *index.IndexDriver, docDB *doc.DocDriver, token tokenize.Token) *Search {
    search := Search{
        indexDB:  indexDB,
        docDB:    docDB,
        Tokenize: token,
    }
    return &search
}

func (search *Search) SearchKey(key string) (map[uint64]float64, error) {
    ids, err := search.indexDB.FindFromIndexDB(key)
    if err != nil {
        return nil, err
    }
    docMap := search.processIds(ids)
    // tf(t in d) = √frequency 词频
    // idf(t) = 1 + log ( numDocs / (docFreq + 1)) 逆向文档频率
    // norm(d) = 1 / √numTerms 字段长度归一值
    // 开始计算分数
    score := search.processScores(docMap)
    return score, nil
}

func (search *Search) processScores(docMap map[uint64]uint64) map[uint64]float64 {
    count := search.docDB.CountDoc()
    idf := math.Log10(float64(count)/float64(len(docMap))) + 1
    scoreMap := make(map[uint64]float64)
    for index, item := range docMap {
        tf := math.Sqrt(float64(item))
        doc, err := search.docDB.FindFromDocDB(index)
        if err != nil {
            log.Printf("Process Score Error %s", err)
            continue
        }
        textLen := len(doc.Text)
        norm := math.Pow(math.Sqrt(float64(textLen)), -1)
        keyScore := norm * tf * math.Pow(idf, 2)
        scoreMap[index] = keyScore
    }
    return scoreMap
}

func (search *Search) processIds(ids []uint64) map[uint64]uint64 {
    docMap := make(map[uint64]uint64)
    for i := range ids {
        docMap[ids[i]]++
    }
    return docMap
}
