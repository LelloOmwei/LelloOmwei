package semantics
import "time"
type Token struct {
    ID   string  `json:"id"`
    Sign string  `json:"sig"`
    Val  float64 `json:"val"`
    Ts   int64   `json:"ts"`
}
func NewToken(id, sign string, val float64) Token {
    return Token{ID: id, Sign: sign, Val: val, Ts: time.Now().Unix()}
}
