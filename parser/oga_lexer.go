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
		"", "'('", "')'", "'{'", "'}'", "'*'", "'/'", "'+'", "'-'", "','", "'make'",
		"'funke'", "'suppose say'", "'dapada'", "'otherwise'", "'dey play'",
		"'big pass'", "'resemble'", "'small pass'", "'no resemble'", "';'",
		"'='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "MAKE", "FUNC", "IF", "RETURN",
		"ELSE", "FOR", "GREATER", "EQUALS", "LESS", "NOT_EQUAL", "SEMI", "ASSIGN",
		"INT", "IDENTIFIER", "LINE_COMMENT", "EOS", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"MAKE", "FUNC", "IF", "RETURN", "ELSE", "FOR", "GREATER", "EQUALS",
		"LESS", "NOT_EQUAL", "SEMI", "ASSIGN", "INT", "IDENTIFIER", "LINE_COMMENT",
		"EOS", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 205, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 169, 8, 21, 10,
		21, 12, 21, 172, 9, 21, 3, 21, 174, 8, 21, 1, 22, 4, 22, 177, 8, 22, 11,
		22, 12, 22, 178, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 185, 8, 23, 10, 23,
		12, 23, 188, 9, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 4, 24, 195, 8, 24,
		11, 24, 12, 24, 196, 1, 25, 4, 25, 200, 8, 25, 11, 25, 12, 25, 201, 1,
		25, 1, 25, 0, 0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		1, 0, 5, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 48, 57, 65, 90, 97, 122, 2,
		0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 210, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 1, 53, 1, 0, 0, 0, 3, 55, 1, 0, 0, 0, 5, 57, 1,
		0, 0, 0, 7, 59, 1, 0, 0, 0, 9, 61, 1, 0, 0, 0, 11, 63, 1, 0, 0, 0, 13,
		65, 1, 0, 0, 0, 15, 67, 1, 0, 0, 0, 17, 69, 1, 0, 0, 0, 19, 71, 1, 0, 0,
		0, 21, 76, 1, 0, 0, 0, 23, 82, 1, 0, 0, 0, 25, 94, 1, 0, 0, 0, 27, 101,
		1, 0, 0, 0, 29, 111, 1, 0, 0, 0, 31, 120, 1, 0, 0, 0, 33, 129, 1, 0, 0,
		0, 35, 138, 1, 0, 0, 0, 37, 149, 1, 0, 0, 0, 39, 161, 1, 0, 0, 0, 41, 163,
		1, 0, 0, 0, 43, 173, 1, 0, 0, 0, 45, 176, 1, 0, 0, 0, 47, 180, 1, 0, 0,
		0, 49, 194, 1, 0, 0, 0, 51, 199, 1, 0, 0, 0, 53, 54, 5, 40, 0, 0, 54, 2,
		1, 0, 0, 0, 55, 56, 5, 41, 0, 0, 56, 4, 1, 0, 0, 0, 57, 58, 5, 123, 0,
		0, 58, 6, 1, 0, 0, 0, 59, 60, 5, 125, 0, 0, 60, 8, 1, 0, 0, 0, 61, 62,
		5, 42, 0, 0, 62, 10, 1, 0, 0, 0, 63, 64, 5, 47, 0, 0, 64, 12, 1, 0, 0,
		0, 65, 66, 5, 43, 0, 0, 66, 14, 1, 0, 0, 0, 67, 68, 5, 45, 0, 0, 68, 16,
		1, 0, 0, 0, 69, 70, 5, 44, 0, 0, 70, 18, 1, 0, 0, 0, 71, 72, 5, 109, 0,
		0, 72, 73, 5, 97, 0, 0, 73, 74, 5, 107, 0, 0, 74, 75, 5, 101, 0, 0, 75,
		20, 1, 0, 0, 0, 76, 77, 5, 102, 0, 0, 77, 78, 5, 117, 0, 0, 78, 79, 5,
		110, 0, 0, 79, 80, 5, 107, 0, 0, 80, 81, 5, 101, 0, 0, 81, 22, 1, 0, 0,
		0, 82, 83, 5, 115, 0, 0, 83, 84, 5, 117, 0, 0, 84, 85, 5, 112, 0, 0, 85,
		86, 5, 112, 0, 0, 86, 87, 5, 111, 0, 0, 87, 88, 5, 115, 0, 0, 88, 89, 5,
		101, 0, 0, 89, 90, 5, 32, 0, 0, 90, 91, 5, 115, 0, 0, 91, 92, 5, 97, 0,
		0, 92, 93, 5, 121, 0, 0, 93, 24, 1, 0, 0, 0, 94, 95, 5, 100, 0, 0, 95,
		96, 5, 97, 0, 0, 96, 97, 5, 112, 0, 0, 97, 98, 5, 97, 0, 0, 98, 99, 5,
		100, 0, 0, 99, 100, 5, 97, 0, 0, 100, 26, 1, 0, 0, 0, 101, 102, 5, 111,
		0, 0, 102, 103, 5, 116, 0, 0, 103, 104, 5, 104, 0, 0, 104, 105, 5, 101,
		0, 0, 105, 106, 5, 114, 0, 0, 106, 107, 5, 119, 0, 0, 107, 108, 5, 105,
		0, 0, 108, 109, 5, 115, 0, 0, 109, 110, 5, 101, 0, 0, 110, 28, 1, 0, 0,
		0, 111, 112, 5, 100, 0, 0, 112, 113, 5, 101, 0, 0, 113, 114, 5, 121, 0,
		0, 114, 115, 5, 32, 0, 0, 115, 116, 5, 112, 0, 0, 116, 117, 5, 108, 0,
		0, 117, 118, 5, 97, 0, 0, 118, 119, 5, 121, 0, 0, 119, 30, 1, 0, 0, 0,
		120, 121, 5, 98, 0, 0, 121, 122, 5, 105, 0, 0, 122, 123, 5, 103, 0, 0,
		123, 124, 5, 32, 0, 0, 124, 125, 5, 112, 0, 0, 125, 126, 5, 97, 0, 0, 126,
		127, 5, 115, 0, 0, 127, 128, 5, 115, 0, 0, 128, 32, 1, 0, 0, 0, 129, 130,
		5, 114, 0, 0, 130, 131, 5, 101, 0, 0, 131, 132, 5, 115, 0, 0, 132, 133,
		5, 101, 0, 0, 133, 134, 5, 109, 0, 0, 134, 135, 5, 98, 0, 0, 135, 136,
		5, 108, 0, 0, 136, 137, 5, 101, 0, 0, 137, 34, 1, 0, 0, 0, 138, 139, 5,
		115, 0, 0, 139, 140, 5, 109, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5,
		108, 0, 0, 142, 143, 5, 108, 0, 0, 143, 144, 5, 32, 0, 0, 144, 145, 5,
		112, 0, 0, 145, 146, 5, 97, 0, 0, 146, 147, 5, 115, 0, 0, 147, 148, 5,
		115, 0, 0, 148, 36, 1, 0, 0, 0, 149, 150, 5, 110, 0, 0, 150, 151, 5, 111,
		0, 0, 151, 152, 5, 32, 0, 0, 152, 153, 5, 114, 0, 0, 153, 154, 5, 101,
		0, 0, 154, 155, 5, 115, 0, 0, 155, 156, 5, 101, 0, 0, 156, 157, 5, 109,
		0, 0, 157, 158, 5, 98, 0, 0, 158, 159, 5, 108, 0, 0, 159, 160, 5, 101,
		0, 0, 160, 38, 1, 0, 0, 0, 161, 162, 5, 59, 0, 0, 162, 40, 1, 0, 0, 0,
		163, 164, 5, 61, 0, 0, 164, 42, 1, 0, 0, 0, 165, 174, 5, 48, 0, 0, 166,
		170, 7, 0, 0, 0, 167, 169, 7, 1, 0, 0, 168, 167, 1, 0, 0, 0, 169, 172,
		1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 174, 1, 0,
		0, 0, 172, 170, 1, 0, 0, 0, 173, 165, 1, 0, 0, 0, 173, 166, 1, 0, 0, 0,
		174, 44, 1, 0, 0, 0, 175, 177, 7, 2, 0, 0, 176, 175, 1, 0, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 46, 1, 0,
		0, 0, 180, 181, 5, 47, 0, 0, 181, 182, 5, 47, 0, 0, 182, 186, 1, 0, 0,
		0, 183, 185, 8, 3, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186,
		1, 0, 0, 0, 189, 190, 3, 49, 24, 0, 190, 191, 1, 0, 0, 0, 191, 192, 6,
		23, 0, 0, 192, 48, 1, 0, 0, 0, 193, 195, 7, 3, 0, 0, 194, 193, 1, 0, 0,
		0, 195, 196, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		50, 1, 0, 0, 0, 198, 200, 7, 4, 0, 0, 199, 198, 1, 0, 0, 0, 200, 201, 1,
		0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0,
		0, 203, 204, 6, 25, 0, 0, 204, 52, 1, 0, 0, 0, 7, 0, 170, 173, 178, 186,
		196, 201, 1, 6, 0, 0,
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
	OgaLexerT__4         = 5
	OgaLexerT__5         = 6
	OgaLexerT__6         = 7
	OgaLexerT__7         = 8
	OgaLexerT__8         = 9
	OgaLexerMAKE         = 10
	OgaLexerFUNC         = 11
	OgaLexerIF           = 12
	OgaLexerRETURN       = 13
	OgaLexerELSE         = 14
	OgaLexerFOR          = 15
	OgaLexerGREATER      = 16
	OgaLexerEQUALS       = 17
	OgaLexerLESS         = 18
	OgaLexerNOT_EQUAL    = 19
	OgaLexerSEMI         = 20
	OgaLexerASSIGN       = 21
	OgaLexerINT          = 22
	OgaLexerIDENTIFIER   = 23
	OgaLexerLINE_COMMENT = 24
	OgaLexerEOS          = 25
	OgaLexerWS           = 26
)
