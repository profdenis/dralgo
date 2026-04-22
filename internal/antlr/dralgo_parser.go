// Code generated from /home/denis/projets/profdenis/dralgo/grammar/dralgo/DrAlgo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package antlr // DrAlgo
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DrAlgoParser struct {
	*antlr.BaseParser
}

var DrAlgoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dralgoParserInit() {
	staticData := &DrAlgoParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "'++'", "'--'", "'**'",
		"'*'", "'/'", "'%'", "'+'", "'-'", "'<='", "'>='", "'<'", "'>'", "'=='",
		"'!='", "'='", "'('", "')'", "'['", "']'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "SI", "SINON", "ALORS", "AFFICHER", "LIRE", "LIRE_ENTIER", "SINON_SI",
		"FIN_SI", "TANT_QUE", "FIN_TANT_QUE", "INCR", "DECR", "POW", "MUL",
		"DIV", "MOD", "PLUS", "MINUS", "LTE", "GTE", "LT", "GT", "EQ", "NEQ",
		"ASSIGN", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "COMMA", "SEMI", "IDENT",
		"FLOAT", "INT", "STRING", "LINE_COMMENT", "NEWLINE", "WS",
	}
	staticData.RuleNames = []string{
		"program", "stmtList", "stmt", "assignStmt", "incDecStmt", "printStmt",
		"printArgs", "readStmt", "readIntStmt", "ifStmt", "elseIfPart", "elsePart",
		"whileStmt", "block", "sep", "expr", "equality", "comparison", "addition",
		"multiplication", "power", "unary", "primary", "listLiteral", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 239, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 5, 0, 52, 8,
		0, 10, 0, 12, 0, 55, 9, 0, 1, 0, 3, 0, 58, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		4, 1, 64, 8, 1, 11, 1, 12, 1, 65, 1, 1, 1, 1, 5, 1, 70, 8, 1, 10, 1, 12,
		1, 73, 9, 1, 1, 1, 5, 1, 76, 8, 1, 10, 1, 12, 1, 79, 9, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 88, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 99, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 104,
		8, 6, 10, 6, 12, 6, 107, 9, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 3, 9, 118, 8, 9, 1, 9, 4, 9, 121, 8, 9, 11, 9, 12, 9, 122,
		1, 9, 1, 9, 5, 9, 127, 8, 9, 10, 9, 12, 9, 130, 9, 9, 1, 9, 3, 9, 133,
		8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 140, 8, 10, 1, 10, 4, 10,
		143, 8, 10, 11, 10, 12, 10, 144, 1, 10, 1, 10, 1, 11, 1, 11, 4, 11, 151,
		8, 11, 11, 11, 12, 11, 152, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 4, 12, 160,
		8, 12, 11, 12, 12, 12, 161, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 176, 8, 16, 10, 16, 12, 16,
		179, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 184, 8, 17, 10, 17, 12, 17, 187,
		9, 17, 1, 18, 1, 18, 1, 18, 5, 18, 192, 8, 18, 10, 18, 12, 18, 195, 9,
		18, 1, 19, 1, 19, 1, 19, 5, 19, 200, 8, 19, 10, 19, 12, 19, 203, 9, 19,
		1, 20, 1, 20, 1, 20, 3, 20, 208, 8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 213,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 222, 8,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 228, 8, 23, 10, 23, 12, 23, 231,
		9, 23, 3, 23, 233, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 0, 0, 25,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
		38, 40, 42, 44, 46, 48, 0, 6, 1, 0, 11, 12, 1, 0, 23, 24, 1, 0, 19, 22,
		1, 0, 17, 18, 1, 0, 14, 16, 1, 0, 33, 35, 245, 0, 53, 1, 0, 0, 0, 2, 61,
		1, 0, 0, 0, 4, 87, 1, 0, 0, 0, 6, 89, 1, 0, 0, 0, 8, 93, 1, 0, 0, 0, 10,
		96, 1, 0, 0, 0, 12, 100, 1, 0, 0, 0, 14, 108, 1, 0, 0, 0, 16, 111, 1, 0,
		0, 0, 18, 114, 1, 0, 0, 0, 20, 136, 1, 0, 0, 0, 22, 148, 1, 0, 0, 0, 24,
		156, 1, 0, 0, 0, 26, 166, 1, 0, 0, 0, 28, 168, 1, 0, 0, 0, 30, 170, 1,
		0, 0, 0, 32, 172, 1, 0, 0, 0, 34, 180, 1, 0, 0, 0, 36, 188, 1, 0, 0, 0,
		38, 196, 1, 0, 0, 0, 40, 204, 1, 0, 0, 0, 42, 212, 1, 0, 0, 0, 44, 221,
		1, 0, 0, 0, 46, 223, 1, 0, 0, 0, 48, 236, 1, 0, 0, 0, 50, 52, 3, 28, 14,
		0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54,
		1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 58, 3, 2, 1, 0,
		57, 56, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 5,
		0, 0, 1, 60, 1, 1, 0, 0, 0, 61, 71, 3, 4, 2, 0, 62, 64, 3, 28, 14, 0, 63,
		62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0,
		0, 66, 67, 1, 0, 0, 0, 67, 68, 3, 4, 2, 0, 68, 70, 1, 0, 0, 0, 69, 63,
		1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0,
		72, 77, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 76, 3, 28, 14, 0, 75, 74, 1,
		0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78,
		3, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 88, 3, 6, 3, 0, 81, 88, 3, 8, 4,
		0, 82, 88, 3, 10, 5, 0, 83, 88, 3, 14, 7, 0, 84, 88, 3, 16, 8, 0, 85, 88,
		3, 18, 9, 0, 86, 88, 3, 24, 12, 0, 87, 80, 1, 0, 0, 0, 87, 81, 1, 0, 0,
		0, 87, 82, 1, 0, 0, 0, 87, 83, 1, 0, 0, 0, 87, 84, 1, 0, 0, 0, 87, 85,
		1, 0, 0, 0, 87, 86, 1, 0, 0, 0, 88, 5, 1, 0, 0, 0, 89, 90, 5, 32, 0, 0,
		90, 91, 5, 25, 0, 0, 91, 92, 3, 30, 15, 0, 92, 7, 1, 0, 0, 0, 93, 94, 5,
		32, 0, 0, 94, 95, 7, 0, 0, 0, 95, 9, 1, 0, 0, 0, 96, 98, 5, 4, 0, 0, 97,
		99, 3, 12, 6, 0, 98, 97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 11, 1, 0,
		0, 0, 100, 105, 3, 30, 15, 0, 101, 102, 5, 30, 0, 0, 102, 104, 3, 30, 15,
		0, 103, 101, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105,
		106, 1, 0, 0, 0, 106, 13, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109, 5,
		5, 0, 0, 109, 110, 5, 32, 0, 0, 110, 15, 1, 0, 0, 0, 111, 112, 5, 6, 0,
		0, 112, 113, 5, 32, 0, 0, 113, 17, 1, 0, 0, 0, 114, 115, 5, 1, 0, 0, 115,
		117, 3, 30, 15, 0, 116, 118, 5, 3, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118,
		1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119, 121, 3, 28, 14, 0, 120, 119, 1,
		0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0,
		0, 123, 124, 1, 0, 0, 0, 124, 128, 3, 26, 13, 0, 125, 127, 3, 20, 10, 0,
		126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128,
		129, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 133,
		3, 22, 11, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 1,
		0, 0, 0, 134, 135, 5, 8, 0, 0, 135, 19, 1, 0, 0, 0, 136, 137, 5, 7, 0,
		0, 137, 139, 3, 30, 15, 0, 138, 140, 5, 3, 0, 0, 139, 138, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 143, 3, 28, 14, 0, 142,
		141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145,
		1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 3, 26, 13, 0, 147, 21, 1, 0,
		0, 0, 148, 150, 5, 2, 0, 0, 149, 151, 3, 28, 14, 0, 150, 149, 1, 0, 0,
		0, 151, 152, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 155, 3, 26, 13, 0, 155, 23, 1, 0, 0, 0, 156, 157,
		5, 9, 0, 0, 157, 159, 3, 30, 15, 0, 158, 160, 3, 28, 14, 0, 159, 158, 1,
		0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 163, 1, 0, 0, 0, 163, 164, 3, 26, 13, 0, 164, 165, 5, 10, 0, 0,
		165, 25, 1, 0, 0, 0, 166, 167, 3, 2, 1, 0, 167, 27, 1, 0, 0, 0, 168, 169,
		5, 37, 0, 0, 169, 29, 1, 0, 0, 0, 170, 171, 3, 32, 16, 0, 171, 31, 1, 0,
		0, 0, 172, 177, 3, 34, 17, 0, 173, 174, 7, 1, 0, 0, 174, 176, 3, 34, 17,
		0, 175, 173, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 33, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 185, 3,
		36, 18, 0, 181, 182, 7, 2, 0, 0, 182, 184, 3, 36, 18, 0, 183, 181, 1, 0,
		0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0,
		186, 35, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 193, 3, 38, 19, 0, 189,
		190, 7, 3, 0, 0, 190, 192, 3, 38, 19, 0, 191, 189, 1, 0, 0, 0, 192, 195,
		1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 37, 1, 0,
		0, 0, 195, 193, 1, 0, 0, 0, 196, 201, 3, 40, 20, 0, 197, 198, 7, 4, 0,
		0, 198, 200, 3, 40, 20, 0, 199, 197, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0,
		201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 39, 1, 0, 0, 0, 203, 201,
		1, 0, 0, 0, 204, 207, 3, 42, 21, 0, 205, 206, 5, 13, 0, 0, 206, 208, 3,
		40, 20, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 41, 1, 0, 0,
		0, 209, 210, 7, 3, 0, 0, 210, 213, 3, 42, 21, 0, 211, 213, 3, 44, 22, 0,
		212, 209, 1, 0, 0, 0, 212, 211, 1, 0, 0, 0, 213, 43, 1, 0, 0, 0, 214, 222,
		3, 46, 23, 0, 215, 222, 3, 48, 24, 0, 216, 222, 5, 32, 0, 0, 217, 218,
		5, 26, 0, 0, 218, 219, 3, 30, 15, 0, 219, 220, 5, 27, 0, 0, 220, 222, 1,
		0, 0, 0, 221, 214, 1, 0, 0, 0, 221, 215, 1, 0, 0, 0, 221, 216, 1, 0, 0,
		0, 221, 217, 1, 0, 0, 0, 222, 45, 1, 0, 0, 0, 223, 232, 5, 28, 0, 0, 224,
		229, 3, 30, 15, 0, 225, 226, 5, 30, 0, 0, 226, 228, 3, 30, 15, 0, 227,
		225, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230,
		1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 224, 1, 0,
		0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 5, 29, 0, 0,
		235, 47, 1, 0, 0, 0, 236, 237, 7, 5, 0, 0, 237, 49, 1, 0, 0, 0, 25, 53,
		57, 65, 71, 77, 87, 98, 105, 117, 122, 128, 132, 139, 144, 152, 161, 177,
		185, 193, 201, 207, 212, 221, 229, 232,
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

// DrAlgoParserInit initializes any static state used to implement DrAlgoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDrAlgoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DrAlgoParserInit() {
	staticData := &DrAlgoParserStaticData
	staticData.once.Do(dralgoParserInit)
}

// NewDrAlgoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDrAlgoParser(input antlr.TokenStream) *DrAlgoParser {
	DrAlgoParserInit()
	this := new(DrAlgoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DrAlgoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DrAlgo.g4"

	return this
}

// DrAlgoParser tokens.
const (
	DrAlgoParserEOF          = antlr.TokenEOF
	DrAlgoParserSI           = 1
	DrAlgoParserSINON        = 2
	DrAlgoParserALORS        = 3
	DrAlgoParserAFFICHER     = 4
	DrAlgoParserLIRE         = 5
	DrAlgoParserLIRE_ENTIER  = 6
	DrAlgoParserSINON_SI     = 7
	DrAlgoParserFIN_SI       = 8
	DrAlgoParserTANT_QUE     = 9
	DrAlgoParserFIN_TANT_QUE = 10
	DrAlgoParserINCR         = 11
	DrAlgoParserDECR         = 12
	DrAlgoParserPOW          = 13
	DrAlgoParserMUL          = 14
	DrAlgoParserDIV          = 15
	DrAlgoParserMOD          = 16
	DrAlgoParserPLUS         = 17
	DrAlgoParserMINUS        = 18
	DrAlgoParserLTE          = 19
	DrAlgoParserGTE          = 20
	DrAlgoParserLT           = 21
	DrAlgoParserGT           = 22
	DrAlgoParserEQ           = 23
	DrAlgoParserNEQ          = 24
	DrAlgoParserASSIGN       = 25
	DrAlgoParserLPAREN       = 26
	DrAlgoParserRPAREN       = 27
	DrAlgoParserLBRACK       = 28
	DrAlgoParserRBRACK       = 29
	DrAlgoParserCOMMA        = 30
	DrAlgoParserSEMI         = 31
	DrAlgoParserIDENT        = 32
	DrAlgoParserFLOAT        = 33
	DrAlgoParserINT          = 34
	DrAlgoParserSTRING       = 35
	DrAlgoParserLINE_COMMENT = 36
	DrAlgoParserNEWLINE      = 37
	DrAlgoParserWS           = 38
)

