package yacc_parseJson

import (
	"strings"
)

func ParseJson(d string, debug bool) (interface{}, error) {
	s := NewScanner(strings.NewReader(d), debug)
	yyParse(s)
	if s.err != nil {
		return nil, s.err
	}
	return s.data, nil
}
