package search

import (
    "log"
    "sort"
)

type DocResult struct {
    Score     float64
    WordCount uint64
}

type SearchContent struct {
    DocId uint64
    Score float64
    Url   string
    Text  string
}

func SearchText(text string, engine *Search) []*ProcessResult {
    words := engine.Tokenize.PartWord(text)
    allScoreMap := make(map[uint64]*DocResult)
    for i := range words {
        scoreMap, err := engine.SearchKey(words[i])
        if err != nil {
            log.Println("when search word ", words[i], "error")
            break
        }
        for key := range scoreMap {
            if _, ok := allScoreMap[key]; ok {
                allScoreMap[key].WordCount++
                allScoreMap[key].Score = allScoreMap[key].Score + scoreMap[key]
            } else {
                allScoreMap[key] = &DocResult{
                    Score:     scoreMap[key],
                    WordCount: 1,
                }
            }
        }
        for key := range allScoreMap {
            allScoreMap[key].Score = allScoreMap[key].Score * float64(allScoreMap[key].WordCount) / float64(len(words))
        }
    }
    processResults := make([]*ProcessResult, len(allScoreMap))
    index := 0
    for key := range allScoreMap {
        processResults[index] = &ProcessResult{
            Score: allScoreMap[key].Score,
            DocId: key,
        }
        index++
    }
    // sort by the score desc
    sort.Sort(ScoreSlice(processResults))
    return processResults
}

func SearchResult(text string, engine *Search) []*SearchContent {
    proRes := SearchText(text, engine)
    res := make([]*SearchContent, len(proRes))
    for i, result := range proRes {
        docRes, err := engine.docDB.FindFromDocDB(result.DocId)
        if err != nil {
            log.Println("Search error , when search the id from the docDB", err)
            break
        }
        res[i] = &SearchContent{
            DocId: result.DocId,
            Score: result.Score,
            Url:   docRes.Url,
            Text:  docRes.Text,
        }
    }
    return res
}
