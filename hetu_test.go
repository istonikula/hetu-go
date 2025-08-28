package hetu_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/istonikula/hetu-go"
	"github.com/istonikula/hetu-go/bday"
	"github.com/istonikula/hetu-go/nnn"
)

func TestHetu(t *testing.T) {
	var bday1900 = time.Date(1980, 2, 22, 0, 0, 0, 0, time.UTC)
	var bday2000 = time.Date(2000, 11, 12, 0, 0, 0, 0, time.UTC)
	var mustRandomBday = func() bday.Val {
		b, err := bday.Random(16, 107)
		require.NoError(t, err)
		return b
	}

	t.Run("century id", func(t *testing.T) {
		f := func(xs []string, format string) []string {
			result := make([]string, len(xs))
			for i, x := range xs {
				result[i] = fmt.Sprintf(format, x)
			}
			return result
		}

		tests := append(
			f([]string{"-", "Y", "X", "W", "V", "U"}, "180851%s229L"),
			f([]string{"A", "B", "C", "D", "E", "F"}, "010100%s030D")...,
		)

		for _, test := range tests {
			t.Run("parse: "+test, func(t *testing.T) {
				v, msg := hetu.Parse(test)
				assert.NoError(t, msg)
				assert.Equal(t, test, v.Str)
			})
		}
	})

	t.Run("valid", func(t *testing.T) {
		tests := []struct {
			input    string
			female   bool
			temporal bool
		}{
			{input: "210721-1703", female: true},
			{input: hetu.Generate(nnn.Generate(nnn.GenderFemale), bday.Val(bday1900)).Str, female: true},
			{input: hetu.Generate(nnn.Generate(nnn.GenderFemale), mustRandomBday()).Str, female: true},

			{input: "260785-309A"},
			{input: hetu.Generate(nnn.Generate(nnn.GenderMale), bday.Val(bday2000)).Str},
			{input: hetu.Generate(nnn.Generate(nnn.GenderMale), mustRandomBday()).Str},

			{input: "090707-998E", female: true, temporal: true},
			{input: hetu.Generate(nnn.GenerateTemporal(nnn.GenderFemale), bday.Val(bday1900)).Str, female: true, temporal: true},
			{input: hetu.Generate(nnn.GenerateTemporal(nnn.GenderFemale), mustRandomBday()).Str, female: true, temporal: true},

			{input: "050162-981N", temporal: true},
			{input: hetu.Generate(nnn.GenerateTemporal(nnn.GenderMale), bday.Val(bday1900)).Str, temporal: true},
			{input: hetu.Generate(nnn.GenerateTemporal(nnn.GenderMale), mustRandomBday()).Str, temporal: true},
		}

		for _, test := range tests {
			t.Run(test.input, func(t *testing.T) {
				v, err := hetu.Parse(test.input)
				assert.NoError(t, err)
				assert.Equal(t, test.female, v.Nnn.IsFemale())
				assert.Equal(t, !test.female, v.Nnn.IsMale())
				assert.Equal(t, test.temporal, v.Nnn.IsTemporal())
				assert.Equal(t, !test.temporal, v.Nnn.IsReal())
			})
		}
	})

	t.Run("invalid", func(t *testing.T) {
		tests := []struct {
			name  string
			input string
			msg   string
		}{
			{name: "birthday, format", input: "xxxxxx-450F", msg: "invalid birthday format"},
			{name: "birthday, date", input: "310425A450F", msg: "invalid birthday: 310425"},
			{name: "century id, lowercase", input: "220305a244S", msg: "unsupported century id: a"},
			{name: "century id, unknown", input: "220305G244S", msg: "unsupported century id: G"},
			{name: "nnn", input: "020146-000F", msg: "invalid nnn: 000"},
			{name: "control, lowercase", input: "020146-450f", msg: "invalid hetu: control char mismatch"},
			{name: "control, mismatch", input: "020146-450A", msg: "invalid hetu: control char mismatch"},
			{name: "too short", input: "020146-450", msg: "invalid hetu format"},
			{name: "too long", input: "020146-450FX", msg: "invalid hetu format"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, msg := hetu.Parse(test.input)
				assert.EqualError(t, msg, test.msg)
			})
		}
	})
}
