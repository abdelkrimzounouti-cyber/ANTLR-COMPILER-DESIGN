
parser grammar GoParser;

options {
    tokenVocab = GoLexer;
}

// 1. Programme complet
program
    : declaration* EOF
    ;

// 2. Déclarations globales
declaration
    : variableDeclaration
    | functionDeclaration
    ;

// 3. Déclaration de variable
variableDeclaration
    : VAR ID ASSIGN expression SEMI
    ;

// 4. Déclaration de fonction
functionDeclaration
    : FUNC ID LPAREN parameterList? RPAREN block
    ;

parameterList
    : ID (COMMA ID)*
    ;

// 5. Bloc d'instructions
block
    : LBRACE statement* RBRACE
    ;

// 6. Instructions
statement
    : variableDeclaration
    | assignment
    | ifStatement
    | whileStatement
    | returnStatement
    | functionCall SEMI
    ;

// 7. Affectation
assignment
    : ID ASSIGN expression SEMI
    ;

// 8. Condition
ifStatement
    : IF LPAREN expression RPAREN block
      (ELSE block)?
    ;

// 9. Boucle
whileStatement
    : WHILE LPAREN expression RPAREN block
    ;

// 10. Retour de fonction
returnStatement
    : RETURN expression SEMI
    ;

// 11. Appel de fonction
functionCall
    : ID LPAREN argumentList? RPAREN
    ;

argumentList
    : expression (COMMA expression)*
    ;

// 12. Expressions avec priorités
expression
    : LPAREN expression RPAREN
    | functionCall
    | ID
    | INT
    | expression (STAR | SLASH) expression
    | expression (PLUS | MINUS) expression
    | expression (LT | GT | LE | GE | EQ | NE) expression
    ;
