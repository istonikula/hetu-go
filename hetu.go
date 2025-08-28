package hetu

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/istonikula/hetu-go/bday"
	"github.com/istonikula/hetu-go/century"
	"github.com/istonikula/hetu-go/ctrl"
	"github.com/istonikula/hetu-go/nnn"
)

type Valid struct {
	Birthday bday.Val
	Century  century.Val
	Nnn      nnn.Val
	Control  rune
	Str      string
}

const (
	groupBday int = iota + 1
	groupCentury
	groupNnn
	groupCtrl
)

var pattern = regexp.MustCompile(`^(.{6})(.)(.{3})(.)$`)

func Parse(candidate string) (Valid, error) {
	m := pattern.FindStringSubmatch(candidate)

	if m == nil {
		return Valid{}, errors.New("invalid hetu format")
	}

	c, err := century.Parse(m[groupCentury])
	if err != nil {
		return Valid{}, err
	}

	b, err := bday.Parse(m[groupBday], c)
	if err != nil {
		return Valid{}, err
	}

	n, err := nnn.Parse(m[groupNnn])
	if err != nil {
		return Valid{}, err
	}

	cc := rune(m[groupCtrl][0])
	if ctrl.From(b, n) != cc {
		return Valid{}, errors.New("invalid hetu: control char mismatch")
	}

	return Valid{Birthday: b, Century: c, Control: cc, Nnn: n, Str: candidate}, nil
}

// NOTE: generates "-" and "A" century ids only
func Generate(n nnn.Val, b bday.Val) Valid {
	cc := ctrl.From(b, n)
	var c century.Val
	if b.Century() == 1900 {
		c, _ = century.Parse("-")
	} else {
		c, _ = century.Parse("A")
	}

	// DDMMYY (6) + Century (1) + NNN (3) + CC (1) = 11 bytes.
	var sb strings.Builder
	sb.Grow(11)
	fmt.Fprintf(&sb, "%06d", b.Number())
	sb.WriteRune(rune(c.Id))
	fmt.Fprintf(&sb, "%03d", n)
	sb.WriteRune(cc)

	return Valid{Birthday: b, Century: c, Nnn: n, Control: cc, Str: sb.String()}
}
