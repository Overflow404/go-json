package main

type JsonToken struct {
	virtualType string
	content     interface{}
}

type JsonLexer interface {
	Execute(rawJson string) []JsonToken
}

type SimpleJsonLexer struct{}

var handlers = []LexerHandler{
	LeftCurlyParenthesisHandler{},
	RightCurlyParenthesisHandler{},
	LeftSquaredParenthesisHandler{},
	RightSquaredParenthesisHandler{},
	ColonHandler{},
	CommaHandler{},
	StringHandler{},
	TrueHandler{},
	FalseHandler{},
	NumberHandler{},
	NullHandler{},
}

func (l SimpleJsonLexer) Execute(rawJson string) []JsonToken {
	var result []JsonToken

	for i := 0; i < len(rawJson); {
		handler := lookupHandler(rawJson, i)
		result, i = handler.Apply(rawJson, result, i)
	}

	return result
}

func lookupHandler(rawJson string, i int) LexerHandler {
	for _, handler := range handlers {

		if handler.Eval(rawJson, i) {
			return handler
		}
	}

	return SkipHandler{}
}
