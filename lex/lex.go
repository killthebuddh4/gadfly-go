package lex

import (
	"errors"

	"github.com/killthebuddh4/gadflai/keywords"
)

func Lex(source string) ([]Lexeme, error) {
	lexer := NewLexer(source)

	for !lexer.isAtEnd() {
		lexer.Start = lexer.Current
		err := lexer.scanToken()

		if err != nil {
			return nil, err
		}
	}

	eof, err := NewEof(source)

	if err != nil {
		return nil, err
	}

	lexer.Tokens = append(lexer.Tokens, eof)

	return lexer.Tokens, nil
}

func (s *Lexer) scanToken() error {
	c, err := s.readCurrent()

	if err != nil {
		return err
	}

	if keywords.IsSpecial(string(c)) {
		return scanSpecial(s, string(c))
	} else if keywords.IsNumber(string(c)) {
		return scanNumber(s, string(c))
	} else if keywords.IsWhitespace(string(c)) {
		return scanWhitespace(s)
	} else if keywords.IsIdentifier(string(c)) {
		case IDENTIFIERS.At:
			s.advance()
			block, err := GetBlockType(s.readLexeme())
			if err != nil {
				return err
			}
			s.addToken(block)
		case IDENTIFIERS.Underscore, IDENTIFIERS.LowerA, IDENTIFIERS.LowerB, IDENTIFIERS.LowerC, IDENTIFIERS.LowerD, IDENTIFIERS.LowerE, IDENTIFIERS.LowerF, IDENTIFIERS.LowerG, IDENTIFIERS.LowerH, IDENTIFIERS.LowerI, IDENTIFIERS.LowerJ, IDENTIFIERS.LowerK, IDENTIFIERS.LowerL, IDENTIFIERS.LowerM, IDENTIFIERS.LowerN, IDENTIFIERS.LowerO, IDENTIFIERS.LowerP, IDENTIFIERS.LowerQ, IDENTIFIERS.LowerR, IDENTIFIERS.LowerS, IDENTIFIERS.LowerT, IDENTIFIERS.LowerU, IDENTIFIERS.LowerV, IDENTIFIERS.LowerW, IDENTIFIERS.LowerX, IDENTIFIERS.LowerY, IDENTIFIERS.LowerZ, IDENTIFIERS.UpperA, IDENTIFIERS.UpperB, IDENTIFIERS.UpperC, IDENTIFIERS.UpperD, IDENTIFIERS.UpperE, IDENTIFIERS.UpperF, IDENTIFIERS.UpperG, IDENTIFIERS.UpperH, IDENTIFIERS.UpperI, IDENTIFIERS.UpperJ, IDENTIFIERS.UpperK, IDENTIFIERS.UpperL, IDENTIFIERS.UpperM, IDENTIFIERS.UpperN, IDENTIFIERS.UpperO, IDENTIFIERS.UpperP, IDENTIFIERS.UpperQ, IDENTIFIERS.UpperR, IDENTIFIERS.UpperS, IDENTIFIERS.UpperT, IDENTIFIERS.UpperU, IDENTIFIERS.UpperV, IDENTIFIERS.UpperW, IDENTIFIERS.UpperX, IDENTIFIERS.UpperY, IDENTIFIERS.UpperZ:
			s.advanceIdentifier()

			block, err := GetBlockType(s.readLexeme())

			if err == nil {
				s.addToken(block)
			} else {
				switch s.readLexeme() {
				case "true":
					s.addToken(keywords.TOKENS.True)
				case "false":
					s.addToken(keywords.TOKENS.False)
				case "nil":
					s.addToken(keywords.TOKENS.Nil)
				default:
					s.addToken(keywords.TOKENS.Identifier)
				}
			}
		default:
			return errors.New("unexpected character" + string(c))
		}
	}

	return nil
}

func scanWhitespace(s *Lexer) error {
	s.advance()
}

func scanNumber(s *Lexer, c string) error {
	s.advanceNumber()
	s.addToken(keywords.TOKENS.Number)
	return nil
}