// DrAlgoParser rules.
const (
	DrAlgoParserRULE_program        = 0
	DrAlgoParserRULE_stmtList       = 1
	DrAlgoParserRULE_stmt           = 2
	DrAlgoParserRULE_assignStmt     = 3
	DrAlgoParserRULE_incDecStmt     = 4
	DrAlgoParserRULE_printStmt      = 5
	DrAlgoParserRULE_printArgs      = 6
	DrAlgoParserRULE_readStmt       = 7
	DrAlgoParserRULE_readIntStmt    = 8
	DrAlgoParserRULE_ifStmt         = 9
	DrAlgoParserRULE_elseIfPart     = 10
	DrAlgoParserRULE_elsePart       = 11
	DrAlgoParserRULE_whileStmt      = 12
	DrAlgoParserRULE_block          = 13
	DrAlgoParserRULE_sep            = 14
	DrAlgoParserRULE_expr           = 15
	DrAlgoParserRULE_equality       = 16
	DrAlgoParserRULE_comparison     = 17
	DrAlgoParserRULE_addition       = 18
	DrAlgoParserRULE_multiplication = 19
	DrAlgoParserRULE_power          = 20
	DrAlgoParserRULE_unary          = 21
	DrAlgoParserRULE_primary        = 22
	DrAlgoParserRULE_listLiteral    = 23
	DrAlgoParserRULE_literal        = 24
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllSep() []ISepContext
	Sep(i int) ISepContext
	StmtList() IStmtListContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserEOF, 0)
}

