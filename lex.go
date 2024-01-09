package main

import (
	"errors"
	"fmt"
)

func Lex(source string) ([]Token, error) {
	scanner := Scanner{
		Source:  source,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
	}

	for !scanner.isAtEnd() {
		scanner.Start = scanner.Current
		err := scanner.scanToken()

		if err != nil {
			return nil, err
		}
	}

	eof, err := EofToken(source)

	if err != nil {
		return nil, err
	}

	scanner.Tokens = append(scanner.Tokens, eof)

	return scanner.Tokens, nil
}

type Scanner struct {
	Source  string
	Tokens  []Token
	Start   int
	Current int
}

func (s *Scanner) scanToken() error {
	c, err := s.readCurrent()

	if err != nil {
		return err
	}

	switch c {
	case OPERATORS.Minus:
		s.advance()
		s.addToken(TOKENS.Minus)
	case OPERATORS.Plus:
		s.advance()
		s.addToken(TOKENS.Plus)
	case OPERATORS.Multiply:
		s.advance()
		s.addToken(TOKENS.Multiply)
	case OPERATORS.Pipe:
		n, _ := s.readNext()

		if n == OPERATORS.Pipe {
			s.advance()
			s.advance()
			s.addToken(TOKENS.Disjunction)
		} else {
			s.advance()
			s.addToken(TOKENS.Pipe)
		}
	case OPERATORS.Ampersand:
		n, _ := s.readNext()

		if n != OPERATORS.Ampersand {
			return errors.New("unexpected character, expected '&' after '&'")
		} else {
			s.advance()
			s.advance()
			s.addToken(TOKENS.Conjunction)
		}
	case OPERATORS.Bang:
		n, _ := s.readNext()

		if n != OPERATORS.Equal {
			s.advance()
			s.addToken(TOKENS.Bang)
		} else {
			s.advance()
			s.advance()
			s.addToken(TOKENS.BangEqual)
		}
	case OPERATORS.Equal:
		n, _ := s.readNext()

		if n != OPERATORS.Equal {
			return errors.New("unexpected character, expected '='")
		} else {
			s.advance()
			s.advance()
			s.addToken(TOKENS.EqualEqual)
		}
	case OPERATORS.LessThan:
		n, _ := s.readNext()

		if n != OPERATORS.Equal {
			s.advance()
			s.addToken(TOKENS.LessThan)
		} else {
			s.advance()
			s.advance()
			s.addToken(TOKENS.LessThanEqual)
		}
	case OPERATORS.GreaterThan:
		n, _ := s.readNext()

		if n != OPERATORS.Equal {
			s.advance()
			s.addToken(TOKENS.GreaterThan)
		} else {
			s.advance()
			s.advance()
			s.addToken(TOKENS.GreaterThanEqual)
		}
	case COMMENT.Line:
		s.advanceLine()
	case OPERATORS.Divide:
		s.advance()
		s.addToken(TOKENS.Divide)
	case STRINGS.Quote:
		s.advanceString()
		s.addToken(TOKENS.String)
	case NUMBERS.Zero, NUMBERS.One, NUMBERS.Two, NUMBERS.Three, NUMBERS.Four, NUMBERS.Five, NUMBERS.Six, NUMBERS.Seven, NUMBERS.Eight, NUMBERS.Nine:
		s.advanceNumber()
		s.addToken(TOKENS.Number)
	case WHITESPACE.NewLine, WHITESPACE.Space, WHITESPACE.Tab, WHITESPACE.Return:
		s.advance()
	case IDENTIFIERS.Underscore, IDENTIFIERS.LowerA, IDENTIFIERS.LowerB, IDENTIFIERS.LowerC, IDENTIFIERS.LowerD, IDENTIFIERS.LowerE, IDENTIFIERS.LowerF, IDENTIFIERS.LowerG, IDENTIFIERS.LowerH, IDENTIFIERS.LowerI, IDENTIFIERS.LowerJ, IDENTIFIERS.LowerK, IDENTIFIERS.LowerL, IDENTIFIERS.LowerM, IDENTIFIERS.LowerN, IDENTIFIERS.LowerO, IDENTIFIERS.LowerP, IDENTIFIERS.LowerQ, IDENTIFIERS.LowerR, IDENTIFIERS.LowerS, IDENTIFIERS.LowerT, IDENTIFIERS.LowerU, IDENTIFIERS.LowerV, IDENTIFIERS.LowerW, IDENTIFIERS.LowerX, IDENTIFIERS.LowerY, IDENTIFIERS.LowerZ, IDENTIFIERS.UpperA, IDENTIFIERS.UpperB, IDENTIFIERS.UpperC, IDENTIFIERS.UpperD, IDENTIFIERS.UpperE, IDENTIFIERS.UpperF, IDENTIFIERS.UpperG, IDENTIFIERS.UpperH, IDENTIFIERS.UpperI, IDENTIFIERS.UpperJ, IDENTIFIERS.UpperK, IDENTIFIERS.UpperL, IDENTIFIERS.UpperM, IDENTIFIERS.UpperN, IDENTIFIERS.UpperO, IDENTIFIERS.UpperP, IDENTIFIERS.UpperQ, IDENTIFIERS.UpperR, IDENTIFIERS.UpperS, IDENTIFIERS.UpperT, IDENTIFIERS.UpperU, IDENTIFIERS.UpperV, IDENTIFIERS.UpperW, IDENTIFIERS.UpperX, IDENTIFIERS.UpperY, IDENTIFIERS.UpperZ:
		s.advanceIdentifier()

		block, err := GetBlock(s.readLexeme())

		if err == nil {
			s.addToken(block)
		} else {
			switch s.readLexeme() {
			case "true":
				s.addToken(TOKENS.True)
			case "false":
				s.addToken(TOKENS.False)
			case "nil":
				s.addToken(TOKENS.Nil)
			default:
				s.addToken(TOKENS.Identifier)
			}
		}
	default:
		return errors.New("unexpected character" + string(c))
	}

	return nil
}

