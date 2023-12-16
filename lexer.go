package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type JsonLexer interface {
	Execute(rawJson string) ([]JsonToken, error)
}

type SimpleJsonLexer struct{}

func (l SimpleJsonLexer) Execute(rawJson string) ([]JsonToken, error) {
	var result []JsonToken

	for i := 0; i < len(rawJson); i++ {
		if rawJson[i] == '{' {
			result = append(result, JsonToken{virtualType: LeftCurlyParenthesis, content: "{"})
		}

		if rawJson[i] == '}' {
			result = append(result, JsonToken{virtualType: RightCurlyParenthesis, content: "}"})
		}

		if rawJson[i] == '[' {
			result = append(result, JsonToken{virtualType: LeftSquaredParenthesis, content: "["})
		}

		if rawJson[i] == ']' {
			fmt.Println("clos")
			result = append(result, JsonToken{virtualType: RightSquaredParenthesis, content: "]"})
		}

		if rawJson[i] == ':' {
			result = append(result, JsonToken{virtualType: Semicolon, content: ":"})
		}

		if rawJson[i] == ',' {
			result = append(result, JsonToken{virtualType: Comma, content: ","})
		}

		if rawJson[i] == '"' {
			closingQuoteIndex, closingQuoteIndexError := closingQuoteIndex(rawJson, i)

			if closingQuoteIndexError != nil {
				return nil, closingQuoteIndexError
			}

			result = append(result, JsonToken{virtualType: String, content: rawJson[i+1 : closingQuoteIndex]})
			i = closingQuoteIndex
		}

		if (i+4) <= len(rawJson) && rawJson[i:i+4] == "true" {
			result = append(result, JsonToken{virtualType: Boolean, content: true})
			i = i + 3
		}

		if (i+5) <= len(rawJson) && rawJson[i:i+5] == "false" {
			result = append(result, JsonToken{virtualType: Boolean, content: false})
			i = i + 4
		}

		if (i+4) <= len(rawJson) && rawJson[i:i+4] == "null" {
			result = append(result, JsonToken{virtualType: Null, content: nil})
			i = i + 3
		}

		if unicode.IsNumber(rune(rawJson[i])) {
			lastDigitIndex := lastDigitIndex(rawJson, i)

			number, atoiError := strconv.Atoi(rawJson[i : lastDigitIndex+1])

			if atoiError != nil {
				return nil, atoiError
			}

			result = append(result, JsonToken{virtualType: Number, content: number})
			i = lastDigitIndex
		}

	}

	fmt.Printf("%+v\n", result)
	return result, nil
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

func closingQuoteIndex(rawJson string, firstQuoteIndex int) (int, error) {
	start := firstQuoteIndex + 1

	end := strings.Index(rawJson[start:], "\"")
	if end == -1 {
		return -1, errors.New("cannot find closing quote")
	}

	return start + end, nil
}

type JsonToken struct {
	virtualType string
	content     interface{}
}

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