func (s *ProgramContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *ProgramContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DrAlgoParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DrAlgoParserNEWLINE {
		{
			p.SetState(50)
			p.Sep()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4294967922) != 0 {
		{
			p.SetState(56)
			p.StmtList()
		}

	}
	{
		p.SetState(59)
		p.Match(DrAlgoParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtListContext is an interface to support dynamic dispatch.
type IStmtListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsStmtListContext differentiates from other interfaces.
	IsStmtListContext()
}

type StmtListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtListContext() *StmtListContext {
	var p = new(StmtListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_stmtList
	return p
}

func InitEmptyStmtListContext(p *StmtListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_stmtList
}

func (*StmtListContext) IsStmtListContext() {}

func NewStmtListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtListContext {
	var p = new(StmtListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_stmtList

	return p
}

func (s *StmtListContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtListContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtListContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *StmtListContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *StmtListContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *StmtListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitStmtList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) StmtList() (localctx IStmtListContext) {
	localctx = NewStmtListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DrAlgoParserRULE_stmtList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Stmt()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == DrAlgoParserNEWLINE {
				{
					p.SetState(62)
					p.Sep()
				}

				p.SetState(65)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(67)
				p.Stmt()
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DrAlgoParserNEWLINE {
		{
			p.SetState(74)
			p.Sep()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignStmt() IAssignStmtContext
	IncDecStmt() IIncDecStmtContext
	PrintStmt() IPrintStmtContext
	ReadStmt() IReadStmtContext
	ReadIntStmt() IReadIntStmtContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StmtContext) IncDecStmt() IIncDecStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncDecStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncDecStmtContext)
}

func (s *StmtContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StmtContext) ReadStmt() IReadStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadStmtContext)
}

func (s *StmtContext) ReadIntStmt() IReadIntStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadIntStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadIntStmtContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DrAlgoParserRULE_stmt)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.AssignStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.IncDecStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.PrintStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.ReadStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.ReadIntStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(85)
			p.IfStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(86)
			p.WhileStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_assignStmt
	return p
}

func InitEmptyAssignStmtContext(p *AssignStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_assignStmt
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserIDENT, 0)
}

func (s *AssignStmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserASSIGN, 0)
}

func (s *AssignStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) AssignStmt() (localctx IAssignStmtContext) {
	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DrAlgoParserRULE_assignStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(DrAlgoParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Match(DrAlgoParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncDecStmtContext is an interface to support dynamic dispatch.
type IIncDecStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	INCR() antlr.TerminalNode
	DECR() antlr.TerminalNode

	// IsIncDecStmtContext differentiates from other interfaces.
	IsIncDecStmtContext()
}

type IncDecStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncDecStmtContext() *IncDecStmtContext {
	var p = new(IncDecStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_incDecStmt
	return p
}

func InitEmptyIncDecStmtContext(p *IncDecStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_incDecStmt
}

func (*IncDecStmtContext) IsIncDecStmtContext() {}

func NewIncDecStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncDecStmtContext {
	var p = new(IncDecStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_incDecStmt

	return p
}

func (s *IncDecStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncDecStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserIDENT, 0)
}

func (s *IncDecStmtContext) INCR() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserINCR, 0)
}

func (s *IncDecStmtContext) DECR() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserDECR, 0)
}

