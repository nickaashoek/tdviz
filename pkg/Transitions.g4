grammar Transitions;

AND: '&';
OR: '|';
NOT: '!';
IMPL: '->';
DOUBLIMPL: '<->';
LPAREN: '(';
RPAREN: ')';
START: [A-Za-z];
END: [A-Za-z]'\'';
TRUE: 'TRUE';
FALSE: 'FALSE';
WHITESPACE: [ \r\n\t]+ -> skip;

start: expression EOF;

expression
    : LPAREN expression RPAREN #nestedExpression
    | NOT expression #notExpression
    | left=expression OP=operator right=expression #opExpression
    | identifier #identifierExpression
    | bl #boolExpression
    ;

operator 
    : AND | OR  | IMPL | DOUBLIMPL;    

identifier
    : START | END;

bl
    : TRUE | FALSE;