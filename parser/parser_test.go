package parser

import (
	"testing"

	"github.com/parikshith078/interpreter/ast"
	"github.com/parikshith078/interpreter/lexer"
)

func testLetStatments(t *testing.T) {
	input := `
  let x = 5;
  let y = 10;
  let foobar = 234;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 Statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatment(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testReturnStatements(t *testing.T) {
	input := `
  return 5;
  return 10;
  return 24218;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 Statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatment)

		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough Statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatment)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatment. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Indentifier)
	if !ok {
		t.Fatalf("exp not *ast.Indentifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}

}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough Statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatment)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatment. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.Indentifier. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("ident.Value not %s. got=%d", "foobar", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}

}
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatment(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letSmt, ok := s.(*ast.LetStatment)
	if !ok {
		t.Errorf("s not *ast.LetStatment. got=%T", s)
		return false
	}

	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("letSmt.Name.Value not '%s'. got=%s", name, letSmt.Name.Value)
		return false
	}

	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letSmt.Name)
		return false
	}
	return true
}