func (s *IncDecStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncDecStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitIncDecStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) IncDecStmt() (localctx IIncDecStmtContext) {
	localctx = NewIncDecStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DrAlgoParserRULE_incDecStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(DrAlgoParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DrAlgoParserINCR || _la == DrAlgoParserDECR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AFFICHER() antlr.TerminalNode
	PrintArgs() IPrintArgsContext

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) AFFICHER() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserAFFICHER, 0)
}

func (s *PrintStmtContext) PrintArgs() IPrintArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintArgsContext)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DrAlgoParserRULE_printStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(DrAlgoParserAFFICHER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64760446976) != 0 {
		{
			p.SetState(97)
			p.PrintArgs()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintArgsContext is an interface to support dynamic dispatch.
type IPrintArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPrintArgsContext differentiates from other interfaces.
	IsPrintArgsContext()
}

type PrintArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintArgsContext() *PrintArgsContext {
	var p = new(PrintArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_printArgs
	return p
}

func InitEmptyPrintArgsContext(p *PrintArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_printArgs
}

func (*PrintArgsContext) IsPrintArgsContext() {}

func NewPrintArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintArgsContext {
	var p = new(PrintArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_printArgs

	return p
}

func (s *PrintArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserCOMMA)
}

func (s *PrintArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserCOMMA, i)
}

func (s *PrintArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitPrintArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) PrintArgs() (localctx IPrintArgsContext) {
	localctx = NewPrintArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DrAlgoParserRULE_printArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Expr()
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DrAlgoParserCOMMA {
		{
			p.SetState(101)
			p.Match(DrAlgoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Expr()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadStmtContext is an interface to support dynamic dispatch.
type IReadStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIRE() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsReadStmtContext differentiates from other interfaces.
	IsReadStmtContext()
}

type ReadStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadStmtContext() *ReadStmtContext {
	var p = new(ReadStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_readStmt
	return p
}

func InitEmptyReadStmtContext(p *ReadStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_readStmt
}

func (*ReadStmtContext) IsReadStmtContext() {}

func NewReadStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadStmtContext {
	var p = new(ReadStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_readStmt

	return p
}

func (s *ReadStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadStmtContext) LIRE() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserLIRE, 0)
}

func (s *ReadStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserIDENT, 0)
}

func (s *ReadStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitReadStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) ReadStmt() (localctx IReadStmtContext) {
	localctx = NewReadStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DrAlgoParserRULE_readStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(DrAlgoParserLIRE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(DrAlgoParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadIntStmtContext is an interface to support dynamic dispatch.
type IReadIntStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIRE_ENTIER() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsReadIntStmtContext differentiates from other interfaces.
	IsReadIntStmtContext()
}

type ReadIntStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadIntStmtContext() *ReadIntStmtContext {
	var p = new(ReadIntStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_readIntStmt
	return p
}

func InitEmptyReadIntStmtContext(p *ReadIntStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_readIntStmt
}

func (*ReadIntStmtContext) IsReadIntStmtContext() {}

func NewReadIntStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadIntStmtContext {
	var p = new(ReadIntStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_readIntStmt

	return p
}

func (s *ReadIntStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadIntStmtContext) LIRE_ENTIER() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserLIRE_ENTIER, 0)
}

func (s *ReadIntStmtContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserIDENT, 0)
}

func (s *ReadIntStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadIntStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadIntStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitReadIntStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) ReadIntStmt() (localctx IReadIntStmtContext) {
	localctx = NewReadIntStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DrAlgoParserRULE_readIntStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(DrAlgoParserLIRE_ENTIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(DrAlgoParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SI() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext
	FIN_SI() antlr.TerminalNode
	ALORS() antlr.TerminalNode
	AllSep() []ISepContext
	Sep(i int) ISepContext
	AllElseIfPart() []IElseIfPartContext
	ElseIfPart(i int) IElseIfPartContext
	ElsePart() IElsePartContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) SI() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserSI, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) FIN_SI() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserFIN_SI, 0)
}

func (s *IfStmtContext) ALORS() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserALORS, 0)
}

func (s *IfStmtContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *IfStmtContext) AllElseIfPart() []IElseIfPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfPartContext); ok {
			len++
		}
	}

	tst := make([]IElseIfPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfPartContext); ok {
			tst[i] = t.(IElseIfPartContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) ElseIfPart(i int) IElseIfPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfPartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfPartContext)
}

