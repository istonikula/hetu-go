package bday

import (
	"errors"
	"fmt"
	"math/rand"
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
		return Val{}, errors.New("invalid birthday format")
	}

	d, _ := strconv.Atoi(match[1])
	m, _ := strconv.Atoi(match[2])
	y, _ := strconv.Atoi(match[3])

	date, ok := validate(d, m, c.Num+y)
	if !ok {
		return Val{}, fmt.Errorf("invalid birthday: %s", candidate)
	}
	return Val(date), nil
}

func From(day, month, year int) (Val, error) {
	date, ok := validate(day, month, year)
	if !ok {
		return Val{}, fmt.Errorf("invalid date: %02d.%02d.%d", day, month, year)
	}
	return Val(date), nil
}

func Random(minAge, maxAge int) (Val, error) {
	if minAge >= maxAge {
		return Val{}, errors.New("minAge must be less than maxAge")
	}

	now := time.Now()

	y1, m1, d1 := now.AddDate(-maxAge, 0, 0).Date()
	min := time.Date(y1, m1, d1, 0, 0, 0, 0, time.UTC).UnixMilli()

	y2, m2, d2 := now.AddDate(-minAge, 0, 0).Date()
	max := time.Date(y2, m2, d2, 0, 0, 0, 0, time.UTC).UnixMilli()

	delta := rand.Int63n(max - min + 1) // +1 to make the range inclusive
	return Val(time.UnixMilli(min + delta)), nil
}

func validate(d, m, y int) (time.Time, bool) {
	date := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	ok := date.Day() == d && int(date.Month()) == m && date.Year() == y
	return date, ok
}
