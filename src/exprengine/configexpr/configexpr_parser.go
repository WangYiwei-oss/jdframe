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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 81, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 25, 10,
	3, 12, 3, 14, 3, 28, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 7, 6, 40, 10, 6, 12, 6, 14, 6, 43, 11, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 5, 6, 49, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 55, 10, 7, 12,
	7, 14, 7, 58, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 64, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 79, 10, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	2, 2, 2, 82, 2, 20, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 29, 3, 2, 2, 2, 8,
	33, 3, 2, 2, 2, 10, 48, 3, 2, 2, 2, 12, 63, 3, 2, 2, 2, 14, 65, 3, 2, 2,
	2, 16, 69, 3, 2, 2, 2, 18, 78, 3, 2, 2, 2, 20, 21, 5, 4, 3, 2, 21, 22,
	7, 2, 2, 3, 22, 3, 3, 2, 2, 2, 23, 25, 5, 6, 4, 2, 24, 23, 3, 2, 2, 2,
	25, 28, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 5, 3, 2,
	2, 2, 28, 26, 3, 2, 2, 2, 29, 30, 5, 8, 5, 2, 30, 31, 7, 3, 2, 2, 31, 32,
	5, 18, 10, 2, 32, 7, 3, 2, 2, 2, 33, 34, 7, 16, 2, 2, 34, 9, 3, 2, 2, 2,
	35, 36, 7, 4, 2, 2, 36, 41, 5, 14, 8, 2, 37, 38, 7, 5, 2, 2, 38, 40, 5,
	14, 8, 2, 39, 37, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41,
	42, 3, 2, 2, 2, 42, 44, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 45, 7, 6, 2,
	2, 45, 49, 3, 2, 2, 2, 46, 47, 7, 4, 2, 2, 47, 49, 7, 6, 2, 2, 48, 35,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 11, 3, 2, 2, 2, 50, 51, 7, 7, 2, 2,
	51, 56, 5, 18, 10, 2, 52, 53, 7, 5, 2, 2, 53, 55, 5, 18, 10, 2, 54, 52,
	3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2,
	57, 59, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 7, 8, 2, 2, 60, 64, 3,
	2, 2, 2, 61, 62, 7, 7, 2, 2, 62, 64, 7, 8, 2, 2, 63, 50, 3, 2, 2, 2, 63,
	61, 3, 2, 2, 2, 64, 13, 3, 2, 2, 2, 65, 66, 5, 16, 9, 2, 66, 67, 7, 9,
	2, 2, 67, 68, 5, 18, 10, 2, 68, 15, 3, 2, 2, 2, 69, 70, 7, 15, 2, 2, 70,
	17, 3, 2, 2, 2, 71, 79, 7, 15, 2, 2, 72, 79, 7, 14, 2, 2, 73, 79, 7, 13,
	2, 2, 74, 79, 7, 17, 2, 2, 75, 79, 7, 10, 2, 2, 76, 79, 5, 10, 6, 2, 77,
	79, 5, 12, 7, 2, 78, 71, 3, 2, 2, 2, 78, 72, 3, 2, 2, 2, 78, 73, 3, 2,
	2, 2, 78, 74, 3, 2, 2, 2, 78, 75, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 77,
	3, 2, 2, 2, 79, 19, 3, 2, 2, 2, 8, 26, 41, 48, 56, 63, 78,
}
var literalNames = []string{
	"", "'='", "'{'", "','", "'}'", "'['", "']'", "':'", "'null'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "SPACE", "COMMENT", "FloatArg", "IntArg",
	"StringArg", "KEY", "Bool",
}

var ruleNames = []string{
	"config", "block", "kvpair", "key", "object", "slice", "pair", "mapkey",
	"value",
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
	ConfigExprParserT__0      = 1
	ConfigExprParserT__1      = 2
	ConfigExprParserT__2      = 3
	ConfigExprParserT__3      = 4
	ConfigExprParserT__4      = 5
	ConfigExprParserT__5      = 6
	ConfigExprParserT__6      = 7
	ConfigExprParserT__7      = 8
	ConfigExprParserSPACE     = 9
	ConfigExprParserCOMMENT   = 10
	ConfigExprParserFloatArg  = 11
	ConfigExprParserIntArg    = 12
	ConfigExprParserStringArg = 13
	ConfigExprParserKEY       = 14
	ConfigExprParserBool      = 15
)

// ConfigExprParser rules.
const (
	ConfigExprParserRULE_config = 0
	ConfigExprParserRULE_block  = 1
	ConfigExprParserRULE_kvpair = 2
	ConfigExprParserRULE_key    = 3
	ConfigExprParserRULE_object = 4
	ConfigExprParserRULE_slice  = 5
	ConfigExprParserRULE_pair   = 6
	ConfigExprParserRULE_mapkey = 7
	ConfigExprParserRULE_value  = 8
)

// IConfigContext is an interface to support dynamic dispatch.
type IConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigContext differentiates from other interfaces.
	IsConfigContext()
}

type ConfigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigContext() *ConfigContext {
	var p = new(ConfigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_config
	return p
}

func (*ConfigContext) IsConfigContext() {}

func NewConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigContext {
	var p = new(ConfigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_config

	return p
}

func (s *ConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConfigContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserEOF, 0)
}

func (s *ConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterConfig(s)
	}
}

func (s *ConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitConfig(s)
	}
}