func (s *IfStmtContext) ElsePart() IElsePartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsePartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsePartContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DrAlgoParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(DrAlgoParserSI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Expr()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DrAlgoParserALORS {
		{
			p.SetState(116)
			p.Match(DrAlgoParserALORS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DrAlgoParserNEWLINE {
		{
			p.SetState(119)
			p.Sep()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Block()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DrAlgoParserSINON_SI {
		{
			p.SetState(125)
			p.ElseIfPart()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DrAlgoParserSINON {
		{
			p.SetState(131)
			p.ElsePart()
		}

	}
	{
		p.SetState(134)
		p.Match(DrAlgoParserFIN_SI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseIfPartContext is an interface to support dynamic dispatch.
type IElseIfPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SINON_SI() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext
	ALORS() antlr.TerminalNode
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsElseIfPartContext differentiates from other interfaces.
	IsElseIfPartContext()
}

type ElseIfPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfPartContext() *ElseIfPartContext {
	var p = new(ElseIfPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_elseIfPart
	return p
}

func InitEmptyElseIfPartContext(p *ElseIfPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_elseIfPart
}

func (*ElseIfPartContext) IsElseIfPartContext() {}

func NewElseIfPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfPartContext {
	var p = new(ElseIfPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_elseIfPart

	return p
}

func (s *ElseIfPartContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfPartContext) SINON_SI() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserSINON_SI, 0)
}

func (s *ElseIfPartContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElseIfPartContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfPartContext) ALORS() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserALORS, 0)
}

func (s *ElseIfPartContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *ElseIfPartContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *ElseIfPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfPartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitElseIfPart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) ElseIfPart() (localctx IElseIfPartContext) {
	localctx = NewElseIfPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DrAlgoParserRULE_elseIfPart)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(DrAlgoParserSINON_SI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Expr()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DrAlgoParserALORS {
		{
			p.SetState(138)
			p.Match(DrAlgoParserALORS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DrAlgoParserNEWLINE {
		{
			p.SetState(141)
			p.Sep()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElsePartContext is an interface to support dynamic dispatch.
type IElsePartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SINON() antlr.TerminalNode
	Block() IBlockContext
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsElsePartContext differentiates from other interfaces.
	IsElsePartContext()
}

type ElsePartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsePartContext() *ElsePartContext {
	var p = new(ElsePartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_elsePart
	return p
}

func InitEmptyElsePartContext(p *ElsePartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_elsePart
}

func (*ElsePartContext) IsElsePartContext() {}

func NewElsePartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsePartContext {
	var p = new(ElsePartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_elsePart

	return p
}

func (s *ElsePartContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsePartContext) SINON() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserSINON, 0)
}

func (s *ElsePartContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElsePartContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *ElsePartContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *ElsePartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsePartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsePartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitElsePart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) ElsePart() (localctx IElsePartContext) {
	localctx = NewElsePartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DrAlgoParserRULE_elsePart)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(DrAlgoParserSINON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DrAlgoParserNEWLINE {
		{
			p.SetState(149)
			p.Sep()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(154)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TANT_QUE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext
	FIN_TANT_QUE() antlr.TerminalNode
	AllSep() []ISepContext
	Sep(i int) ISepContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) TANT_QUE() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserTANT_QUE, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) FIN_TANT_QUE() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserFIN_TANT_QUE, 0)
}

func (s *WhileStmtContext) AllSep() []ISepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISepContext); ok {
			len++
		}
	}

	tst := make([]ISepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISepContext); ok {
			tst[i] = t.(ISepContext)
			i++
		}
	}

	return tst
}

func (s *WhileStmtContext) Sep(i int) ISepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISepContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DrAlgoParserRULE_whileStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(DrAlgoParserTANT_QUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Expr()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DrAlgoParserNEWLINE {
		{
			p.SetState(158)
			p.Sep()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(163)
		p.Block()
	}
	{
		p.SetState(164)
		p.Match(DrAlgoParserFIN_TANT_QUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StmtList() IStmtListContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DrAlgoParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.StmtList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISepContext is an interface to support dynamic dispatch.
type ISepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEWLINE() antlr.TerminalNode

	// IsSepContext differentiates from other interfaces.
	IsSepContext()
}

type SepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySepContext() *SepContext {
	var p = new(SepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_sep
	return p
}

func InitEmptySepContext(p *SepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_sep
}

func (*SepContext) IsSepContext() {}

func NewSepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SepContext {
	var p = new(SepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_sep

	return p
}

func (s *SepContext) GetParser() antlr.Parser { return s.parser }

func (s *SepContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserNEWLINE, 0)
}

func (s *SepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitSep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Sep() (localctx ISepContext) {
	localctx = NewSepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DrAlgoParserRULE_sep)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(DrAlgoParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Equality() IEqualityContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Equality() IEqualityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DrAlgoParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Equality()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityContext is an interface to support dynamic dispatch.
type IEqualityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparison() []IComparisonContext
	Comparison(i int) IComparisonContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNEQ() []antlr.TerminalNode
	NEQ(i int) antlr.TerminalNode

	// IsEqualityContext differentiates from other interfaces.
	IsEqualityContext()
}

type EqualityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityContext() *EqualityContext {
	var p = new(EqualityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_equality
	return p
}

func InitEmptyEqualityContext(p *EqualityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_equality
}

func (*EqualityContext) IsEqualityContext() {}

func NewEqualityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityContext {
	var p = new(EqualityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_equality

	return p
}

func (s *EqualityContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityContext) AllComparison() []IComparisonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonContext); ok {
			len++
		}
	}

	tst := make([]IComparisonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonContext); ok {
			tst[i] = t.(IComparisonContext)
			i++
		}
	}

	return tst
}

func (s *EqualityContext) Comparison(i int) IComparisonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *EqualityContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserEQ)
}

func (s *EqualityContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserEQ, i)
}

func (s *EqualityContext) AllNEQ() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserNEQ)
}

func (s *EqualityContext) NEQ(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserNEQ, i)
}

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Equality() (localctx IEqualityContext) {
	localctx = NewEqualityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DrAlgoParserRULE_equality)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Comparison()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DrAlgoParserEQ || _la == DrAlgoParserNEQ {
		{
			p.SetState(173)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DrAlgoParserEQ || _la == DrAlgoParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(174)
			p.Comparison()
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAddition() []IAdditionContext
	Addition(i int) IAdditionContext
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllLTE() []antlr.TerminalNode
	LTE(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllGTE() []antlr.TerminalNode
	GTE(i int) antlr.TerminalNode

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_comparison
	return p
}

func InitEmptyComparisonContext(p *ComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_comparison
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllAddition() []IAdditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditionContext); ok {
			len++
		}
	}

	tst := make([]IAdditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditionContext); ok {
			tst[i] = t.(IAdditionContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) Addition(i int) IAdditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditionContext)
}

func (s *ComparisonContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserLT)
}

func (s *ComparisonContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserLT, i)
}

func (s *ComparisonContext) AllLTE() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserLTE)
}

