package yacc_parseJson

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
)

type Scanner struct {
	buf   *bufio.Reader
	data  interface{}
	err   error
	debug bool
}

func NewScanner(r io.Reader, debug bool) *Scanner {
	return &Scanner{
		buf:   bufio.NewReader(r),
		debug: debug,
	}
}

func (sc *Scanner) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func (sc *Scanner) Reduced(rule, state int, lval *yySymType) bool {
	if sc.debug {
		fmt.Printf("rule: %v; state %v; lval: %v\n", rule, state, lval)
	}
	return false
}

func (s *Scanner) Lex(lval *yySymType) int {
	return s.lex(lval)
}

func (s *Scanner) lex(lval *yySymType) int {
	for {
		r := s.read()
		if r == 0 {
			return 0
		}
		if isWhitespace(r) {
			continue
		}

		if isDigit(r) {
			s.unread()
			lval.i = s.scanNumber()
			return NUMBER
		}

		switch r {
		case '[':
			return LS
		case ']':
			return RS
		case '{':
			return LC
		case '}':
			return RC
		case ',':
			return COMMA
		case ':':
			return COLON
		case '"':
			s.unread()
			lval.s = s.scanStr()
			return STRING
		case 't':
			s.unread()
			if s.scanTrue() {
				return TRUE
			}
		case 'f':
			s.unread()
			if s.scanFalse() {
				return FALSE
			}
		case 'n':
			s.unread()
			if s.scanNull() {
				return NULL
			}
		default:
			s.err = errors.New("error")
			return 0
		}
	}
}

func (s *Scanner) scanTrue() bool {
	t := []rune{'t', 'r', 'u', 'e'}
	for _, i := range t {
		r := s.read()
		if r != i {
			s.err = errors.New("true is error")
			return false
		}
	}
	return true
}

func (s *Scanner) scanFalse() bool {
	t := []rune{'f', 'a', 'l', 's', 'e'}
	for _, i := range t {
		r := s.read()
		if r != i {
			s.err = errors.New("false is error")
			return false
		}
	}
	return true
}

func (s *Scanner) scanNull() bool {
	t := []rune{'n', 'u', 'l', 'l'}
	for _, i := range t {
		r := s.read()
		if r != i {
			s.err = errors.New("null is error")
			return false
		}
	}
	return true
}

func (s *Scanner) scanStr() string {
	var str []rune
	//begin with ", end with "
	r := s.read()
	if r != '"' {
		os.Exit(1)
	}

	for {
		r := s.read()
		if r == '"' || r == 1 {
			break
		}
		str = append(str, r)
	}
	return string(str)
}

func (s *Scanner) scanNumber() interface{} {
	var number []rune
	var isFloat bool
	for {
		r := s.read()
		if r == '.' && len(number) > 0 && !isFloat {
			isFloat = true
			number = append(number, r)
			continue
		}

		if isWhitespace(r) || r == ',' || r == '}' || r == ']' {
			s.unread()
			break
		}
		if !isDigit(r) {
			return nil
		}
		number = append(number, r)
	}
	if isFloat {
		f, _ := strconv.ParseFloat(string(number), 64)
		return f
	}
	i, _ := strconv.Atoi(string(number))
	return i
}

func (s *Scanner) read() rune {
	ch, _, _ := s.buf.ReadRune()
	return ch
}

func (s *Scanner) unread() { _ = s.buf.UnreadRune() }

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
