package main

import (
	"fmt"
	parser "github.com/WangYiwei-oss/jdframe/src/exprengine/configexpr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct{
	*parser.BaseConfigExprListener
}

func (c *calcListener)ExitFuncKey(ctx *parser.FuncKeyContext) {
	fmt.Println(ctx.GetText())
}

func main(){
	is,_ := antlr.NewFileStream("./Settings.conf")
	lexer := parser.NewConfigExprLexer(is)
	ts := antlr.NewCommonTokenStream(lexer,antlr.TokenDefaultChannel)
	p:=parser.NewConfigExprParser(ts)
	tree := p.Chunk()
	antlr.ParseTreeWalkerDefault.Walk(&calcListener{},tree)
	//jdft.NewJdft().
	//	Beans(jdft.NewGormAdapter()).
	//	Mount("v1",testcontroller.NewUserController()).
	//	Mount("v2",testcontroller.NewUserController()).
	//Launch()
}
