package operation

import (
	"math/rand"
	"time"

	db "github.com/SG143s/DGRPL/api/product/database"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Genid() string {
	b := make([]byte, 8)
	var idexist = false

	for !idexist {
		for i := range b {
			b[i] = letters[r.Intn(len(letters))]
		}
		idexist = db.ChId(string(b))
	}
	return string(b)
}
