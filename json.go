package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Filename is required")
	}

	bytes, readFileError := os.ReadFile(os.Args[1])
	if readFileError != nil {
		log.Fatalf(readFileError.Error())
	}

	run(string(bytes), SimpleJsonLexer{})
}

func run(rawJson string, lexer JsonLexer) {
	fmt.Println(prettyPrint(lexer.Execute(rawJson)))
}

func prettyPrint(tokens []JsonToken) string {
	var result strings.Builder
	indent := 0

	for _, token := range tokens {
		switch token.virtualType {
		case LeftCurlyParenthesis:
			result.WriteString(fmt.Sprintf("%s {\n%s", redColor("Object"), strings.Repeat("  ", indent+1)))
			indent++
		case LeftSquaredParenthesis:
			result.WriteString(fmt.Sprintf("%s [\n%s", redColor("Array"), strings.Repeat("  ", indent+1)))
			indent++
		case RightCurlyParenthesis:
			indent--
			result.WriteString(fmt.Sprintf("\n%s}", strings.Repeat("  ", indent)))
		case RightSquaredParenthesis:
			indent--
			result.WriteString(fmt.Sprintf("\n%s]", strings.Repeat("  ", indent)))
		case Comma:
			result.WriteString(",\n")
			result.WriteString(strings.Repeat("  ", indent))
		case Semicolon:
			result.WriteString(": ")
		default:
			result.WriteString(fmt.Sprintf("(%s %v)", blueColor(token.virtualType), token.content))
		}
	}

	return result.String()
}

func redColor(content interface{}) string {
	return fmt.Sprintf("%s%v\x1b[0m", "\x1b[31m", content)
}

func blueColor(content interface{}) string {
	return fmt.Sprintf("%s%v\x1b[0m", "\x1b[34m", content)
}