func (s *Scanner) addToken(tokenType string) {
	s.Tokens = append(s.Tokens, Token{
		Type:   tokenType,
		Start:  s.Start,
		Length: s.Current - s.Start,
		Lexeme: s.readLexeme(),
	})
}

func (s *Scanner) advance() error {
	if s.isAtEnd() {
		return errors.New("unexpected end of file")
	}
	s.Current++

	return nil
}

func (s *Scanner) advanceNumber() error {
	for !s.isAtEnd() {
		c, _ := s.readCurrent()

		if c < '0' || c > '9' {
			break
		} else {
			s.advance()
		}
	}

	c, _ := s.readCurrent()

	if c != '.' {
		return nil
	} else {
		s.advance()
	}

	c, _ = s.readCurrent()

	if c < '0' && c > '9' {
		return errors.New("unexpected character, expected digit after decimal")
	}

	for !s.isAtEnd() {
		c, _ := s.readCurrent()

		if c < '0' || c > '9' {
			break
		} else {
			s.advance()
		}
	}

	return nil
}

func (s *Scanner) advanceString() error {

	s.advance()

	for !s.isAtEnd() {
		n, _ := s.readCurrent()

		if n == '"' {
			break
		} else {
			s.advance()
		}
	}

	if s.isAtEnd() {
		return errors.New("unexpected end of file, unterminated string")
	}

	s.advance()

	return nil
}

func (s *Scanner) advanceIdentifier() error {
	for {
		n, _ := s.readCurrent()

		if !IsIdentifierChar(n) {
			break
		} else {
			s.advance()
		}
	}

	return nil
}

func (s *Scanner) advanceLine() {
	fmt.Println("advancing line")
	n, _ := s.readCurrent()

	for !s.isAtEnd() {
		if n == '\n' {
			break
		}
		s.advance()
		n, _ = s.readCurrent()
	}
}

func (s Scanner) readLexeme() string {
	return s.Source[s.Start:s.Current]
}

func (s Scanner) readCurrent() (rune, error) {
	if s.isAtEnd() {
		return ' ', errors.New("unexpected end of file")
	}
	return rune(s.Source[s.Current]), nil
}

func (s Scanner) readNext() (rune, error) {
	if s.Current+1 > len(s.Source) {
		return ' ', errors.New("unexpected end of file")
	}
	return rune(s.Source[s.Current+1]), nil
}

func (s Scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}
