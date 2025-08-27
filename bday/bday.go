package bday

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/istonikula/hetu-go/century"
)

type Val time.Time

func (v Val) Number() int {
	y, m, d := time.Time(v).Date()
	return d*10000 + int(m)*100 + y%100
}

func (v Val) Century() int {
	y, _, _ := time.Time(v).Date()
	return y / 100 * 100
}

var pattern = regexp.MustCompile(`^(\d{2})(\d{2})(\d{2})$`)

func Parse(candidate string, c century.Val) (Val, error) {
	match := pattern.FindStringSubmatch(candidate)
	if match == nil {
		return Val{}, errors.New("invalid bday: pattern mismatch")
	}

	d, _ := strconv.Atoi(match[1])
	m, _ := strconv.Atoi(match[2])
	y, _ := strconv.Atoi(match[3])

	date := time.Date(c.Num+y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	return Val(date), nil
}
