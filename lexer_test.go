package graphql_lexer

import (
	"testing"
)

func TestLexerForTypeDef(t *testing.T) {
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

	tests := []struct {
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
		tok := lex.NextToken()
		if tok.Type != ti.expectedType {
			t.Errorf("Failed to parse Token Type correctly. Expected %s Actual %s in test %d",
				ti.expectedType,
				tok.Type,
				i+1)
		}

		if tok.Literal != ti.expectedLiteral {
			t.Errorf(
				"Failed to parse Token Literal correctly. Expected %s Actual %s in test %d",
				ti.expectedLiteral,
				tok.Literal,
				i+1)
		}
	}
}

func TestLexerForUnionType(t *testing.T) {
	input := `
	type Cat {
		name: String!
	}

	type Dog {
		name: String!
	}

	union Animal = Cat | Dog
`

	lex := NewLexer(input)

	tests := []struct {
		expectedLiteral string
		expectedType    TokenType
	}{
		{"type", TYPE},
		{"Cat", IDENTIFIER},
		{"{", LBRACE},
		{"name", IDENTIFIER},
		{":", COLON},
		{"String", STRING},
		{"!", BANG},
		{"}", RBRACE},

		{"type", TYPE},
		{"Dog", IDENTIFIER},
		{"{", LBRACE},
		{"name", IDENTIFIER},
		{":", COLON},
		{"String", STRING},
		{"!", BANG},
		{"}", RBRACE},

		{"union", UNION},
		{"Animal", IDENTIFIER},
		{"=", EQUALS},
		{"Cat", IDENTIFIER},
		{"|", PIPE},
		{"Dog", IDENTIFIER},
	}

	for i, ti := range tests {
		tok := lex.NextToken()
		if tok.Type != ti.expectedType {
			t.Errorf("Failed to parse Token Type correctly. Expected %s Actual %s in test %d",
				ti.expectedType,
				tok.Type,
				i+1)
		}

		if tok.Literal != ti.expectedLiteral {
			t.Errorf(
				"Failed to parse Token Literal correctly. Expected %s Actual %s in test %d",
				ti.expectedLiteral,
				tok.Literal,
				i+1)
		}
	}
}

func TestLexerForQuery(t *testing.T) {
	input := `
	query {
		product(id: "dummy") {
			id
			name
			uri
		}
	}
`

	lex := NewLexer(input)

	tests := []struct {
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
		tok := lex.NextToken()
		if tok.Type != ti.expectedType {
			t.Errorf("Failed to parse Token Type correctly. Expected %s Actual %s in test %d",
				ti.expectedType,
				tok.Type,
				i+1)
		}

		if tok.Literal != ti.expectedLiteral {
			t.Errorf(
				"Failed to parse Token Literal correctly. Expected %s Actual %s in test %d",
				ti.expectedLiteral,
				tok.Literal,
				i+1)
		}
	}
}

func TestLexerForMutation(t *testing.T) {
	input := `
	mutation {
		addProduct(input: "dummy") {
			id
			name
		}
	}
`

	lex := NewLexer(input)

	tests := []struct {
		expectedLiteral string
		expectedType    TokenType
	}{
		{"mutation", MUTATION},
		{"{", LBRACE},
		{"addProduct", IDENTIFIER},
		{"(", LPAREN},
		{"input", INPUT},
		{":", COLON},
		{"dummy", STRING},
		{")", RPAREN},
		{"{", LBRACE},
		{"id", IDENTIFIER},
		{"name", IDENTIFIER},
		{"}", RBRACE},
		{"}", RBRACE},
	}

	for i, ti := range tests {
		tok := lex.NextToken()
		if tok.Type != ti.expectedType {
			t.Errorf("Failed to parse Token Type correctly. Expected %s Actual %s in test %d",
				ti.expectedType,
				tok.Type,
				i+1)
		}

		if tok.Literal != ti.expectedLiteral {
			t.Errorf(
				"Failed to parse Token Literal correctly. Expected %s Actual %s in test %d",
				ti.expectedLiteral,
				tok.Literal,
				i+1)
		}
	}
}
