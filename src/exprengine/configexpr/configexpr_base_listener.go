// Code generated from C:/Users/WangYiWei/Desktop/jdframe/src/exprengine/antlr\ConfigExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ConfigExpr

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// BaseConfigExprListener is a complete listener for a parse tree produced by ConfigExprParser.
type BaseConfigExprListener struct{}

var _ ConfigExprListener = &BaseConfigExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConfigExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConfigExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConfigExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConfigExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChunk is called when production chunk is entered.
func (s *BaseConfigExprListener) EnterChunk(ctx *ChunkContext) {}

// ExitChunk is called when production chunk is exited.
func (s *BaseConfigExprListener) ExitChunk(ctx *ChunkContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseConfigExprListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseConfigExprListener) ExitBlock(ctx *BlockContext) {}

// EnterKvpair is called when production kvpair is entered.
func (s *BaseConfigExprListener) EnterKvpair(ctx *KvpairContext) {}

// ExitKvpair is called when production kvpair is exited.
func (s *BaseConfigExprListener) ExitKvpair(ctx *KvpairContext) {}

// EnterFuncKey is called when production FuncKey is entered.
func (s *BaseConfigExprListener) EnterFuncKey(ctx *FuncKeyContext) {}

// ExitFuncKey is called when production FuncKey is exited.
func (s *BaseConfigExprListener) ExitFuncKey(ctx *FuncKeyContext) {}

// EnterFuncValue is called when production FuncValue is entered.
func (s *BaseConfigExprListener) EnterFuncValue(ctx *FuncValueContext) {}

// ExitFuncValue is called when production FuncValue is exited.
func (s *BaseConfigExprListener) ExitFuncValue(ctx *FuncValueContext) {}
