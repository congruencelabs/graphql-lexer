# graphql-lexer
A minimal Graphql lexer written in Go. The lexer tokenizes below Graphql syntaxes:
- Query
- Mutation
- Subscription
- Schema Definition Language(SDL)

[![Build Status](https://travis-ci.org/congruencelabs/graphql-lexer.svg?branch=master)](https://travis-ci.org/congruencelabs/graphql-lexer)

## Prerequisite
```
go 1.12
```

# Usage
1. Install the go module
```
go get github.com/congruencelabs/graphql-lexer
```
2. Import the `graphql-lexer` package
```go
import (
    lexer "github.com/congruencelabs/graphql-lexer"
)
```
3. use the `lexer` to get the tokens or _lexemes_ for a given input

```go
input := `query {
    product(id: "dummy") {
        id
        name
        uri
    }
}`

lex := lexer.NewLexer(input)

for {
    tok := lex.NextToken()
    if tok.Type == token.EOF {
        return
    }

    fmt.Printf("TokenType %v TokenValue %s \n", tok.Type, tok.Literal)
}
```

4. This will print to stdout the below
```
TokenType query TokenValue query 
TokenType { TokenValue { 
TokenType IDENTIFIER TokenValue product 
TokenType ( TokenValue ( 
TokenType IDENTIFIER TokenValue id 
TokenType : TokenValue : 
TokenType String TokenValue dummy 
TokenType ) TokenValue ) 
TokenType { TokenValue { 
TokenType IDENTIFIER TokenValue id 
TokenType IDENTIFIER TokenValue name 
TokenType IDENTIFIER TokenValue uri 
TokenType } TokenValue } 
TokenType } TokenValue }
```
