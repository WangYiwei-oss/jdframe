// Code generated from C:/Users/WangYiWei/Desktop/jdframe/src/exprengine/antlr\ConfigExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ConfigExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConfigExprListener is a complete listener for a parse tree produced by ConfigExprParser.
type ConfigExprListener interface {
	antlr.ParseTreeListener

	// EnterChunk is called when entering the chunk production.
	EnterChunk(c *ChunkContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterKvpair is called when entering the kvpair production.
	EnterKvpair(c *KvpairContext)

	// EnterFuncKey is called when entering the FuncKey production.
	EnterFuncKey(c *FuncKeyContext)

	// EnterFuncValue is called when entering the FuncValue production.
	EnterFuncValue(c *FuncValueContext)

	// ExitChunk is called when exiting the chunk production.
	ExitChunk(c *ChunkContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitKvpair is called when exiting the kvpair production.
	ExitKvpair(c *KvpairContext)

	// ExitFuncKey is called when exiting the FuncKey production.
	ExitFuncKey(c *FuncKeyContext)

	// ExitFuncValue is called when exiting the FuncValue production.
	ExitFuncValue(c *FuncValueContext)
}
