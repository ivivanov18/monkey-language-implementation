package parser

import (
	"testing",
	"go-monkey/ast",
	"go-monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 1983;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	
	if program == nil {
		t.Fatalf("ParsedProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program statements does not contain 3 statements. got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	} {
		{"x"},
		{"y"},
		{"foobar"}
	}

	for i, tt := range tests {
		stmt := 
	}


}