%{
// headers
package dentaku

import "text/scanner"

type Expression interface{}

type Token struct {
    token int
    literal string
}

type NumExpr struct {
    literal string
}

type BinOpExpr struct {
    left, right Expression
    operator rune
}
%}

%union{
    expr Expression
    token Token
}

%type<expr> program
%type<expr> expr
%token<token> NUMBER

%left '+' '-'
%left '*' '/'
%left '^'

%%

program
    : expr
    {
        $$ = $1
        yylex.(*Lexer).Result = $$
    }

expr
    : NUMBER
    {
        $$ = NumExpr{literal: $1.literal}
    }
    | expr '+' expr
    {
        $$ = BinOpExpr{left: $1, right: $3, operator: '+'}
    }
    | expr '-' expr
    {
        $$ = BinOpExpr{left: $1, right: $3, operator: '-'}
    }
    | expr '*' expr
    {
        $$ = BinOpExpr{left: $1, right: $3, operator: '*'}
    }
    | expr '/' expr
    {
        $$ = BinOpExpr{left: $1, right: $3, operator: '/'}
    }
    | expr '^' expr
    {
        $$ = BinOpExpr{left: $1, right: $3, operator: '^'}
    }

%%

type Lexer struct {
    scanner.Scanner
    Result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
    token := int(l.Scan())
    if token == scanner.Int {
        token = NUMBER
    }
    lval.token = Token{token: token, literal: l.TokenText()}
    return token
}

func (l *Lexer) Error(e string) {
    panic(e)
}

func Parse(l *Lexer) int {
    return yyParse(l)
}
