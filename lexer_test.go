package graphql_lexer

import (
	"testing"
)

func TestSDLLexerForTypeDef(t *testing.T) {
	input := `
	type Query {
		product(id!: Int): Product
	}


	type Product {
		id: Int!
		name: String!
		uri: String!
	}

`
	lex := NewLexer(input)

	tests := []struct{
		expectedLiteral string
		expectedType    TokenType
	}{
		{"type", TYPE},
		{"Query", QUERY},
		{"{", LBRACE},
		{"product", IDENTIFIER},
		{"(", LPAREN},
		{"id", IDENTIFIER},
		{"!", BANG},
		{":", COLON},
		{"Int", INT},
		{")", RPAREN},
		{":", COLON},
		{"Product", IDENTIFIER},
		{"}", RBRACE},
		{"type", TYPE},
		{"Product", IDENTIFIER},
		{"{", LBRACE},
		{"id", IDENTIFIER},
		{":", COLON},
		{"Int", INT},
		{"!", BANG},
		{"name", IDENTIFIER},
		{":", COLON},
		{"String", STRING},
		{"!", BANG},
		{"uri", IDENTIFIER},
		{":", COLON},
		{"String", STRING},
		{"!", BANG},
		{"}", RBRACE},
	}

	for i, ti := range tests {
		tok := NextToken()
		if tok.Type != ti.expectedType {
			t.Errorf("Failed to parse Token Type correctly. Expected %s Actual %s in test %d",
				ti.expectedType,
				tok.Type,
				i + 1)
		}

		if tok.Literal != ti.expectedLiteral{
			t.Errorf(
				"Failed to parse Token Literal correctly. Expected %s Actual %s in test %d",
				ti.expectedLiteral,
				tok.Literal,
				i + 1)
		}
	}
}


func TestSDLLexerForQuery(t *testing.T) {
	input := `
	query {
		product(id: "dummy") {
			id
			name
			uri
		}
	}
`

	lexer := NewLexer(input)

	tests := []struct{
		expectedLiteral string
		expectedType    TokenType
	}{
		{"query", QUERY},
		{"{", LBRACE},
		{"product", IDENTIFIER},
		{"(", LPAREN},
		{"id", IDENTIFIER},
		{":", COLON},
		{"dummy", STRING},
		{")", RPAREN},
		{"{", LBRACE},
		{"id", IDENTIFIER},
		{"name", IDENTIFIER},
		{"uri", IDENTIFIER},
		{"}", RBRACE},
		{"}", RBRACE},
	}

	for i, ti := range tests {
		tok := NextToken()
		if tok.Type != ti.expectedType {
			t.Errorf("Failed to parse Token Type correctly. Expected %s Actual %s in test %d",
				ti.expectedType,
				tok.Type,
				i + 1)
		}

		if tok.Literal != ti.expectedLiteral{
			t.Errorf(
				"Failed to parse Token Literal correctly. Expected %s Actual %s in test %d",
				ti.expectedLiteral,
				tok.Literal,
				i + 1)
		}
	}
}

