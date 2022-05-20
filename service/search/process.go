package search

import "log"

type DocResult struct {
    Score     float64
    WordCount uint64
}

func SearchText(text string, engine *Search) map[uint64]*DocResult {
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
                allScoreMap[key].WordCount = 1
                allScoreMap[key].Score = scoreMap[key]
            }
        }
        for key := range allScoreMap {
            allScoreMap[key].Score = allScoreMap[key].Score * float64(allScoreMap[key].WordCount) / float64(len(words))
        }
    }
    return allScoreMap
}
