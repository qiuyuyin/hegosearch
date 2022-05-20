package tokenize

import (
    "github.com/go-ego/gse"
    "github.com/go-ego/gse/hmm/pos"
)

var ()

type Token struct {
    seg    *gse.Segmenter
    posSeg *pos.Segmenter
}

func NewToken() *Token {
    var seg gse.Segmenter
    var posSeg *pos.Segmenter
    seg.LoadDict()
    seg.LoadStop("zh")
    return &Token{
        seg:    &seg,
        posSeg: posSeg,
    }

}

func (token *Token) PartWord(text string) []string {
    return token.partWordCut(text)
}
func (token *Token) partWordCut(text string) []string {
    res := token.seg.CutSearch(text, true)
    res = token.seg.TrimPunct(res)
    return res
}