func (s *ComparisonContext) LTE(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserLTE, i)
}

func (s *ComparisonContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserGT)
}

func (s *ComparisonContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserGT, i)
}

func (s *ComparisonContext) AllGTE() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserGTE)
}

func (s *ComparisonContext) GTE(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserGTE, i)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DrAlgoParserRULE_comparison)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Addition()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0 {
		{
			p.SetState(181)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(182)
			p.Addition()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditionContext is an interface to support dynamic dispatch.
type IAdditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplication() []IMultiplicationContext
	Multiplication(i int) IMultiplicationContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditionContext differentiates from other interfaces.
	IsAdditionContext()
}

type AdditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionContext() *AdditionContext {
	var p = new(AdditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_addition
	return p
}

func InitEmptyAdditionContext(p *AdditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_addition
}

func (*AdditionContext) IsAdditionContext() {}

func NewAdditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionContext {
	var p = new(AdditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_addition

	return p
}

func (s *AdditionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionContext) AllMultiplication() []IMultiplicationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicationContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicationContext); ok {
			tst[i] = t.(IMultiplicationContext)
			i++
		}
	}

	return tst
}

func (s *AdditionContext) Multiplication(i int) IMultiplicationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicationContext)
}

func (s *AdditionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserPLUS)
}

func (s *AdditionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserPLUS, i)
}

func (s *AdditionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserMINUS)
}

func (s *AdditionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserMINUS, i)
}

func (s *AdditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitAddition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Addition() (localctx IAdditionContext) {
	localctx = NewAdditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DrAlgoParserRULE_addition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Multiplication()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DrAlgoParserPLUS || _la == DrAlgoParserMINUS {
		{
			p.SetState(189)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DrAlgoParserPLUS || _la == DrAlgoParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(190)
			p.Multiplication()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicationContext is an interface to support dynamic dispatch.
type IMultiplicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPower() []IPowerContext
	Power(i int) IPowerContext
	AllMUL() []antlr.TerminalNode
	MUL(i int) antlr.TerminalNode
	AllDIV() []antlr.TerminalNode
	DIV(i int) antlr.TerminalNode
	AllMOD() []antlr.TerminalNode
	MOD(i int) antlr.TerminalNode

	// IsMultiplicationContext differentiates from other interfaces.
	IsMultiplicationContext()
}

type MultiplicationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicationContext() *MultiplicationContext {
	var p = new(MultiplicationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_multiplication
	return p
}

func InitEmptyMultiplicationContext(p *MultiplicationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_multiplication
}

func (*MultiplicationContext) IsMultiplicationContext() {}

func NewMultiplicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicationContext {
	var p = new(MultiplicationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_multiplication

	return p
}

func (s *MultiplicationContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicationContext) AllPower() []IPowerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPowerContext); ok {
			len++
		}
	}

	tst := make([]IPowerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPowerContext); ok {
			tst[i] = t.(IPowerContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicationContext) Power(i int) IPowerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerContext)
}

func (s *MultiplicationContext) AllMUL() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserMUL)
}

