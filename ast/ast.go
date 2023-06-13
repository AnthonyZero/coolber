package ast

import "coolber/token"

type Node interface {
	TokenLiteral() string
}

// Statement All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// Expression All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// LetStatement let语句
type LetStatement struct {
	Token token.Token // the token.LET词法单元
	Name  *Identifier //绑定的标识符
	Value Expression  //产生值的表达式
}

// TokenLiteral 实现Node接口
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// 实现Statement接口
func (ls *LetStatement) statementNode() {
}

// Identifier 标识符
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {

}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
