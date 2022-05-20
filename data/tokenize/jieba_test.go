package tokenize

import (
    "fmt"
    "testing"
)

func TestName(t *testing.T) {
    JiebaInit()
    words := PartWord("《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的的科幻片.")
    fmt.Println(words)
}
