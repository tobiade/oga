package sem

import "fmt"

type Type int

const (
	VAR Type = iota
	FUNC
	INT
	STRING
	NONE
)

const MAIN = "mehn"

type Metadata struct {
	scope  Scope
	symbol Symbol
}

type Symbol interface {
	Name() string
	Type() Type
}

type VarSymbol struct {
	VarName string
	VarType Type
}

func (v *VarSymbol) Name() string {
	return v.VarName
}

func (v *VarSymbol) Type() Type {
	return v.VarType
}

type FuncSymbol struct {
	FuncName string
	FuncType Type
	Params   []*VarSymbol
}

func (v *FuncSymbol) Name() string {
	return v.FuncName
}

func (v *FuncSymbol) Type() Type {
	return v.FuncType
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
