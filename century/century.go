package century

import (
	"fmt"
	"math/rand"
)

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

func Random(hint Id) Val {
	ids := IdsByNum[ById[hint].Num]
	id := ids[rand.Intn(len(ids))]
	return ById[id]
}

var ById map[Id]Val
var IdsByNum map[int][]Id

func init() {
	ById = make(map[Id]Val)

	ById['+'] = Val{Id: '+', Num: 1800}

	for _, id := range []Id{'-', 'Y', 'X', 'W', 'V', 'U'} {
		ById[id] = Val{Id: id, Num: 1900}
	}

	for _, id := range []Id{'A', 'B', 'C', 'D', 'E', 'F'} {
		ById[id] = Val{Id: id, Num: 2000}
	}

	IdsByNum = make(map[int][]Id)
	for key, val := range ById {
		IdsByNum[val.Num] = append(IdsByNum[val.Num], key)
	}
}
