package data

import (
    "github.com/go-ego/gse"
    "github.com/go-ego/gse/hmm/pos"
)

var (
    seg    gse.Segmenter
    posSeg pos.Segmenter
)

func JiebaInit() {
    seg.LoadDict()
    seg.LoadStop("zh")
}

func PartWord(text string) []string {
    return partWordCut(text)
}
func partWordCut(text string) []string {
    res := seg.CutSearch(text, true)
    res = seg.TrimPunct(res)
    return res
}
