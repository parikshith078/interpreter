package parser

import (
	"fmt"
	"strconv"

	"github.com/parikshith078/interpreter/ast"
	"github.com/parikshith078/interpreter/lexer"
	"github.com/parikshith078/interpreter/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGRATER
	SUM
	PRODUCT
	PREFIX // -x or !x
	CALL   // myFunction(x)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	// Reading two position so that curToken and peekToken both are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Indentifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := new(ast.Program)
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		fmt.Println("stmt: ", stmt)
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatment()
	default:
		return p.parseExpressionStatment()
	}
}

func (p *Parser) parseExpressionStatment() *ast.ExpressionStatment {
	stmt := &ast.ExpressionStatment{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	val, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = val

	return lit
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatment{Token: p.curToken}

	// If peek token is Indentifier, call nextToken which will place curToken to Indentifier token
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Indentifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// skiping till end of the line
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatment() *ast.ReturnStatment {
	stmt := &ast.ReturnStatment{Token: p.curToken}
	p.nextToken()

	// skiping till semicolon or line end
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	return prefix()
}
