package ast

import (
	"bytes"

	"github.com/parikshith078/interpreter/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// Returns head token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatment struct {
	Token token.Token
	Name  *Indentifier
	Value Expression
}

func (ls *LetStatment) statementNode()       {}
func (ls *LetStatment) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatment) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ReturnStatment struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatment) statementNode()       {}
func (rs *ReturnStatment) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatment) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	return out.String()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type ExpressionStatment struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatment) statementNode()       {}
func (es *ExpressionStatment) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatment) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Indentifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Indentifier) expressionNode()      {}
func (i *Indentifier) TokenLiteral() string { return i.Token.Literal }
func (i *Indentifier) String() string       { return i.Value }
