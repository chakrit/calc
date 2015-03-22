package main

import "testing"
import "strings"
import "strconv"
import "bytes"
import a "github.com/stretchr/testify/assert"

func TestRun(t *testing.T) {
	cases := map[string]int{
		"1+2+3+4+5+6+7+8+9+10": 55,
		"1+2*3":                7,
		"1+2*3+4+5":            16,
		"1+2*(3+4)+5":          20,
	}

	for expr, expected := range cases {
		input, output := strings.NewReader(expr), &bytes.Buffer{}
		run(input, output)

		stdout := strings.Trim(string(output.Bytes()), "\r\n")
		actual, e := strconv.Atoi(stdout)
		a.NoError(t, e)
		a.Equal(t, expected, actual)
	}
}
