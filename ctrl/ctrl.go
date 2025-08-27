package ctrl

import (
	"github.com/istonikula/hetu-go/bday"
	"github.com/istonikula/hetu-go/nnn"
)

var lookup = "0123456789ABCDEFHJKLMNPRSTUVWXY"

func From(b bday.Val, n nnn.Val) rune {
	num := b.Number()*1000 + int(n)
	idx := num % len(lookup)
	return rune(lookup[idx])
}
