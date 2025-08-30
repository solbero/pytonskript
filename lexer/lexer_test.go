package lexer

import (
	"testing"

	"github.com/solbero/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `la five = 5;
la ten = 10;

la add = funksjon(x, y) {
	x + y;
};

la result = add(five, ten);
!-/*5;
5 < 10 > 5;

hvis (5 < 10) {
	returner sant;
} ellers {
	returner falskt;
}

10 == 10;
10 != 9;

"foobar"
"foo bar"

"hello \"world\""
"hello\nworld"
"hello\t\t\tworld"
"hello\\world"
"hello\bworld"

[1, 2];
{"foo": "bar"};

len("123")

`

	tests := []struct {
		expectedType     token.TokenType
		expectedLiterial string
	}{
		{token.LET, "la"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "la"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "la"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "funksjon"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "la"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "hvis"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "returner"},
		{token.TRUE, "sant"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "ellers"},
		{token.LBRACE, "{"},
		{token.RETURN, "returner"},
		{token.FALSE, "falskt"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.STRING, "hello \"world\""},
		{token.STRING, "hello\nworld"},
		{token.STRING, "hello\t\t\tworld"},
		{token.STRING, "hello\\world"},
		{token.STRING, "helloworld"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "len"},
		{token.LPAREN, "("},
		{token.STRING, "123"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong: expected %q got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiterial {
			t.Fatalf("tests[%d] - literal wrong: expected %q got %q", i, tt.expectedLiterial, tok.Literal)
		}
	}
}
