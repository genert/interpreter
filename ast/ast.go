package ast

import (
	"bytes"
	"genert.org/interpreter/token"
)

type Node interface {
	// TokenLiteral will be used only for debugging and testing.
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

// It is important to keep in mind that expressions produce values,
// statements don't. let x = 5 doesn't produce a value,
// whereas 5 does (the value it produces is 5).
// A return 5; statement doesn't produce a value, but add(5,5) does.
type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the root node of every AST our parser produces.
type Program struct {
	Statements []Statement
}

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

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
