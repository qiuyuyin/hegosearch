package model

type SearchReq struct {
    Text     string `json:"text"`
    StopWord string `json:"stopWord"`
    Limit    uint64 `json:"limit"`
}
