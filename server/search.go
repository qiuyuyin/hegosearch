package server

import "hegosearch/service/search"
import "github.com/gin-gonic/gin"

type SearchReq struct {
    Text     string `json:"text"`
    StopWord string `json:"stopWord"`
    Limit    int64  `json:"limit"`
}

type SearchResp struct {
    Result []*ResultDoc
}

type ResultDoc struct {
    DocId uint64  `json:"docId"`
    Score float64 `json:"score"`
    Url   string  `json:"url"`
    Text  string  `json:"text"`
}

type SearchSever struct {
    Engine *search.Search
}

func NewSearchSever(engine *search.Search) *SearchSever {
    return &SearchSever{Engine: engine}
}

func (ss *SearchSever) Search(c *gin.Context) {
    var searchReq SearchReq
    if err := c.ShouldBindJSON(&searchReq); err == nil {
        if err != nil {
            FailWithMessage("解析输入错误", c)
            return
        }
        result := search.SearchResult(searchReq.Text, ss.Engine)
        res := make([]*ResultDoc, len(result))
        for i, content := range result {
            res[i] = &ResultDoc{
                DocId: content.DocId,
                Score: content.Score,
                Url:   content.Url,
                Text:  content.Text,
            }
        }
        OkWithDetailed(SearchResp{Result: res}, "success", c)
    } else {
        FailWithMessage("参数错误", c)
    }
}
