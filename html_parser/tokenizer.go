package html_parser

import (
	"bufio"
	"io"
	"log/slog"
	"os"
)

type TokenType int

const (
	StartTag TokenType = iota
	EndTag
	Text
)

type Token struct {
	Type  TokenType
	Tag   string
	Attrs map[string]string
}

type State int

const (
	Data = iota
	TagOpen
	TagName
	AttrName
	AttrValue
	EndTagOpen
)

type TagType string

const (
	StartTagType TagType = "start"
	EndTagType   TagType = "end"
)

func NewToken() Token {
	return Token{
		Attrs: make(map[string]string),
	}
}

func Tokenize(rd io.Reader) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	reader := bufio.NewReader(rd)

	var state State = Data
	var buff string
	token := NewToken()
	var quoteChar rune
	currentAttrName := ""

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		switch state {
		case Data:
			if r == '<' {
				logger.Debug("Tag open", "rune", r)
				state = TagOpen
				if buff != "" {
					logger.Info("Text", "value", buff)
					buff = ""
				}
			} else {
				buff += string(r)
			}
		case TagOpen:
			if r == '/' {
				logger.Debug("End tag open")
				state = EndTagOpen
			}
			if r == 'a' {
				buff += string(r)
				state = TagName
			}
		case EndTagOpen:
			if r == '>' {
				logger.Debug("End tag closed")
				token.Tag = buff
				token.Type = EndTag
				printToken(token, EndTagType)
				token = NewToken()
				state = Data
			} else {
				buff += string(r)
			}
		case TagName:
			if r == ' ' {
				logger.Debug("Attribute started", "tag", token.Tag)
				token.Tag = buff
				state = AttrName
				buff = ""
			} else {
				buff += string(r)
			}
		case AttrName:
			if r == '=' {
				logger.Debug("Attribute value started", "name", buff)
				currentAttrName = buff
				state = AttrValue
				buff = ""
			} else {
				buff += string(r)
			}
		case AttrValue:
			switch r {
			case ' ':
				logger.Debug("Found space after attr value")
			case quoteChar:
				logger.Debug("Quote closed")
				token.Attrs[currentAttrName] = buff
				currentAttrName = ""
				buff = ""
				quoteChar = 0
			case '"', '\'':
				logger.Debug("Quote opened", "char", string(r))
				quoteChar = r
			case '>':
				logger.Debug("Tag closed")
				state = Data
				printToken(token, StartTagType)
				// emit start tag
			default:
				buff += string(r)
			}
		}
	}
}

func printToken(token Token, tagType TagType) {
	slog.Info("Token", "type", token.Type, "tagType", tagType, "tag", token.Tag, "attributes", token.Attrs)
}
