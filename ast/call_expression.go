package ast

import (
	"bytes"
	"genert.org/interpreter/token"
	"strings"
)

// Parse the calling of a function: call expressions. Here is their structure:
//   <expression>(<comma separated expressions>)
// Some examples:
//   add(2, 3)
//   add(2 + 2, 3 * 3 * 3)
//   callsFunction(2, 3, function(x, y) { x + y; });

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	var args []string

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}
