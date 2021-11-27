// Code generated from C:/Users/WangYiWei/Desktop/jdframe/src/exprengine/antlr\ConfigExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ConfigExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

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

// EnterConfig is called when production config is entered.
func (s *BaseConfigExprListener) EnterConfig(ctx *ConfigContext) {}

// ExitConfig is called when production config is exited.
func (s *BaseConfigExprListener) ExitConfig(ctx *ConfigContext) {}

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

// EnterObject is called when production object is entered.
func (s *BaseConfigExprListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseConfigExprListener) ExitObject(ctx *ObjectContext) {}

// EnterSlice is called when production slice is entered.
func (s *BaseConfigExprListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production slice is exited.
func (s *BaseConfigExprListener) ExitSlice(ctx *SliceContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseConfigExprListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseConfigExprListener) ExitPair(ctx *PairContext) {}

// EnterMapKey is called when production MapKey is entered.
func (s *BaseConfigExprListener) EnterMapKey(ctx *MapKeyContext) {}

// ExitMapKey is called when production MapKey is exited.
func (s *BaseConfigExprListener) ExitMapKey(ctx *MapKeyContext) {}

// EnterStringValue is called when production StringValue is entered.
func (s *BaseConfigExprListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production StringValue is exited.
func (s *BaseConfigExprListener) ExitStringValue(ctx *StringValueContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseConfigExprListener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseConfigExprListener) ExitIntValue(ctx *IntValueContext) {}

// EnterFloatValue is called when production FloatValue is entered.
func (s *BaseConfigExprListener) EnterFloatValue(ctx *FloatValueContext) {}

// ExitFloatValue is called when production FloatValue is exited.
func (s *BaseConfigExprListener) ExitFloatValue(ctx *FloatValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseConfigExprListener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseConfigExprListener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterNullValue is called when production NullValue is entered.
func (s *BaseConfigExprListener) EnterNullValue(ctx *NullValueContext) {}

// ExitNullValue is called when production NullValue is exited.
func (s *BaseConfigExprListener) ExitNullValue(ctx *NullValueContext) {}

// EnterObjectValue is called when production ObjectValue is entered.
func (s *BaseConfigExprListener) EnterObjectValue(ctx *ObjectValueContext) {}

// ExitObjectValue is called when production ObjectValue is exited.
func (s *BaseConfigExprListener) ExitObjectValue(ctx *ObjectValueContext) {}

// EnterSliceValue is called when production SliceValue is entered.
func (s *BaseConfigExprListener) EnterSliceValue(ctx *SliceValueContext) {}

// ExitSliceValue is called when production SliceValue is exited.
func (s *BaseConfigExprListener) ExitSliceValue(ctx *SliceValueContext) {}
