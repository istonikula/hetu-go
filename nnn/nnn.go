package nnn

import (
	"fmt"
	"math/rand"
	"strconv"
)

// spec: three digits between 002-899. 900-999 used for temporal ssns
// female: even
// male: odd
type Val int

type Gender int

const (
	GenderFemale Gender = 0
	GenderMale   Gender = 1
)

func (v Val) IsFemale() bool {
	return isFemale(int(v))
}

func (v Val) IsMale() bool {
	return !v.IsFemale()
}

func (v Val) IsReal() bool {
	return 2 <= v && v <= 899
}

func (v Val) IsTemporal() bool {
	return 900 <= v && v <= 999
}

func Parse(candidate string) (Val, error) {
	parsed, err := strconv.Atoi(candidate)
	if err != nil {
		return 0, fmt.Errorf("invalid nnn: %s", candidate)
	}

	v := Val(parsed)
	if v.IsReal() || v.IsTemporal() {
		return v, nil
	}

	return 0, fmt.Errorf("invalid nnn: %s", candidate)
}

var Generate = g(2, 899)
var GenerateTemporal = g(900, 999)

func g(min, max int) func(gender Gender) Val {
	return func(gender Gender) Val {
		for {
			candidate := rand.Intn(max-min+1) + min
			if gender == GenderFemale {
				if isFemale(candidate) {
					return Val(candidate)
				}
				continue
			}
			if gender == GenderMale {
				if !isFemale(candidate) {
					return Val(candidate)
				}
				continue
			}
		}
	}
}

func isFemale(x int) bool {
	return x%2 == int(GenderFemale)
}
