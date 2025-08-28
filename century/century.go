package century

import "fmt"

type Id rune

type Val struct {
	Id  Id
	Num int
}

func Parse(candidate string) (Val, error) {
	c, ok := ById[Id(rune(candidate[0]))]
	if !ok {
		return c, fmt.Errorf("unsupported century id: %s", candidate)
	}
	return c, nil
}

var ById map[Id]Val

func init() {
	ById = make(map[Id]Val)

	ById['+'] = Val{Id: '+', Num: 1800}

	for _, id := range []Id{'-', 'Y', 'X', 'W', 'V', 'U'} {
		ById[id] = Val{Id: id, Num: 1900}
	}

	for _, id := range []Id{'A', 'B', 'C', 'D', 'E', 'F'} {
		ById[id] = Val{Id: id, Num: 2000}
	}
}
