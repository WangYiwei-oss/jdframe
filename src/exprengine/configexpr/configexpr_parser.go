// Code generated from C:/Users/WangYiWei/Desktop/jdframe/src/exprengine/antlr\ConfigExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ConfigExpr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 30, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 3, 7, 3, 17, 10, 3, 12, 3, 14, 3, 20, 11, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 3, 4, 2, 8,
	8, 11, 12, 2, 25, 2, 12, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 21, 3, 2, 2,
	2, 8, 25, 3, 2, 2, 2, 10, 27, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2, 13, 14, 7,
	2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 17, 5, 6, 4, 2, 16, 15, 3, 2, 2, 2, 17,
	20, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 5, 3, 2, 2,
	2, 20, 18, 3, 2, 2, 2, 21, 22, 5, 8, 5, 2, 22, 23, 7, 9, 2, 2, 23, 24,
	5, 10, 6, 2, 24, 7, 3, 2, 2, 2, 25, 26, 7, 13, 2, 2, 26, 9, 3, 2, 2, 2,
	27, 28, 9, 2, 2, 2, 28, 11, 3, 2, 2, 2, 3, 18,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "'='",
}
var symbolicNames = []string{
	"", "SPACE", "COMMENT", "Float", "FloatArg", "IntArg", "StringArg", "ASSIGN",
	"ARG", "SLICE", "DICT", "KEY",
}

var ruleNames = []string{
	"chunk", "block", "kvpair", "key", "value",
}

type ConfigExprParser struct {
	*antlr.BaseParser
}

// NewConfigExprParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ConfigExprParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewConfigExprParser(input antlr.TokenStream) *ConfigExprParser {
	this := new(ConfigExprParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ConfigExpr.g4"

	return this
}

// ConfigExprParser tokens.
const (
	ConfigExprParserEOF       = antlr.TokenEOF
	ConfigExprParserSPACE     = 1
	ConfigExprParserCOMMENT   = 2
	ConfigExprParserFloat     = 3
	ConfigExprParserFloatArg  = 4
	ConfigExprParserIntArg    = 5
	ConfigExprParserStringArg = 6
	ConfigExprParserASSIGN    = 7
	ConfigExprParserARG       = 8
	ConfigExprParserSLICE     = 9
	ConfigExprParserDICT      = 10
	ConfigExprParserKEY       = 11
)

// ConfigExprParser rules.
const (
	ConfigExprParserRULE_chunk  = 0
	ConfigExprParserRULE_block  = 1
	ConfigExprParserRULE_kvpair = 2
	ConfigExprParserRULE_key    = 3
	ConfigExprParserRULE_value  = 4
)

// IChunkContext is an interface to support dynamic dispatch.
type IChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChunkContext differentiates from other interfaces.
	IsChunkContext()
}

type ChunkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChunkContext() *ChunkContext {
	var p = new(ChunkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_chunk
	return p
}

func (*ChunkContext) IsChunkContext() {}

func NewChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChunkContext {
	var p = new(ChunkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_chunk

	return p
}

func (s *ChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *ChunkContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ChunkContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserEOF, 0)
}

func (s *ChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChunkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterChunk(s)
	}
}

func (s *ChunkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitChunk(s)
	}
}

func (p *ConfigExprParser) Chunk() (localctx IChunkContext) {
	localctx = NewChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfigExprParserRULE_chunk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Block()
	}
	{
		p.SetState(11)
		p.Match(ConfigExprParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllKvpair() []IKvpairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKvpairContext)(nil)).Elem())
	var tst = make([]IKvpairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKvpairContext)
		}
	}

	return tst
}

func (s *BlockContext) Kvpair(i int) IKvpairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvpairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKvpairContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *ConfigExprParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfigExprParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConfigExprParserKEY {
		{
			p.SetState(13)
			p.Kvpair()
		}

		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKvpairContext is an interface to support dynamic dispatch.
type IKvpairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKvpairContext differentiates from other interfaces.
	IsKvpairContext()
}

type KvpairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKvpairContext() *KvpairContext {
	var p = new(KvpairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_kvpair
	return p
}

func (*KvpairContext) IsKvpairContext() {}

func NewKvpairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvpairContext {
	var p = new(KvpairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_kvpair

	return p
}

func (s *KvpairContext) GetParser() antlr.Parser { return s.parser }

func (s *KvpairContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KvpairContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserASSIGN, 0)
}

func (s *KvpairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *KvpairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvpairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvpairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterKvpair(s)
	}
}

func (s *KvpairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitKvpair(s)
	}
}

func (p *ConfigExprParser) Kvpair() (localctx IKvpairContext) {
	localctx = NewKvpairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConfigExprParserRULE_kvpair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Key()
	}
	{
		p.SetState(20)
		p.Match(ConfigExprParserASSIGN)
	}
	{
		p.SetState(21)
		p.Value()
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) CopyFrom(ctx *KeyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncKeyContext struct {
	*KeyContext
}

func NewFuncKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncKeyContext {
	var p = new(FuncKeyContext)

	p.KeyContext = NewEmptyKeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*KeyContext))

	return p
}

func (s *FuncKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncKeyContext) KEY() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserKEY, 0)
}

func (s *FuncKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterFuncKey(s)
	}
}

func (s *FuncKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitFuncKey(s)
	}
}

func (p *ConfigExprParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConfigExprParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFuncKeyContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(ConfigExprParserKEY)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncValueContext struct {
	*ValueContext
}

func NewFuncValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncValueContext {
	var p = new(FuncValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FuncValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncValueContext) StringArg() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserStringArg, 0)
}

func (s *FuncValueContext) SLICE() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserSLICE, 0)
}

func (s *FuncValueContext) DICT() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserDICT, 0)
}

func (s *FuncValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterFuncValue(s)
	}
}

func (s *FuncValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitFuncValue(s)
	}
}

func (p *ConfigExprParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConfigExprParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFuncValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigExprParserStringArg)|(1<<ConfigExprParserSLICE)|(1<<ConfigExprParserDICT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
