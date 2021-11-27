// Code generated from C:/Users/WangYiWei/Desktop/jdframe/src/exprengine/antlr\ConfigExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ConfigExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConfigExprListener is a complete listener for a parse tree produced by ConfigExprParser.
type ConfigExprListener interface {
	antlr.ParseTreeListener

	// EnterConfig is called when entering the config production.
	EnterConfig(c *ConfigContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterKvpair is called when entering the kvpair production.
	EnterKvpair(c *KvpairContext)

	// EnterFuncKey is called when entering the FuncKey production.
	EnterFuncKey(c *FuncKeyContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterMapKey is called when entering the MapKey production.
	EnterMapKey(c *MapKeyContext)

	// EnterStringValue is called when entering the StringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterFloatValue is called when entering the FloatValue production.
	EnterFloatValue(c *FloatValueContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterNullValue is called when entering the NullValue production.
	EnterNullValue(c *NullValueContext)

	// EnterObjectValue is called when entering the ObjectValue production.
	EnterObjectValue(c *ObjectValueContext)

	// EnterSliceValue is called when entering the SliceValue production.
	EnterSliceValue(c *SliceValueContext)

	// ExitConfig is called when exiting the config production.
	ExitConfig(c *ConfigContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitKvpair is called when exiting the kvpair production.
	ExitKvpair(c *KvpairContext)

	// ExitFuncKey is called when exiting the FuncKey production.
	ExitFuncKey(c *FuncKeyContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitMapKey is called when exiting the MapKey production.
	ExitMapKey(c *MapKeyContext)

	// ExitStringValue is called when exiting the StringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitFloatValue is called when exiting the FloatValue production.
	ExitFloatValue(c *FloatValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitNullValue is called when exiting the NullValue production.
	ExitNullValue(c *NullValueContext)

	// ExitObjectValue is called when exiting the ObjectValue production.
	ExitObjectValue(c *ObjectValueContext)

	// ExitSliceValue is called when exiting the SliceValue production.
	ExitSliceValue(c *SliceValueContext)
}
