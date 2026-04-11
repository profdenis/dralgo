package antlr_test

import (
	"dralgo/internal/antlr"
	"testing"

	antlr_lib "github.com/antlr4-go/antlr/v4"
)

type errorListener struct {
	*antlr_lib.DefaultErrorListener
	errors []string
}

func (l *errorListener) SyntaxError(recognizer antlr_lib.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr_lib.RecognitionException) {
	l.errors = append(l.errors, msg)
}

func TestParser(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Basic Assignment",
			input:   "x = 10\n",
			wantErr: false,
		},
		{
			name:    "Multiple Statements",
			input:   "x = 10\ny = 20\nAfficher x + y\n",
			wantErr: false,
		},
		{
			name:    "If Statement",
			input:   "Si x > 10 Alors\n  Afficher \"Plus grand\"\nFin Si\n",
			wantErr: false,
		},
		{
			name:    "If Else Statement",
			input:   "Si x > 10 Alors\n  Afficher \"Plus grand\"\nSinon\n  Afficher \"Plus petit\"\nFin Si\n",
			wantErr: false,
		},
		{
			name:    "While Loop",
			input:   "Tant Que x < 10\n  x = x + 1\nFin Tant Que\n",
			wantErr: false,
		},
		{
			name:    "For Loop",
			input:   "Pour i = 0; i < 10; i++\n  Afficher i\nFin Pour\n",
			wantErr: false,
		},
		{
			name:    "Read Statements",
			input:   "Lire x\nLire Entier y\n",
			wantErr: false,
		},
		{
			name:    "Complex Expressions",
			input:   "z = (x + y) * 2 ** 3 / 4 % 5\n",
			wantErr: false,
		},
		{
			name:    "List Literal",
			input:   "l = [1, 2, 3]\n",
			wantErr: false,
		},
		{
			name:    "For Each Loop",
			input:   "Pour Chaque x Dans [1, 2, 3]\n  Afficher x\nFin Pour\n",
			wantErr: false,
		},
		{
			name:    "Do While Loop",
			input:   "Fait\n  x = x + 1\nTant Que x < 10\n",
			wantErr: false,
		},
		{
			name:    "Nested Structures",
			input:   "Si x > 0 Alors\n  Tant Que x < 10\n    x = x + 1\n    Si x == 5 Alors\n      Afficher \"Milieu\"\n    Fin Si\n  Fin Tant Que\nFin Si\n",
			wantErr: false,
		},
		{
			name:    "Empty Program",
			input:   "\n",
			wantErr: false,
		},
		{
			name:    "Invalid Syntax",
			input:   "x = \n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputStream := antlr_lib.NewInputStream(tt.input)
			lexer := antlr.NewDrAlgoLexer(inputStream)
			tokenStream := antlr_lib.NewCommonTokenStream(lexer, antlr_lib.TokenDefaultChannel)
			parser := antlr.NewDrAlgoParser(tokenStream)

			errListener := &errorListener{
				DefaultErrorListener: &antlr_lib.DefaultErrorListener{},
			}
			parser.RemoveErrorListeners()
			parser.AddErrorListener(errListener)

			parser.Program()

			if (len(errListener.errors) > 0) != tt.wantErr {
				t.Errorf("TestParser() %s: errors = %v, wantErr %v", tt.name, errListener.errors, tt.wantErr)
			}
		})
	}
}
