package search

type ProcessResult struct {
    DocId uint64
    Score float64
}

type ScoreSlice []ProcessResult

func (x ScoreSlice) Len() int {
    return len(x)
}
func (x ScoreSlice) Less(i, j int) bool {
    return x[i].Score < x[j].Score
}
func (x ScoreSlice) Swap(i, j int) {
    x[i], x[j] = x[j], x[i]
}
