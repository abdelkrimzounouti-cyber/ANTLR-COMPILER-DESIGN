lexer grammar GoLexer;

ELSE    : 'else';
IF      : 'if';
WHILE   : 'while';
FUNC    : 'func';
VAR     : 'var';
RETURN  : 'return';

LE      : '<=';
GE      : '>=';
EQ      : '==';
NE      : '!=';
LT      : '<';
GT      : '>';

PLUS    : '+';
MINUS   : '-';
STAR    : '*';
SLASH   : '/';
ASSIGN  : '=';

LPAREN  : '(';
RPAREN  : ')';
LBRACE  : '{';
RBRACE  : '}';
COMMA   : ',';
SEMI    : ';';

INT : [0-9]+;
ID  : [a-zA-Z_][a-zA-Z_0-9]*;

LINE_COMMENT  : '//' ~[\r\n]* -> skip;
BLOCK_COMMENT : '/*' .*? '*/' -> skip;

WS : [ \t\r\n]+ -> skip;