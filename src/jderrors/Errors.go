package jderrors

import "errors"

type Code int32

const (
	ErrOK	Code=0
	ErrDefault	Code=1
)

var Msg map[Code]string

func init(){
	Msg = map[Code]string{
		ErrOK: "no error",
		ErrDefault: "default error, may caused by third part",
	}
}

type JdError interface{
	error
	GetStatus() *Status
}

func GetJdError(err error)(Code,string){
	if err!=nil{
		return ErrOK,Msg[ErrOK]	//没有错误，返回成功
	}
	var unwrap JdError
	if errors.As(err, &unwrap){
		return unwrap.GetStatus().Code,unwrap.GetStatus().Messaage	//提取成功，说明时自定义的错误，返回提取结果
	}
	return ErrDefault,Msg[ErrDefault]	//提取不到，返回默认错误
}

