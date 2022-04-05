package kvalscanner

//Customised from: https://github.com/benbjohnson/sql-parser/

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strconv"
	"strings"
)

var minUnicoderune rune
var maxUnicoderune rune

// Initialize the code...
func init() {
	var err error
	//UNICODE: http://unicode-table.com/en/
	minUnicoderune, _, _, err = strconv.UnquoteChar("\u00A1", 0) //inverted exclamation mark
	if err != nil {
		log.Fatal("Cannot initialize scanner with min unicode value.")
	}
	maxUnicoderune, _, _, err = strconv.UnquoteChar("\uFF1F00", 0) //extended symbols: squid
	if err != nil {
		log.Fatal("Cannot initialize scanner with max unicode value.")
	}
}

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an literal or reserved word.
	// If we see a digit then consume as a number.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isOperator(ch) {
		s.unread()
		return s.scanOperator()
	} else if isLetter(ch) {
		s.unread()
		return s.scanLiteral()
	} else if isDigit(ch) {
		s.unread()
		return s.scanLiteral()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '_':
		return USCORE, string(ch)
	case '{':
		return OPATT, string(ch)
	case '}':
		return CPATT, string(ch)
	}
	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanOperator() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent operator character into the buffer.
	// Non-operator characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isOperator(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a custom operator then return that operator.
	switch strings.ToUpper(buf.String()) {
	case ">>>>":
		return BUCKEY, buf.String()
	case ">>":
		return BUCBUC, buf.String()
	case "=>":
		return ASSIGN, buf.String()
	case "::":
		return KEYVAL, buf.String()
	}

	return LITERAL, buf.String()
}

// scanLiteral consumes the current rune and all contiguous literal runes.
func (s *Scanner) scanLiteral() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent literal character into the buffer.
	// Non-literal characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "INS":
		return INS, buf.String()
	case "GET":
		return GET, buf.String()
	case "LIS":
		return LIS, buf.String()
	case "DEL":
		return DEL, buf.String()
	case "REN":
		return REN, buf.String()
	}

	// Otherwise return as a regular identifier.
	return LITERAL, buf.String()
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' }

// isLetter returns true if the rune is a letter.
//TODO: Need to expand this for a greater range of values... maybe NOT(the other classes?)
//func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '-' }

var asciiLetterSymbols = []rune{';', '=', '?', '[', '\\', ']', '^', '`', '|', '~', '@'}

func isLetter(ch rune) bool {

	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= '!' && ch <= '/' {
		return true
	}
	for _, v := range asciiLetterSymbols {
		if ch == v {
			return true
		}
	}
	//some rudimentary unicode handling... (need to improve)
	//init code sets min and max unicode values
	if ch >= minUnicoderune && ch <= maxUnicoderune {
		return true
	}
	return false
}

func isOperator(ch rune) bool { return (ch == '>') || (ch == ':') || (ch == '=') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
