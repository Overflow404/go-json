package main

import (
	"reflect"
	"testing"
)

func TestSimpleJsonLexer_Execute(t *testing.T) {
	type args struct {
		rawJson string
	}
	tests := []struct {
		name string
		args args
		want []JsonToken
	}{
		{
			"should lex an empty json object",
			args{"{}"},
			[]JsonToken{
				{virtualType: LeftCurlyParenthesis, content: "{"},
				{virtualType: RightCurlyParenthesis, content: "}"},
			},
		},
		{
			"should lex an empty json array",
			args{"[]"},
			[]JsonToken{
				{virtualType: LeftSquaredParenthesis, content: "["},
				{virtualType: RightSquaredParenthesis, content: "]"},
			},
		},
		{
			"should lex a comma",
			args{","},
			[]JsonToken{
				{virtualType: Comma, content: ","},
			},
		},
		{
			"should lex a semicolon",
			args{":"},
			[]JsonToken{
				{virtualType: Semicolon, content: ":"},
			},
		},
		{
			"should lex a string",
			args{"\"stringContentHere\""},
			[]JsonToken{
				{virtualType: String, content: "stringContentHere"},
			},
		},
		{
			"should lex an empty string",
			args{"\"\""},
			[]JsonToken{
				{virtualType: String, content: ""},
			},
		},
		{
			"should lex an array with two strings",
			args{"[\"stringContentHere1\", \"stringContentHere2\"]"},
			[]JsonToken{
				{virtualType: LeftSquaredParenthesis, content: "["},
				{virtualType: String, content: "stringContentHere1"},
				{virtualType: Comma, content: ","},
				{virtualType: String, content: "stringContentHere2"},
				{virtualType: RightSquaredParenthesis, content: "]"},
			},
		},
		{
			"should lex an object of strings",
			args{"{\"key1\": \"value1\", \"key2\": \"value2\"}"},
			[]JsonToken{
				{virtualType: LeftCurlyParenthesis, content: "{"},
				{virtualType: String, content: "key1"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: String, content: "value1"},
				{virtualType: Comma, content: ","},
				{virtualType: String, content: "key2"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: String, content: "value2"},
				{virtualType: RightCurlyParenthesis, content: "}"},
			},
		},
		{
			"should lex a boolean value",
			args{"[true, false]"},
			[]JsonToken{
				{virtualType: LeftSquaredParenthesis, content: "["},
				{virtualType: Boolean, content: true},
				{virtualType: Comma, content: ","},
				{virtualType: Boolean, content: false},
				{virtualType: RightSquaredParenthesis, content: "]"},
			},
		},
		{
			"should lex a number",
			args{"{\"key\": 1}"},
			[]JsonToken{
				{virtualType: LeftCurlyParenthesis, content: "{"},
				{virtualType: String, content: "key"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: Number, content: 1},
				{virtualType: RightCurlyParenthesis, content: "}"},
			},
		},
		{
			"should lex a big number",
			args{"{\"key\": 1234567}"},
			[]JsonToken{
				{virtualType: LeftCurlyParenthesis, content: "{"},
				{virtualType: String, content: "key"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: Number, content: 1234567},
				{virtualType: RightCurlyParenthesis, content: "}"},
			},
		},
		{
			"should lex an object with multiple types",
			args{"{\"key1\": 156, \"key2\": \"dummyString\", \"key3\": false}"},
			[]JsonToken{
				{virtualType: LeftCurlyParenthesis, content: "{"},
				{virtualType: String, content: "key1"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: Number, content: 156},
				{virtualType: Comma, content: ","},
				{virtualType: String, content: "key2"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: String, content: "dummyString"},
				{virtualType: Comma, content: ","},
				{virtualType: String, content: "key3"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: Boolean, content: false},
				{virtualType: RightCurlyParenthesis, content: "}"},
			},
		},
		{
			"should lex an array with multiple types",
			args{"[1, \"dummyString\", true]"},
			[]JsonToken{
				{virtualType: LeftSquaredParenthesis, content: "["},
				{virtualType: Number, content: 1},
				{virtualType: Comma, content: ","},
				{virtualType: String, content: "dummyString"},
				{virtualType: Comma, content: ","},
				{virtualType: Boolean, content: true},
				{virtualType: RightSquaredParenthesis, content: "]"},
			},
		},
		{
			"should lex an array with a nested object",
			args{"[1, \"dummyString\", true, {\"key\": false}]"},
			[]JsonToken{
				{virtualType: LeftSquaredParenthesis, content: "["},
				{virtualType: Number, content: 1},
				{virtualType: Comma, content: ","},
				{virtualType: String, content: "dummyString"},
				{virtualType: Comma, content: ","},
				{virtualType: Boolean, content: true},
				{virtualType: Comma, content: ","},
				{virtualType: LeftCurlyParenthesis, content: "{"},
				{virtualType: String, content: "key"},
				{virtualType: Semicolon, content: ":"},
				{virtualType: Boolean, content: false},
				{virtualType: RightCurlyParenthesis, content: "}"},
				{virtualType: RightSquaredParenthesis, content: "]"},
			},
		},
		{
			"should lex a null value",
			args{"[null, false]"},
			[]JsonToken{
				{virtualType: LeftSquaredParenthesis, content: "["},
				{virtualType: Null, content: nil},
				{virtualType: Comma, content: ","},
				{virtualType: Boolean, content: false},
				{virtualType: RightSquaredParenthesis, content: "]"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := SimpleJsonLexer{}
			got, lexerError := l.Execute(tt.args.rawJson)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nExpected = %+v, \nActual = %+v", got, tt.want)
			}

			if lexerError != nil {
				t.Errorf("Error = %v", lexerError)
			}
		})
	}
}
