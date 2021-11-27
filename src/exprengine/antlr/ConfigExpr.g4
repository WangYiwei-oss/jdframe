grammar ConfigExpr;  //表明此文件是词法分析器（lexer）定义文件,词法分析器名称要和文件保持一致

//定义token

config
    : block EOF
    ;
block
    : kvpair*
    ;
kvpair
    :  key '=' value
    ;
key
    : KEY   #FuncKey
    ;
object
    : '{' pair ( ',' pair )* '}'
    | '{' '}'
    ;
slice
    : '[' value ( ',' value )* ']'
    | '[' ']'
    ;
pair: mapkey ':' value
    ;

mapkey: StringArg  #MapKey;

value
    : StringArg #StringValue
    | IntArg    #IntValue
    | FloatArg  #FloatValue
    | Bool      #BoolValue
    | 'null'    #NullValue
    | object    #ObjectValue
    | slice     #SliceValue
    ;

SPACE: [ \r\n\t]+ -> skip;   //意思是遇到空格\r\n\t就扔到HIDDEN通道中，HIDDEN通道就是扔掉不要
COMMENT: '#' ~('\n'|'\r')* '\r'? '\n' ->skip;    //遇到注释就跳过
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
fragment ESC: '\\' (["\\bfnrt] | UNICODE);
fragment DIGIT: [0-9];   //数字
fragment Float: (DIGIT+)? '.' DIGIT+;  //无符号float，(DIGIT+)?表示至少匹配一个数字，但这个数字可有可无。然后.然后后面至少一个数字
FloatArg: '-'?Float;    //有符号float
IntArg: '0' | '-'? [1-9] DIGIT*;     //有符号int,可以有正负号，但除了0不能从0开始
StringArg: '"' (ESC | ~["\\])* '"';    //字符串必须为单引号'a'但支持单引号转义
KEY: [A-Z] ([A-Z] | DIGIT | '_')+;  //一个KEY必须是大写字母组成的
Bool: ('true' | 'false');