func scanSpecial(s Lexer, c string) error {
	switch c {
	case keywords.SPECIALS.Minus:
		s.advance()
		s.addToken(keywords.TOKENS.Minus)
	case keywords.SPECIALS.Plus:
		s.advance()
		s.addToken(keywords.TOKENS.Plus)
	case keywords.SPECIALS.Multiply:
		s.advance()
		s.addToken(keywords.TOKENS.Multiply)
	case keywords.SPECIALS.Pipe:
		n, _ := s.readNext()

		if string(n) != keywords.SPECIALS.Pipe {
			s.advance()
			s.addToken(keywords.SPECIALS.Pipe)
		} else {
			s.advance()
			s.advance()
			s.addToken(keywords.TOKENS.Disjunction)
		}
	case keywords.SPECIALS.Ampersand:
		n, _ := s.readNext()

		if string(n) != keywords.SPECIALS.Ampersand {
			return errors.New("unexpected character, expected '&' after '&'")
		} else {
			s.advance()
			s.advance()
			s.addToken(keywords.TOKENS.Conjunction)
		}
	case keywords.SPECIALS.Bang:
		n, _ := s.readNext()

		if string(n) != keywords.SPECIALS.Equal {
			s.advance()
			s.addToken(keywords.TOKENS.Bang)
		} else {
			s.advance()
			s.advance()
			s.addToken(keywords.TOKENS.BangEqual)
		}
	case keywords.SPECIALS.Equal:
		n, _ := s.readNext()

		if string(n) != keywords.SPECIALS.Equal {
			return errors.New("unexpected character, expected '='")
		} else {
			s.advance()
			s.advance()
			s.addToken(keywords.TOKENS.EqualEqual)
		}
	case keywords.SPECIALS.LessThan:
		n, _ := s.readNext()

		if string(n) != keywords.SPECIALS.Equal {
			s.advance()
			s.addToken(keywords.TOKENS.LessThan)
		} else {
			s.advance()
			s.advance()
			s.addToken(keywords.TOKENS.LessThanEqual)
		}
	case keywords.SPECIALS.GreaterThan:
		n, _ := s.readNext()

		if string(n) != keywords.SPECIALS.Equal {
			s.advance()
			s.addToken(keywords.TOKENS.GreaterThan)
		} else {
			s.advance()
			s.advance()
			s.addToken(keywords.TOKENS.GreaterThanEqual)
		}
	case keywords.SPECIALS.Comment:
		s.advanceLine()
	case keywords.SPECIALS.Divide:
		s.advance()
		s.addToken(keywords.TOKENS.Divide)
	case keywords.SPECIALS.Quote:
		s.advanceString()
		s.addToken(keywords.TOKENS.String)
	default:
		return errors.New("unexpected character" + string(c))
	}

	return nil
}

func (s *Lexer) addToken(tokenType string) {
	s.Tokens = append(s.Tokens, Lexeme{
		Keyword: tokenType,
		Start:   s.Start,
		Length:  s.Current - s.Start,
		Text:    s.readLexeme(),
	})
}

func (s *Lexer) advance() error {
	if s.isAtEnd() {
		return errors.New("unexpected end of file")
	}
	s.Current++

	return nil
}

func (s *Lexer) advanceNumber() error {
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

func (s *Lexer) advanceString() error {

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

func (s *Lexer) advanceIdentifier() error {
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

func (s *Lexer) advanceLine() {
	n, _ := s.readCurrent()

	for !s.isAtEnd() {
		if n == '\n' {
			break
		}
		s.advance()
		n, _ = s.readCurrent()
	}
}

func (s Lexer) readLexeme() string {
	return s.Source[s.Start:s.Current]
}

func (s Lexer) readCurrent() (rune, error) {
	if s.isAtEnd() {
		return ' ', errors.New("unexpected end of file")
	}
	return rune(s.Source[s.Current]), nil
}

func (s Lexer) readNext() (rune, error) {
	if s.Current+1 > len(s.Source) {
		return ' ', errors.New("unexpected end of file")
	}
	return rune(s.Source[s.Current+1]), nil
}

func (s Lexer) isAtEnd() bool {
	return s.Current >= len(s.Source)
}
