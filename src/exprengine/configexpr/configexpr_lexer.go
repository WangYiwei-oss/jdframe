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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 156,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 6, 2, 37, 10, 2, 13, 2, 14, 2, 38, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 45,
	10, 3, 12, 3, 14, 3, 48, 11, 3, 3, 3, 5, 3, 51, 10, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 60, 10, 5, 13, 5, 14, 5, 61, 5, 5, 64, 10,
	5, 3, 5, 3, 5, 6, 5, 68, 10, 5, 13, 5, 14, 5, 69, 3, 6, 5, 6, 73, 10, 6,
	3, 6, 3, 6, 3, 7, 5, 7, 78, 10, 7, 3, 7, 6, 7, 81, 10, 7, 13, 7, 14, 7,
	82, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 91, 10, 8, 12, 8, 14, 8,
	94, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 6, 14, 111, 10, 14, 13, 14, 14,
	14, 112, 3, 15, 3, 15, 5, 15, 117, 10, 15, 3, 15, 3, 15, 3, 15, 7, 15,
	122, 10, 15, 12, 15, 14, 15, 125, 11, 15, 3, 15, 5, 15, 128, 10, 15, 5,
	15, 130, 10, 15, 3, 15, 3, 15, 3, 16, 5, 16, 135, 10, 16, 3, 16, 3, 16,
	7, 16, 139, 10, 16, 12, 16, 14, 16, 142, 11, 16, 3, 16, 5, 16, 145, 10,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 153, 10, 17, 13, 17,
	14, 17, 154, 2, 2, 18, 3, 3, 5, 4, 7, 2, 9, 5, 11, 6, 13, 7, 15, 8, 17,
	9, 19, 2, 21, 2, 23, 2, 25, 2, 27, 10, 29, 11, 31, 12, 33, 13, 3, 2, 8,
	5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 3, 2, 50, 59, 4, 2,
	41, 41, 94, 94, 4, 2, 94, 94, 127, 127, 3, 2, 67, 92, 2, 175, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 3, 36, 3, 2, 2, 2,
	5, 42, 3, 2, 2, 2, 7, 56, 3, 2, 2, 2, 9, 63, 3, 2, 2, 2, 11, 72, 3, 2,
	2, 2, 13, 77, 3, 2, 2, 2, 15, 84, 3, 2, 2, 2, 17, 97, 3, 2, 2, 2, 19, 99,
	3, 2, 2, 2, 21, 101, 3, 2, 2, 2, 23, 103, 3, 2, 2, 2, 25, 105, 3, 2, 2,
	2, 27, 110, 3, 2, 2, 2, 29, 114, 3, 2, 2, 2, 31, 134, 3, 2, 2, 2, 33, 148,
	3, 2, 2, 2, 35, 37, 9, 2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2,
	38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 8,
	2, 2, 2, 41, 4, 3, 2, 2, 2, 42, 46, 7, 37, 2, 2, 43, 45, 10, 3, 2, 2, 44,
	43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2,
	2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 51, 7, 15, 2, 2, 50, 49,
	3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 7, 12, 2, 2,
	53, 54, 3, 2, 2, 2, 54, 55, 8, 3, 2, 2, 55, 6, 3, 2, 2, 2, 56, 57, 9, 4,
	2, 2, 57, 8, 3, 2, 2, 2, 58, 60, 5, 7, 4, 2, 59, 58, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2,
	63, 59, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 7,
	48, 2, 2, 66, 68, 5, 7, 4, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69,
	67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 10, 3, 2, 2, 2, 71, 73, 7, 47,
	2, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75,
	5, 9, 5, 2, 75, 12, 3, 2, 2, 2, 76, 78, 7, 47, 2, 2, 77, 76, 3, 2, 2, 2,
	77, 78, 3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 81, 5, 7, 4, 2, 80, 79, 3,
	2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83,
	14, 3, 2, 2, 2, 84, 92, 7, 41, 2, 2, 85, 86, 7, 94, 2, 2, 86, 91, 11, 2,
	2, 2, 87, 88, 7, 41, 2, 2, 88, 91, 7, 41, 2, 2, 89, 91, 10, 5, 2, 2, 90,
	85, 3, 2, 2, 2, 90, 87, 3, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2,
	2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 92,
	3, 2, 2, 2, 95, 96, 7, 41, 2, 2, 96, 16, 3, 2, 2, 2, 97, 98, 7, 63, 2,
	2, 98, 18, 3, 2, 2, 2, 99, 100, 7, 60, 2, 2, 100, 20, 3, 2, 2, 2, 101,
	102, 7, 46, 2, 2, 102, 22, 3, 2, 2, 2, 103, 104, 7, 48, 2, 2, 104, 24,
	3, 2, 2, 2, 105, 106, 7, 97, 2, 2, 106, 26, 3, 2, 2, 2, 107, 111, 5, 15,
	8, 2, 108, 111, 5, 13, 7, 2, 109, 111, 5, 11, 6, 2, 110, 107, 3, 2, 2,
	2, 110, 108, 3, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 28, 3, 2, 2, 2, 114, 129, 7,
	93, 2, 2, 115, 117, 5, 27, 14, 2, 116, 115, 3, 2, 2, 2, 116, 117, 3, 2,
	2, 2, 117, 130, 3, 2, 2, 2, 118, 119, 5, 27, 14, 2, 119, 120, 5, 21, 11,
	2, 120, 122, 3, 2, 2, 2, 121, 118, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123,
	121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123,
	3, 2, 2, 2, 126, 128, 5, 27, 14, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3,
	2, 2, 2, 128, 130, 3, 2, 2, 2, 129, 116, 3, 2, 2, 2, 129, 123, 3, 2, 2,
	2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 95, 2, 2, 132, 30, 3, 2, 2, 2, 133,
	135, 7, 94, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136,
	3, 2, 2, 2, 136, 140, 7, 125, 2, 2, 137, 139, 10, 6, 2, 2, 138, 137, 3,
	2, 2, 2, 139, 142, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2,
	2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 143, 145, 7, 94, 2, 2, 144,
	143, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147,
	7, 127, 2, 2, 147, 32, 3, 2, 2, 2, 148, 152, 9, 7, 2, 2, 149, 153, 9, 7,
	2, 2, 150, 153, 5, 7, 4, 2, 151, 153, 5, 25, 13, 2, 152, 149, 3, 2, 2,
	2, 152, 150, 3, 2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154,
	152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 34, 3, 2, 2, 2, 25, 2, 38,
	46, 50, 61, 63, 69, 72, 77, 82, 90, 92, 110, 112, 116, 123, 127, 129, 134,
	140, 144, 152, 154, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "", "'='",
}

var lexerSymbolicNames = []string{
	"", "SPACE", "COMMENT", "Float", "FloatArg", "IntArg", "StringArg", "ASSIGN",
	"ARG", "SLICE", "DICT", "KEY",
}

var lexerRuleNames = []string{
	"SPACE", "COMMENT", "DIGIT", "Float", "FloatArg", "IntArg", "StringArg",
	"ASSIGN", "COLON", "COMMA", "Dot", "UNDERLINE", "ARG", "SLICE", "DICT",
	"KEY",
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
	ConfigExprLexerSPACE     = 1
	ConfigExprLexerCOMMENT   = 2
	ConfigExprLexerFloat     = 3
	ConfigExprLexerFloatArg  = 4
	ConfigExprLexerIntArg    = 5
	ConfigExprLexerStringArg = 6
	ConfigExprLexerASSIGN    = 7
	ConfigExprLexerARG       = 8
	ConfigExprLexerSLICE     = 9
	ConfigExprLexerDICT      = 10
	ConfigExprLexerKEY       = 11
)
