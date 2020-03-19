package object

import (
	"bytes"
	"fmt"
	"genert.org/interpreter/ast"
	"hash/fnv"
	"strings"
)

type Type string
type BuiltinFunction func(args ...Object) Object

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct {
	Type  Type
	Value uint64
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

type Object interface {
	Type() Type
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type      { return INTEGER_OBJ }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type      { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

type Null struct{}

func (n *Null) Type() Type      { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type      { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() Type      { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() Type      { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type      { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() Type { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type HashPair struct {
	Key   Object
	Value Object
}
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() Type { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	var pairs []string
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
