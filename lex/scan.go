package lex

import (
	"errors"
	"fmt"
	"os"
)

func (s *Lexer) scan() error {
	char, err := s.readCurrent()

	if err != nil {
		return err
	}

	switch char {
	case "\n", "\r", "\t", " ":
		s.advance()
		return nil
	case "#":
		n, _ := s.readCurrent()

		for !s.isAtEnd() {
			if n == "\n" {
				break
			}
			s.advance()
			n, _ = s.readCurrent()
		}

		return nil
	case "@":
		s.advance()
	case "-":
		s.advance()
	case "+":
		s.advance()
	case "*":
		s.advance()
	case "/":
		s.advance()
	case "\"":
		s.advance()

		for !s.isAtEnd() {
			n, _ := s.readCurrent()

			if n == "\"" {
				break
			} else {
				s.advance()
			}
		}

		if s.isAtEnd() {
			return errors.New("unexpected end of file, unterminated string")
		}

		s.advance()
	case "|":
		n, _ := s.readNext()

		if n != "|" {
			s.advance()
		} else {
			s.advance()
			s.advance()
		}
	case "&":
		n, _ := s.readNext()

		if n != "&" {
			return errors.New("unexpected character, expected '&' after '&'")
		} else {
			s.advance()
			s.advance()
		}
	case "!":
		n, _ := s.readNext()

		if n != "=" {
			s.advance()
		} else {
			s.advance()
			s.advance()
		}
	case "=":
		n, _ := s.readNext()

		if n != "=" {
			return errors.New("unexpected character, expected '='")
		} else {
			s.advance()
			s.advance()
		}
	case "<":
		n, _ := s.readNext()

		if n != "=" {
			s.advance()
		} else {
			s.advance()
			s.advance()
		}

	case ">":
		n, _ := s.readNext()

		if n != "=" {
			s.advance()
		} else {
			s.advance()
			s.advance()
		}
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		for !s.isAtEnd() {
			c, _ := s.readCurrent()

			if c < "0" || c > "9" {
				break
			} else {
				s.advance()
			}
		}

		c, _ := s.readCurrent()

		if c != "." {
			break
		} else {
			s.advance()
		}

		c, _ = s.readCurrent()

		if c < "0" && c > "9" {
			return errors.New("unexpected character, expected digit after decimal")
		}

		for !s.isAtEnd() {
			c, _ := s.readCurrent()

			if c < "0" || c > "9" {
				break
			} else {
				s.advance()
			}
		}
	case "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
	IdentifierLoop:
		for {
			n, _ := s.readCurrent()

			switch n {
			case "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
				s.advance()
			default:
				break IdentifierLoop
			}
		}
	default:
		return errors.New("unexpected character " + char)
	}

	_, debug := os.LookupEnv("GADFLY_DEBUG_LEX")

	if debug {
		fmt.Println("adding token <" + s.Source[s.Start:s.Current] + ">")
	}

	s.Tokens = append(s.Tokens, Lexeme{
		Start:  s.Start,
		Length: s.Current - s.Start,
		Text:   s.Source[s.Start:s.Current],
	})

	return nil
}
