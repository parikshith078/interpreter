package ast

import "github.com/parikshith078/interpreter/token"

type Node interface {
	TokenLiteral() string
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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatment struct {
	Token token.Token
	Name  *Indentifier
	Value Expression
}

func (ls *LetStatment) statementNode() {}
func (ls *LetStatment) TokenLiteral() string { return ls.Token.Literal }

type Indentifier struct {
	Token token.Token // the token.IDENT token
  Value string
}

func (i *Indentifier) expressionNode() {} 

func (i *Indentifier) TokenLiteral() string { return i.Token.Literal }
