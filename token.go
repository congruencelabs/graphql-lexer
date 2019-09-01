package graphql_lexer

const 	(
	EOF          = "EOF"
	ILLEGAL      = "ILLEGAL"
	IDENTIFIER   = "IDENTIFIER"
	COLON        = ":"
	COMMA        = ","
	LPAREN       = "("
	RPAREN       = ")"
	LBRACE       = "{"
	RBRACE       = "}"
	LBRACKET     = "["
	RBRACKET     = "]"
	EQUALS       = "="
	BANG         = "!"
	DOLLAR       = "$"
	SPREAD       = "..."
	AT           = "@"
	PIPE         = "|"
	INT          = "Int"
	STRING       = "String"
	FLOAT        = "Float"
	BLOCKSTRING  = "BlockString"

	DIRECTIVE    = "directive"
	ENUM         = "enum"
	EXTEND       = "extend"
	FRAGMENT     = "fragment"
	INPUT        = "input"
	INTERFACE    = "interface"
	QUERY        = "query"
	MUTATION     = "mutation"
	SCALAR       = "scalar"
	SCHEMA       = "schema"
	SUBSCRIPTION = "subscription"
	TYPE         = "type"
	UNION        = "union"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var tokenKeywordsMap = map[string]TokenType{
	"...": SPREAD,
	"enum": ENUM,
	"extend": EXTEND,
	"fragment": FRAGMENT,
	"input": INPUT,
	"interface": INTERFACE,
	"query": QUERY,
	"mutation": MUTATION,
	"scalar": SCALAR,
	"schema": SCHEMA,
	"subscription": SUBSCRIPTION,
	"type": TYPE,
	"union": UNION,

	"Query": QUERY,
	"Int": INT,
	"String": STRING,
	"Float": FLOAT,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := tokenKeywordsMap[ident]; ok {
		return tok
	}
	return IDENTIFIER
}


