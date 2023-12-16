package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

const (
	LeftCurlyParenthesis    string = "LeftCurlyParenthesis"
	RightCurlyParenthesis   string = "RightCurlyParenthesis"
	LeftSquaredParenthesis  string = "LeftSquaredParenthesis"
	RightSquaredParenthesis string = "RightSquaredParenthesis"
	Semicolon               string = "Semicolon"
	Comma                   string = "Comma"
	String                  string = "String"
	Boolean                 string = "Boolean"
	Number                  string = "Number"
	Null                    string = "Null"
)

type LexerHandler interface {
	Eval(rawJson string, base int) bool
	Apply(rawJson string, tokens []JsonToken, base int) ([]JsonToken, int)
}

type LeftCurlyParenthesisHandler struct{}
type RightCurlyParenthesisHandler struct{}
type LeftSquaredParenthesisHandler struct{}
type RightSquaredParenthesisHandler struct{}
type SemicolonHandler struct{}
type CommaHandler struct{}
type StringHandler struct{}
type TrueHandler struct{}
type FalseHandler struct{}
type NumberHandler struct{}
type NullHandler struct{}
type SkipHandler struct{}

func (h LeftCurlyParenthesisHandler) Eval(rawJson string, base int) bool {
	return rawJson[base] == '{'
}

func (h LeftCurlyParenthesisHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: LeftCurlyParenthesis, content: "{"}), advancePointerTo(base, 1)
}

func (h RightCurlyParenthesisHandler) Eval(rawJson string, base int) bool {
	return rawJson[base] == '}'
}

func (h RightCurlyParenthesisHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: RightCurlyParenthesis, content: "}"}), advancePointerTo(base, 1)
}

func (h LeftSquaredParenthesisHandler) Eval(rawJson string, base int) bool {
	return rawJson[base] == '['
}

func (h LeftSquaredParenthesisHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: LeftSquaredParenthesis, content: "["}), advancePointerTo(base, 1)
}

func (h RightSquaredParenthesisHandler) Eval(rawJson string, base int) bool {
	return rawJson[base] == ']'
}

func (h RightSquaredParenthesisHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: RightSquaredParenthesis, content: "]"}), advancePointerTo(base, 1)
}

func (h SemicolonHandler) Eval(rawJson string, base int) bool {
	return rawJson[base] == ':'
}

func (h SemicolonHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: Semicolon, content: ":"}), advancePointerTo(base, 1)
}

func (h CommaHandler) Eval(rawJson string, base int) bool {
	return rawJson[base] == ','
}

func (h CommaHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: Comma, content: ","}), advancePointerTo(base, 1)
}

func (h StringHandler) Eval(rawJson string, base int) bool {
	return rawJson[base] == '"'
}

func (h StringHandler) Apply(rawJson string, tokens []JsonToken, base int) ([]JsonToken, int) {
	closingQuoteIndex := closingQuoteIndex(rawJson, base)
	return append(tokens, JsonToken{virtualType: String, content: rawJson[base+1 : closingQuoteIndex]}), setPointerTo(closingQuoteIndex + 1)

}

func closingQuoteIndex(rawJson string, firstQuoteIndex int) int {
	start := firstQuoteIndex + 1

	end := strings.Index(rawJson[start:], "\"")
	if end == -1 {
		log.Fatalf("cannot find closing quote")
	}

	return start + end
}

func (h TrueHandler) Eval(rawJson string, base int) bool {
	return (base+4) <= len(rawJson) && rawJson[base:base+4] == "true"
}

func (h TrueHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: Boolean, content: true}), setPointerTo(base + 4)
}

func (h FalseHandler) Eval(rawJson string, base int) bool {
	return (base+5) <= len(rawJson) && rawJson[base:base+5] == "false"
}

func (h FalseHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: Boolean, content: false}), setPointerTo(base + 5)
}

func (h NullHandler) Eval(rawJson string, base int) bool {
	return (base+4) <= len(rawJson) && rawJson[base:base+4] == "null"
}

func (h NullHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return append(tokens, JsonToken{virtualType: Null, content: nil}), setPointerTo(base + 4)
}

func (h NumberHandler) Eval(rawJson string, base int) bool {
	return unicode.IsNumber(rune(rawJson[base]))
}

func (h NumberHandler) Apply(rawJson string, tokens []JsonToken, base int) ([]JsonToken, int) {
	lastDigitIndex := lastDigitIndex(rawJson, base)

	number, atoiError := strconv.Atoi(rawJson[base : lastDigitIndex+1])

	if atoiError != nil {
		log.Fatalf("%v", atoiError)
	}

	return append(tokens, JsonToken{virtualType: Number, content: number}), setPointerTo(lastDigitIndex + 1)
}

func lastDigitIndex(rawJson string, start int) int {
	end := start

	for i := start; i < len(rawJson); i++ {
		if unicode.IsNumber(rune(rawJson[i])) {
			end = i
		} else {
			break
		}
	}

	return end
}

func (h SkipHandler) Eval(_ string, _ int) bool {
	return true
}

func (h SkipHandler) Apply(_ string, tokens []JsonToken, base int) ([]JsonToken, int) {
	return tokens, advancePointerTo(base, 1)
}

func advancePointerTo(base int, offset int) int {
	return base + offset
}

func setPointerTo(offset int) int {
	return offset
}
