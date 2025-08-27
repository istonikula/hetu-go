package nnn

import (
	"fmt"
	"strconv"
)

// spec: three digits between 002-899. 900-999 used for temporal ssns
// female: even
// male: odd
type Val int

func (v Val) IsFemale() bool {
	return v%2 == 0
}

func (v Val) IsMale() bool {
	return v%2 != 0
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
