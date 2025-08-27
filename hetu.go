package hetu

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/istonikula/hetu-go/bday"
	"github.com/istonikula/hetu-go/century"
	"github.com/istonikula/hetu-go/ctrl"
	"github.com/istonikula/hetu-go/nnn"
)

type Valid struct {
	Val string
}

func New(b bday.Val, c century.Val, n nnn.Val, cc rune) Valid {
	// YYYYMMDD (8) + Century (1) + NNN (3) + CC (1) = 13 bytes.
	var sb strings.Builder
	sb.Grow(13)

	sb.WriteString(strconv.Itoa(b.Number()))
	sb.WriteRune(rune(c.Id))
	fmt.Fprintf(&sb, "%03d", n)
	sb.WriteRune(cc)

	return Valid{Val: sb.String()}
}

const (
	groupBday int = iota + 1
	groupCentury
	groupNnn
	groupCtrl
)

var pattern = regexp.MustCompile(`^(.{6})(.)(.{3})(.)`)

func Parse(candidate string) (Valid, error) {
	m := pattern.FindStringSubmatch(candidate)

	if m == nil {
		return Valid{}, errors.New("invalid hetu: pattern mismatch")
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

	return New(b, c, n, cc), nil
}
