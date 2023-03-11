package lang

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Type int

const (
	VAR Type = iota
	FUNC
	INT
	STRING
	NONE
)

const MAIN = "mehn"

// Holds a return value for a function call. Need this so we can catch and return the value from a panic
type Return struct {
	val any
}

type Metadata struct {
	scope  Scope
	symbol Symbol
}

type Symbol interface {
	Name() string
	Type() Type
}

type FuncSymbol interface {
	Symbol
	Params() []*VarSymbol
	Call(v antlr.ParseTreeVisitor, args []any) Return
}

type VarSymbol struct {
	VarName string
	VarType Type
	Node    antlr.ParseTree
}

func (v *VarSymbol) Name() string {
	return v.VarName
}

func (v *VarSymbol) Type() Type {
	return v.VarType
}

type FuncSymbolBase struct {
	FuncName   string
	FuncType   Type
	FuncParams []*VarSymbol
}

func (f *FuncSymbolBase) Name() string {
	return f.FuncName
}

func (f *FuncSymbolBase) Type() Type {
	return f.FuncType
}

func (f *FuncSymbolBase) Params() []*VarSymbol {
	return f.FuncParams
}

type UserFuncSymbol struct {
	FuncSymbolBase
	Node antlr.ParseTree
}

func (f *UserFuncSymbol) Call(v antlr.ParseTreeVisitor, _ []any) (ret Return) {
	defer CatchReturn(&ret)
	f.Node.Accept(v)
	return Return{}
}

type PrintFunc struct {
	FuncSymbolBase
}

func (f *PrintFunc) Call(_ antlr.ParseTreeVisitor, args []any) (ret Return) {
	fmt.Println(args...)
	return Return{}
}

func CatchReturn(ret *Return) {
	if r := recover(); r != nil {
		switch v := r.(type) {
		case *ReturnError:
			ret.val = v.val
		default:
			panic("continue panic") // TODO: verify what the error logs look like in this case
		}
	}
}

type Scope interface {
	EnclosingScope() Scope
	Name() string
	Define(s Symbol)
	Resolve(name string) (Symbol, error)
}

type DefaultScope struct {
	Enclosing   Scope
	SymbolTable map[string]Symbol
	ScopeName   string
}

func NewDefaultScope(name string, enclosingScope Scope) DefaultScope {
	return DefaultScope{
		ScopeName:   name,
		Enclosing:   enclosingScope,
		SymbolTable: make(map[string]Symbol),
	}
}

func (s *DefaultScope) EnclosingScope() Scope {
	return s.Enclosing
}

func (s *DefaultScope) Name() string {
	return s.ScopeName
}

func (s *DefaultScope) Define(sym Symbol) {
	s.SymbolTable[sym.Name()] = sym
}

func (s *DefaultScope) Resolve(name string) (Symbol, error) {
	sym, ok := s.SymbolTable[name]
	if !ok && s.Enclosing != nil {
		return s.Enclosing.Resolve(name)
	}
	if !ok {
		return nil, fmt.Errorf("unable to resolve symbol [%s]", name)
	}
	return sym, nil
}