func (s *MultiplicationContext) MUL(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserMUL, i)
}

func (s *MultiplicationContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserDIV)
}

func (s *MultiplicationContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserDIV, i)
}

func (s *MultiplicationContext) AllMOD() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserMOD)
}

func (s *MultiplicationContext) MOD(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserMOD, i)
}

func (s *MultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Multiplication() (localctx IMultiplicationContext) {
	localctx = NewMultiplicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DrAlgoParserRULE_multiplication)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Power()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&114688) != 0 {
		{
			p.SetState(197)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&114688) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(198)
			p.Power()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPowerContext is an interface to support dynamic dispatch.
type IPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unary() IUnaryContext
	POW() antlr.TerminalNode
	Power() IPowerContext

	// IsPowerContext differentiates from other interfaces.
	IsPowerContext()
}

type PowerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerContext() *PowerContext {
	var p = new(PowerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_power
	return p
}

func InitEmptyPowerContext(p *PowerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_power
}

func (*PowerContext) IsPowerContext() {}

func NewPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerContext {
	var p = new(PowerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_power

	return p
}

func (s *PowerContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *PowerContext) POW() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserPOW, 0)
}

func (s *PowerContext) Power() IPowerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerContext)
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Power() (localctx IPowerContext) {
	localctx = NewPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DrAlgoParserRULE_power)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Unary()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DrAlgoParserPOW {
		{
			p.SetState(205)
			p.Match(DrAlgoParserPOW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Power()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unary() IUnaryContext
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	Primary() IPrimaryContext

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *UnaryContext) PLUS() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserPLUS, 0)
}

func (s *UnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserMINUS, 0)
}

func (s *UnaryContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DrAlgoParserRULE_unary)
	var _la int

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DrAlgoParserPLUS, DrAlgoParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DrAlgoParserPLUS || _la == DrAlgoParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(210)
			p.Unary()
		}

	case DrAlgoParserLPAREN, DrAlgoParserLBRACK, DrAlgoParserIDENT, DrAlgoParserFLOAT, DrAlgoParserINT, DrAlgoParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Primary()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListLiteral() IListLiteralContext
	Literal() ILiteralContext
	IDENT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) ListLiteral() IListLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListLiteralContext)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserIDENT, 0)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserLPAREN, 0)
}

func (s *PrimaryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserRPAREN, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DrAlgoParserRULE_primary)
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DrAlgoParserLBRACK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.ListLiteral()
		}

	case DrAlgoParserFLOAT, DrAlgoParserINT, DrAlgoParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Literal()
		}

	case DrAlgoParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Match(DrAlgoParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DrAlgoParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.Match(DrAlgoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Expr()
		}
		{
			p.SetState(219)
			p.Match(DrAlgoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListLiteralContext is an interface to support dynamic dispatch.
type IListLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListLiteralContext differentiates from other interfaces.
	IsListLiteralContext()
}

type ListLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListLiteralContext() *ListLiteralContext {
	var p = new(ListLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_listLiteral
	return p
}

func InitEmptyListLiteralContext(p *ListLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_listLiteral
}

func (*ListLiteralContext) IsListLiteralContext() {}

func NewListLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListLiteralContext {
	var p = new(ListLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_listLiteral

	return p
}

func (s *ListLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ListLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserLBRACK, 0)
}

func (s *ListLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserRBRACK, 0)
}

func (s *ListLiteralContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListLiteralContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DrAlgoParserCOMMA)
}

func (s *ListLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DrAlgoParserCOMMA, i)
}

func (s *ListLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitListLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) ListLiteral() (localctx IListLiteralContext) {
	localctx = NewListLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DrAlgoParserRULE_listLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(DrAlgoParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64760446976) != 0 {
		{
			p.SetState(224)
			p.Expr()
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DrAlgoParserCOMMA {
			{
				p.SetState(225)
				p.Match(DrAlgoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(226)
				p.Expr()
			}

			p.SetState(231)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(234)
		p.Match(DrAlgoParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	INT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DrAlgoParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrAlgoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserFLOAT, 0)
}

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserINT, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(DrAlgoParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DrAlgoVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DrAlgoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DrAlgoParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&60129542144) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