func (p *ConfigExprParser) Config() (localctx IConfigContext) {
	localctx = NewConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfigExprParserRULE_config)

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
		p.SetState(18)
		p.Block()
	}
	{
		p.SetState(19)
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
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConfigExprParserKEY {
		{
			p.SetState(21)
			p.Kvpair()
		}

		p.SetState(26)
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
		p.SetState(27)
		p.Key()
	}
	{
		p.SetState(28)
		p.Match(ConfigExprParserT__0)
	}
	{
		p.SetState(29)
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
		p.SetState(31)
		p.Match(ConfigExprParserKEY)
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjectContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *ConfigExprParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConfigExprParserRULE_object)
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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(ConfigExprParserT__1)
		}
		{
			p.SetState(34)
			p.Pair()
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConfigExprParserT__2 {
			{
				p.SetState(35)
				p.Match(ConfigExprParserT__2)
			}
			{
				p.SetState(36)
				p.Pair()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
			p.Match(ConfigExprParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(ConfigExprParserT__1)
		}
		{
			p.SetState(45)
			p.Match(ConfigExprParserT__3)
		}

	}

	return localctx
}

// ISliceContext is an interface to support dynamic dispatch.
type ISliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceContext differentiates from other interfaces.
	IsSliceContext()
}

type SliceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceContext() *SliceContext {
	var p = new(SliceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_slice
	return p
}

func (*SliceContext) IsSliceContext() {}

func NewSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceContext {
	var p = new(SliceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_slice

	return p
}

func (s *SliceContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *SliceContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitSlice(s)
	}
}

func (p *ConfigExprParser) Slice() (localctx ISliceContext) {
	localctx = NewSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConfigExprParserRULE_slice)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(ConfigExprParserT__4)
		}
		{
			p.SetState(49)
			p.Value()
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ConfigExprParserT__2 {
			{
				p.SetState(50)
				p.Match(ConfigExprParserT__2)
			}
			{
				p.SetState(51)
				p.Value()
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(57)
			p.Match(ConfigExprParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(ConfigExprParserT__4)
		}
		{
			p.SetState(60)
			p.Match(ConfigExprParserT__5)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) Mapkey() IMapkeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapkeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapkeyContext)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *ConfigExprParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConfigExprParserRULE_pair)

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
		p.SetState(63)
		p.Mapkey()
	}
	{
		p.SetState(64)
		p.Match(ConfigExprParserT__6)
	}
	{
		p.SetState(65)
		p.Value()
	}

	return localctx
}

// IMapkeyContext is an interface to support dynamic dispatch.
type IMapkeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapkeyContext differentiates from other interfaces.
	IsMapkeyContext()
}

type MapkeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapkeyContext() *MapkeyContext {
	var p = new(MapkeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigExprParserRULE_mapkey
	return p
}

func (*MapkeyContext) IsMapkeyContext() {}

func NewMapkeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapkeyContext {
	var p = new(MapkeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigExprParserRULE_mapkey

	return p
}

func (s *MapkeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MapkeyContext) CopyFrom(ctx *MapkeyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MapkeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapkeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MapKeyContext struct {
	*MapkeyContext
}

func NewMapKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapKeyContext {
	var p = new(MapKeyContext)

	p.MapkeyContext = NewEmptyMapkeyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MapkeyContext))

	return p
}

func (s *MapKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapKeyContext) StringArg() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserStringArg, 0)
}

func (s *MapKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterMapKey(s)
	}
}

func (s *MapKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitMapKey(s)
	}
}

func (p *ConfigExprParser) Mapkey() (localctx IMapkeyContext) {
	localctx = NewMapkeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConfigExprParserRULE_mapkey)

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

	localctx = NewMapKeyContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(ConfigExprParserStringArg)
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

type NullValueContext struct {
	*ValueContext
}

func NewNullValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValueContext {
	var p = new(NullValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitNullValue(s)
	}
}

type ObjectValueContext struct {
	*ValueContext
}

func NewObjectValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

type BoolValueContext struct {
	*ValueContext
}

func NewBoolValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolValueContext {
	var p = new(BoolValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) Bool() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserBool, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitBoolValue(s)
	}
}

type SliceValueContext struct {
	*ValueContext
}

func NewSliceValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceValueContext {
	var p = new(SliceValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *SliceValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceValueContext) Slice() ISliceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
}

func (s *SliceValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterSliceValue(s)
	}
}

func (s *SliceValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitSliceValue(s)
	}
}

type FloatValueContext struct {
	*ValueContext
}

func NewFloatValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatValueContext {
	var p = new(FloatValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) FloatArg() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserFloatArg, 0)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

type StringValueContext struct {
	*ValueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) StringArg() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserStringArg, 0)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitStringValue(s)
	}
}

type IntValueContext struct {
	*ValueContext
}

func NewIntValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntValueContext {
	var p = new(IntValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *IntValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValueContext) IntArg() antlr.TerminalNode {
	return s.GetToken(ConfigExprParserIntArg, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigExprListener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (p *ConfigExprParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ConfigExprParserRULE_value)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ConfigExprParserStringArg:
		localctx = NewStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(ConfigExprParserStringArg)
		}

	case ConfigExprParserIntArg:
		localctx = NewIntValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(ConfigExprParserIntArg)
		}

	case ConfigExprParserFloatArg:
		localctx = NewFloatValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.Match(ConfigExprParserFloatArg)
		}

	case ConfigExprParserBool:
		localctx = NewBoolValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)
			p.Match(ConfigExprParserBool)
		}

	case ConfigExprParserT__7:
		localctx = NewNullValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)
			p.Match(ConfigExprParserT__7)
		}

	case ConfigExprParserT__1:
		localctx = NewObjectValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(74)
			p.Object()
		}

	case ConfigExprParserT__4:
		localctx = NewSliceValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(75)
			p.Slice()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
