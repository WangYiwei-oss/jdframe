// Code generated from C:/Users/WangYiWei/Desktop/jdframe/src/exprengine/antlr\ConfigExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 158,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 64, 10, 10, 13, 10, 14, 10, 65, 3,
	10, 3, 10, 3, 11, 3, 11, 7, 11, 72, 10, 11, 12, 11, 14, 11, 75, 11, 11,
	3, 11, 5, 11, 78, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 95,
	10, 14, 3, 15, 3, 15, 3, 16, 6, 16, 100, 10, 16, 13, 16, 14, 16, 101, 5,
	16, 104, 10, 16, 3, 16, 3, 16, 6, 16, 108, 10, 16, 13, 16, 14, 16, 109,
	3, 17, 5, 17, 113, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 5, 18, 119, 10,
	18, 3, 18, 3, 18, 7, 18, 123, 10, 18, 12, 18, 14, 18, 126, 11, 18, 5, 18,
	128, 10, 18, 3, 19, 3, 19, 3, 19, 7, 19, 133, 10, 19, 12, 19, 14, 19, 136,
	11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 6, 20, 144, 10, 20, 13,
	20, 14, 20, 145, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 5, 21, 157, 10, 21, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 2, 25, 2, 27, 2, 29, 2, 31, 2, 33,
	13, 35, 14, 37, 15, 39, 16, 41, 17, 3, 2, 10, 5, 2, 11, 12, 15, 15, 34,
	34, 4, 2, 12, 12, 15, 15, 5, 2, 50, 59, 67, 72, 99, 104, 9, 2, 36, 36,
	94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 3, 2, 50, 59,
	3, 2, 51, 59, 4, 2, 36, 36, 94, 94, 3, 2, 67, 92, 2, 169, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 45,
	3, 2, 2, 2, 7, 47, 3, 2, 2, 2, 9, 49, 3, 2, 2, 2, 11, 51, 3, 2, 2, 2, 13,
	53, 3, 2, 2, 2, 15, 55, 3, 2, 2, 2, 17, 57, 3, 2, 2, 2, 19, 63, 3, 2, 2,
	2, 21, 69, 3, 2, 2, 2, 23, 83, 3, 2, 2, 2, 25, 89, 3, 2, 2, 2, 27, 91,
	3, 2, 2, 2, 29, 96, 3, 2, 2, 2, 31, 103, 3, 2, 2, 2, 33, 112, 3, 2, 2,
	2, 35, 127, 3, 2, 2, 2, 37, 129, 3, 2, 2, 2, 39, 139, 3, 2, 2, 2, 41, 156,
	3, 2, 2, 2, 43, 44, 7, 63, 2, 2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 125, 2,
	2, 46, 6, 3, 2, 2, 2, 47, 48, 7, 46, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50, 7,
	127, 2, 2, 50, 10, 3, 2, 2, 2, 51, 52, 7, 93, 2, 2, 52, 12, 3, 2, 2, 2,
	53, 54, 7, 95, 2, 2, 54, 14, 3, 2, 2, 2, 55, 56, 7, 60, 2, 2, 56, 16, 3,
	2, 2, 2, 57, 58, 7, 112, 2, 2, 58, 59, 7, 119, 2, 2, 59, 60, 7, 110, 2,
	2, 60, 61, 7, 110, 2, 2, 61, 18, 3, 2, 2, 2, 62, 64, 9, 2, 2, 2, 63, 62,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 68, 8, 10, 2, 2, 68, 20, 3, 2, 2, 2, 69, 73, 7,
	37, 2, 2, 70, 72, 10, 3, 2, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2,
	73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3,
	2, 2, 2, 76, 78, 7, 15, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78,
	79, 3, 2, 2, 2, 79, 80, 7, 12, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 8, 11,
	2, 2, 82, 22, 3, 2, 2, 2, 83, 84, 7, 119, 2, 2, 84, 85, 5, 25, 13, 2, 85,
	86, 5, 25, 13, 2, 86, 87, 5, 25, 13, 2, 87, 88, 5, 25, 13, 2, 88, 24, 3,
	2, 2, 2, 89, 90, 9, 4, 2, 2, 90, 26, 3, 2, 2, 2, 91, 94, 7, 94, 2, 2, 92,
	95, 9, 5, 2, 2, 93, 95, 5, 23, 12, 2, 94, 92, 3, 2, 2, 2, 94, 93, 3, 2,
	2, 2, 95, 28, 3, 2, 2, 2, 96, 97, 9, 6, 2, 2, 97, 30, 3, 2, 2, 2, 98, 100,
	5, 29, 15, 2, 99, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 99, 3, 2,
	2, 2, 101, 102, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2,
	103, 104, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 7, 48, 2, 2, 106,
	108, 5, 29, 15, 2, 107, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 107,
	3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 32, 3, 2, 2, 2, 111, 113, 7, 47,
	2, 2, 112, 111, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2,
	114, 115, 5, 31, 16, 2, 115, 34, 3, 2, 2, 2, 116, 128, 7, 50, 2, 2, 117,
	119, 7, 47, 2, 2, 118, 117, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120,
	3, 2, 2, 2, 120, 124, 9, 7, 2, 2, 121, 123, 5, 29, 15, 2, 122, 121, 3,
	2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2,
	2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 116, 3, 2, 2, 2, 127,
	118, 3, 2, 2, 2, 128, 36, 3, 2, 2, 2, 129, 134, 7, 36, 2, 2, 130, 133,
	5, 27, 14, 2, 131, 133, 10, 8, 2, 2, 132, 130, 3, 2, 2, 2, 132, 131, 3,
	2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2,
	2, 135, 137, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 138, 7, 36, 2, 2, 138,
	38, 3, 2, 2, 2, 139, 143, 9, 9, 2, 2, 140, 144, 9, 9, 2, 2, 141, 144, 5,
	29, 15, 2, 142, 144, 7, 97, 2, 2, 143, 140, 3, 2, 2, 2, 143, 141, 3, 2,
	2, 2, 143, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2,
	145, 146, 3, 2, 2, 2, 146, 40, 3, 2, 2, 2, 147, 148, 7, 118, 2, 2, 148,
	149, 7, 116, 2, 2, 149, 150, 7, 119, 2, 2, 150, 157, 7, 103, 2, 2, 151,
	152, 7, 104, 2, 2, 152, 153, 7, 99, 2, 2, 153, 154, 7, 110, 2, 2, 154,
	155, 7, 117, 2, 2, 155, 157, 7, 103, 2, 2, 156, 147, 3, 2, 2, 2, 156, 151,
	3, 2, 2, 2, 157, 42, 3, 2, 2, 2, 19, 2, 65, 73, 77, 94, 101, 103, 109,
	112, 118, 124, 127, 132, 134, 143, 145, 156, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'{'", "','", "'}'", "'['", "']'", "':'", "'null'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "SPACE", "COMMENT", "FloatArg", "IntArg",
	"StringArg", "KEY", "Bool",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "SPACE",
	"COMMENT", "UNICODE", "HEX", "ESC", "DIGIT", "Float", "FloatArg", "IntArg",
	"StringArg", "KEY", "Bool",
}

type ConfigExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewConfigExprLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ConfigExprLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewConfigExprLexer(input antlr.CharStream) *ConfigExprLexer {
	l := new(ConfigExprLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ConfigExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ConfigExprLexer tokens.
const (
	ConfigExprLexerT__0      = 1
	ConfigExprLexerT__1      = 2
	ConfigExprLexerT__2      = 3
	ConfigExprLexerT__3      = 4
	ConfigExprLexerT__4      = 5
	ConfigExprLexerT__5      = 6
	ConfigExprLexerT__6      = 7
	ConfigExprLexerT__7      = 8
	ConfigExprLexerSPACE     = 9
	ConfigExprLexerCOMMENT   = 10
	ConfigExprLexerFloatArg  = 11
	ConfigExprLexerIntArg    = 12
	ConfigExprLexerStringArg = 13
	ConfigExprLexerKEY       = 14
	ConfigExprLexerBool      = 15
)
