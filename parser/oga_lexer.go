// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type OgaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ogalexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ogalexerLexerInit() {
	staticData := &ogalexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'make'", "'funke'", "'suppose say'",
		"'dapada'", "'otherwise'", "'dey play'", "'big pass'", "'resemble'",
		"'small pass'", "'no resemble'", "';'", "'='", "','", "'*'", "'/'",
		"'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "MAKE", "FUNC", "IF", "RETURN", "ELSE", "FOR", "GREATER",
		"EQUALS", "LESS", "NOT_EQUAL", "SEMI", "ASSIGN", "COMMA", "MUL", "DIV",
		"PLUS", "MINUS", "INT", "STR", "IDENTIFIER", "LINE_COMMENT", "EOS",
		"WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "MAKE", "FUNC", "IF", "RETURN", "ELSE",
		"FOR", "GREATER", "EQUALS", "LESS", "NOT_EQUAL", "SEMI", "ASSIGN", "COMMA",
		"MUL", "DIV", "PLUS", "MINUS", "INT", "STR", "ESC", "IDENTIFIER", "LINE_COMMENT",
		"EOS", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 225, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 173, 8, 21, 10,
		21, 12, 21, 176, 9, 21, 3, 21, 178, 8, 21, 1, 22, 1, 22, 1, 22, 5, 22,
		183, 8, 22, 10, 22, 12, 22, 186, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 194, 8, 23, 1, 24, 4, 24, 197, 8, 24, 11, 24, 12, 24,
		198, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 205, 8, 25, 10, 25, 12, 25, 208,
		9, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 4, 26, 215, 8, 26, 11, 26, 12,
		26, 216, 1, 27, 4, 27, 220, 8, 27, 11, 27, 12, 27, 221, 1, 27, 1, 27, 1,
		184, 0, 28, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 0, 49, 24, 51, 25, 53, 26, 55,
		27, 1, 0, 5, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 48, 57, 65, 90, 97, 122,
		2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 232, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 1, 57, 1, 0, 0, 0, 3, 59,
		1, 0, 0, 0, 5, 61, 1, 0, 0, 0, 7, 63, 1, 0, 0, 0, 9, 65, 1, 0, 0, 0, 11,
		70, 1, 0, 0, 0, 13, 76, 1, 0, 0, 0, 15, 88, 1, 0, 0, 0, 17, 95, 1, 0, 0,
		0, 19, 105, 1, 0, 0, 0, 21, 114, 1, 0, 0, 0, 23, 123, 1, 0, 0, 0, 25, 132,
		1, 0, 0, 0, 27, 143, 1, 0, 0, 0, 29, 155, 1, 0, 0, 0, 31, 157, 1, 0, 0,
		0, 33, 159, 1, 0, 0, 0, 35, 161, 1, 0, 0, 0, 37, 163, 1, 0, 0, 0, 39, 165,
		1, 0, 0, 0, 41, 167, 1, 0, 0, 0, 43, 177, 1, 0, 0, 0, 45, 179, 1, 0, 0,
		0, 47, 193, 1, 0, 0, 0, 49, 196, 1, 0, 0, 0, 51, 200, 1, 0, 0, 0, 53, 214,
		1, 0, 0, 0, 55, 219, 1, 0, 0, 0, 57, 58, 5, 40, 0, 0, 58, 2, 1, 0, 0, 0,
		59, 60, 5, 41, 0, 0, 60, 4, 1, 0, 0, 0, 61, 62, 5, 123, 0, 0, 62, 6, 1,
		0, 0, 0, 63, 64, 5, 125, 0, 0, 64, 8, 1, 0, 0, 0, 65, 66, 5, 109, 0, 0,
		66, 67, 5, 97, 0, 0, 67, 68, 5, 107, 0, 0, 68, 69, 5, 101, 0, 0, 69, 10,
		1, 0, 0, 0, 70, 71, 5, 102, 0, 0, 71, 72, 5, 117, 0, 0, 72, 73, 5, 110,
		0, 0, 73, 74, 5, 107, 0, 0, 74, 75, 5, 101, 0, 0, 75, 12, 1, 0, 0, 0, 76,
		77, 5, 115, 0, 0, 77, 78, 5, 117, 0, 0, 78, 79, 5, 112, 0, 0, 79, 80, 5,
		112, 0, 0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 115, 0, 0, 82, 83, 5, 101,
		0, 0, 83, 84, 5, 32, 0, 0, 84, 85, 5, 115, 0, 0, 85, 86, 5, 97, 0, 0, 86,
		87, 5, 121, 0, 0, 87, 14, 1, 0, 0, 0, 88, 89, 5, 100, 0, 0, 89, 90, 5,
		97, 0, 0, 90, 91, 5, 112, 0, 0, 91, 92, 5, 97, 0, 0, 92, 93, 5, 100, 0,
		0, 93, 94, 5, 97, 0, 0, 94, 16, 1, 0, 0, 0, 95, 96, 5, 111, 0, 0, 96, 97,
		5, 116, 0, 0, 97, 98, 5, 104, 0, 0, 98, 99, 5, 101, 0, 0, 99, 100, 5, 114,
		0, 0, 100, 101, 5, 119, 0, 0, 101, 102, 5, 105, 0, 0, 102, 103, 5, 115,
		0, 0, 103, 104, 5, 101, 0, 0, 104, 18, 1, 0, 0, 0, 105, 106, 5, 100, 0,
		0, 106, 107, 5, 101, 0, 0, 107, 108, 5, 121, 0, 0, 108, 109, 5, 32, 0,
		0, 109, 110, 5, 112, 0, 0, 110, 111, 5, 108, 0, 0, 111, 112, 5, 97, 0,
		0, 112, 113, 5, 121, 0, 0, 113, 20, 1, 0, 0, 0, 114, 115, 5, 98, 0, 0,
		115, 116, 5, 105, 0, 0, 116, 117, 5, 103, 0, 0, 117, 118, 5, 32, 0, 0,
		118, 119, 5, 112, 0, 0, 119, 120, 5, 97, 0, 0, 120, 121, 5, 115, 0, 0,
		121, 122, 5, 115, 0, 0, 122, 22, 1, 0, 0, 0, 123, 124, 5, 114, 0, 0, 124,
		125, 5, 101, 0, 0, 125, 126, 5, 115, 0, 0, 126, 127, 5, 101, 0, 0, 127,
		128, 5, 109, 0, 0, 128, 129, 5, 98, 0, 0, 129, 130, 5, 108, 0, 0, 130,
		131, 5, 101, 0, 0, 131, 24, 1, 0, 0, 0, 132, 133, 5, 115, 0, 0, 133, 134,
		5, 109, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 108, 0, 0, 136, 137,
		5, 108, 0, 0, 137, 138, 5, 32, 0, 0, 138, 139, 5, 112, 0, 0, 139, 140,
		5, 97, 0, 0, 140, 141, 5, 115, 0, 0, 141, 142, 5, 115, 0, 0, 142, 26, 1,
		0, 0, 0, 143, 144, 5, 110, 0, 0, 144, 145, 5, 111, 0, 0, 145, 146, 5, 32,
		0, 0, 146, 147, 5, 114, 0, 0, 147, 148, 5, 101, 0, 0, 148, 149, 5, 115,
		0, 0, 149, 150, 5, 101, 0, 0, 150, 151, 5, 109, 0, 0, 151, 152, 5, 98,
		0, 0, 152, 153, 5, 108, 0, 0, 153, 154, 5, 101, 0, 0, 154, 28, 1, 0, 0,
		0, 155, 156, 5, 59, 0, 0, 156, 30, 1, 0, 0, 0, 157, 158, 5, 61, 0, 0, 158,
		32, 1, 0, 0, 0, 159, 160, 5, 44, 0, 0, 160, 34, 1, 0, 0, 0, 161, 162, 5,
		42, 0, 0, 162, 36, 1, 0, 0, 0, 163, 164, 5, 47, 0, 0, 164, 38, 1, 0, 0,
		0, 165, 166, 5, 43, 0, 0, 166, 40, 1, 0, 0, 0, 167, 168, 5, 45, 0, 0, 168,
		42, 1, 0, 0, 0, 169, 178, 5, 48, 0, 0, 170, 174, 7, 0, 0, 0, 171, 173,
		7, 1, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0,
		0, 0, 174, 175, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0,
		177, 169, 1, 0, 0, 0, 177, 170, 1, 0, 0, 0, 178, 44, 1, 0, 0, 0, 179, 184,
		5, 34, 0, 0, 180, 183, 3, 47, 23, 0, 181, 183, 9, 0, 0, 0, 182, 180, 1,
		0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 185, 1, 0, 0,
		0, 184, 182, 1, 0, 0, 0, 185, 187, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187,
		188, 5, 34, 0, 0, 188, 46, 1, 0, 0, 0, 189, 190, 5, 92, 0, 0, 190, 194,
		5, 34, 0, 0, 191, 192, 5, 92, 0, 0, 192, 194, 5, 92, 0, 0, 193, 189, 1,
		0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 48, 1, 0, 0, 0, 195, 197, 7, 2, 0,
		0, 196, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 50, 1, 0, 0, 0, 200, 201, 5, 47, 0, 0, 201, 202,
		5, 47, 0, 0, 202, 206, 1, 0, 0, 0, 203, 205, 8, 3, 0, 0, 204, 203, 1, 0,
		0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0,
		207, 209, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 210, 3, 53, 26, 0, 210,
		211, 1, 0, 0, 0, 211, 212, 6, 25, 0, 0, 212, 52, 1, 0, 0, 0, 213, 215,
		7, 3, 0, 0, 214, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 214, 1, 0,
		0, 0, 216, 217, 1, 0, 0, 0, 217, 54, 1, 0, 0, 0, 218, 220, 7, 4, 0, 0,
		219, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 6, 27, 0, 0, 224, 56,
		1, 0, 0, 0, 10, 0, 174, 177, 182, 184, 193, 198, 206, 216, 221, 1, 6, 0,
		0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// OgaLexerInit initializes any static state used to implement OgaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewOgaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func OgaLexerInit() {
	staticData := &ogalexerLexerStaticData
	staticData.once.Do(ogalexerLexerInit)
}

// NewOgaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewOgaLexer(input antlr.CharStream) *OgaLexer {
	OgaLexerInit()
	l := new(OgaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ogalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Oga.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// OgaLexer tokens.
const (
	OgaLexerT__0         = 1
	OgaLexerT__1         = 2
	OgaLexerT__2         = 3
	OgaLexerT__3         = 4
	OgaLexerMAKE         = 5
	OgaLexerFUNC         = 6
	OgaLexerIF           = 7
	OgaLexerRETURN       = 8
	OgaLexerELSE         = 9
	OgaLexerFOR          = 10
	OgaLexerGREATER      = 11
	OgaLexerEQUALS       = 12
	OgaLexerLESS         = 13
	OgaLexerNOT_EQUAL    = 14
	OgaLexerSEMI         = 15
	OgaLexerASSIGN       = 16
	OgaLexerCOMMA        = 17
	OgaLexerMUL          = 18
	OgaLexerDIV          = 19
	OgaLexerPLUS         = 20
	OgaLexerMINUS        = 21
	OgaLexerINT          = 22
	OgaLexerSTR          = 23
	OgaLexerIDENTIFIER   = 24
	OgaLexerLINE_COMMENT = 25
	OgaLexerEOS          = 26
	OgaLexerWS           = 27
)
