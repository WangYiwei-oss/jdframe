package configparser

import (
	parser "github.com/WangYiwei-oss/jdframe/src/exprengine/configexpr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	json "github.com/json-iterator/go"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

//TODO antlr分析配置文件中，应该区分出来map键的不同类型，而不是只能[string]interface
var GlobalSettings map[string]interface{}

func init() {
	GlobalSettings = make(map[string]interface{})
	wd, _ := os.Getwd()
	is, _ := antlr.NewFileStream(path.Join(wd, "Settings.conf"))
	lexer := parser.NewConfigExprLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewConfigExprParser(ts)
	tree := p.Config()
	antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, tree)
}

type calcListener struct {
	*parser.BaseConfigExprListener
	key string
}

func (c *calcListener) EnterFuncKey(ctx *parser.FuncKeyContext) {
	c.key = ctx.GetText()
}

func (c *calcListener) ExitIntValue(ctx *parser.IntValueContext) {
	if c.key != "" {
		text := ctx.GetText()
		i, err := strconv.Atoi(text)
		if err != nil {
			log.Fatalln("解析配置文件错误，转换int错误", err)
		}
		GlobalSettings[c.key] = i
		c.key = ""
	}
}

func (c *calcListener) ExitFloatValue(ctx *parser.FloatValueContext) {
	if c.key != "" {
		text := ctx.GetText()
		f, err := strconv.ParseFloat(text, 32)
		if err != nil {
			log.Fatalln("解析配置文件错误，转换float错误", err)
		}
		GlobalSettings[c.key] = f
		c.key = ""
	}
}

func (c *calcListener) ExitBoolValue(ctx *parser.BoolValueContext) {
	if c.key != "" {
		text := ctx.GetText()
		if text == `"true"` {
			GlobalSettings[c.key] = true
		} else {
			GlobalSettings[c.key] = false
		}
		c.key = ""
	}
}

func (c *calcListener) ExitNullValue(ctx *parser.NullValueContext) {
	if c.key != "" {
		GlobalSettings[c.key] = nil
		c.key = ""
	}
}

func (c *calcListener) ExitStringValue(ctx *parser.StringValueContext) {
	if c.key != "" {
		text := ctx.GetText()
		GlobalSettings[c.key] = strings.Trim(text, `"`)
		c.key = ""
	}
}

func (c *calcListener) EnterSliceValue(ctx *parser.SliceValueContext) {
	if c.key != "" {
		text := ctx.GetText()
		obj := make([]interface{}, 0)
		err := json.Unmarshal([]byte(text), &obj)
		if err != nil {
			log.Fatalln("slice配置解析错误", err)
		}
		GlobalSettings[c.key] = obj
		c.key = ""
	}
}

func (c *calcListener) EnterObjectValue(ctx *parser.ObjectValueContext) {
	if c.key != "" {
		text := ctx.GetText()
		obj := make(map[string]interface{})
		err := json.Unmarshal([]byte(text), &obj)
		if err != nil {
			log.Fatalln("map配置解析错误", err)
		}
		GlobalSettings[c.key] = obj
		c.key = ""
	}
}

func (c *calcListener) EnterMapKey(ctx *parser.MapKeyContext) {
}
