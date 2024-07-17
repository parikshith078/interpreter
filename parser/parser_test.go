package parser

import (
	"testing"

	"github.com/parikshith078/interpreter/ast"
	"github.com/parikshith078/interpreter/lexer"
)

func TestLetStatments(t *testing.T) {
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
