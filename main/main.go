package main

import (
	"fmt"
	"os"

	grammar "compiler/generated/grammar"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors int
}

func (e *ErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	exception antlr.RecognitionException,
) {
	e.Errors++
	fmt.Printf("Erreur ligne %d, colonne %d : %s\n",
		line, column, msg)
}

func runLexer(source string) {
	fmt.Println("=== Analyse lexicale ===")

	input := antlr.NewInputStream(source)
	lexer := grammar.NewGoLexer(input)

	errors := &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errors)

	for {
		token := lexer.NextToken()

		if token.GetTokenType() == antlr.TokenEOF {
			break
		}

		tokenType := token.GetTokenType()
		tokenName := "INCONNU"

		if tokenType >= 0 &&
			int(tokenType) < len(lexer.SymbolicNames) {
			tokenName = lexer.SymbolicNames[tokenType]
		}

		fmt.Printf("%-20s %q\n", tokenName, token.GetText())
	}

	if errors.Errors > 0 {
		fmt.Printf("Erreurs lexicales : %d\n", errors.Errors)
	}
}

func runParser(source string) int {
	fmt.Println("\n=== Analyse syntaxique ===")

	input := antlr.NewInputStream(source)
	lexer := grammar.NewGoLexer(input)

	errors := &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errors)

	tokens := antlr.NewCommonTokenStream(
		lexer,
		antlr.TokenDefaultChannel,
	)

	parser := grammar.NewGoParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errors)

	parser.Program()

	if errors.Errors > 0 {
		fmt.Printf("Échec : %d erreur(s) détectée(s).\n",
			errors.Errors)
		return errors.Errors
	}

	fmt.Println("Syntaxe correcte !")
	return 0
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run ./main <fichier.pcg>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		os.Exit(1)
	}

	source := string(data)

	runLexer(source)

	if runParser(source) > 0 {
		os.Exit(1)
	}
}
