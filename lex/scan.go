package lex

import (
	"errors"
	"fmt"
	"os"

	"github.com/killthebuddh4/gadflai/types"
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
	case ".", "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
	IdentifierLoop:
		for {
			n, _ := s.readCurrent()

			switch n {
			case ".", "_", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
				s.advance()
			default:
				break IdentifierLoop
			}
		}
	case "G", "D", "O":
		first, _ := s.readCurrent()
		s.advance()
		second, _ := s.readCurrent()
		s.advance()
		third, _ := s.readCurrent()
		s.advance()
		switch first + second + third {
		case "GAD", "DAE", "ORA":
			// FLY, MON, CLE
			s.advance()
			s.advance()
			s.advance()
		case "GHO":
			// ST
			s.advance()
			s.advance()
		default:
			return errors.New("unexpected AI keyword prefix, got " + first + second + third)
		}

	default:
		return errors.New("unexpected character " + char)
	}

	_, debug := os.LookupEnv("GADFLY_DEBUG_LEX")

	if debug {
		fmt.Println("adding token <" + s.Source[s.Start:s.Current] + ">")
	}

	s.Tokens = append(s.Tokens, types.Lexeme{
		Start:  s.Start,
		Length: s.Current - s.Start,
		Text:   s.Source[s.Start:s.Current],
	})

	return nil
}
