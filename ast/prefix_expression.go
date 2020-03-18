package ast

import (
	"bytes"
	"genert.org/interpreter/token"
)

// PrefixExpression node has two noteworthy fields:
// Operator and Right.
// Operator is a string that’s going to contain either "-" or "!".
// The Right field contains the expression to the right of the operator.
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	// Add deliberately parentheses around the operator and its operand,
	// the expression in Right.
	// That allows us to see which operands belong to which operator.
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}
