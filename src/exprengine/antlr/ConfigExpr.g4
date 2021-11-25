grammar ConfigExpr;  //表明此文件是词法分析器（lexer）定义文件,词法分析器名称要和文件保持一致

//定义token

chunk
    : block EOF
    ;
block
    : kvpair*
    ;
kvpair
    : key ASSIGN value
    ;
key
    : KEY   #FuncKey
    ;
value
    : (StringArg | SLICE | DICT) #FuncValue
    ;

SPACE: [ \r\n\t]+ -> skip;   //意思是遇到空格\r\n\t就扔到HIDDEN通道中，HIDDEN通道就是扔掉不要
COMMENT: '#' ~('\n'|'\r')* '\r'? '\n' ->skip;    //遇到注释就跳过

fragment DIGIT: [0-9];   //数字
Float: (DIGIT+)?'.'DIGIT+;  //无符号float，(DIGIT+)?表示至少匹配一个数字，但这个数字可有可无。然后.然后后面至少一个数字
FloatArg: '-'?Float;    //有符号float
IntArg: '-'?DIGIT+;     //有符号int
StringArg: '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';    //字符串必须为单引号'a'但支持单引号转义

ASSIGN: '=';
fragment COLON: ':';
fragment COMMA: ',';
fragment Dot: '.';
fragment UNDERLINE: '_';


ARG: (StringArg|IntArg|FloatArg)+;
SLICE: '[' (ARG? | (ARG COMMA)* ARG?) ']'; //[], [a], [a,b], [a,b,]
DICT: '\\'? '{' ~('\\'|'}')* '\\'? '}';
KEY: [A-Z] ([A-Z] | DIGIT | UNDERLINE)+;  //一个KEY必须是大写字母组成的


