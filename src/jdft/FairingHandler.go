package jdft

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var fairingHandler *FairingHandler
var fairingOnce sync.Once

type FairingHandler struct {
	fairings []Fairing
}

func getFairingHandler() *FairingHandler {
	fairingOnce.Do(func() {
		fairingHandler = &FairingHandler{}
	})
	return fairingHandler
}

func (f *FairingHandler) AddFairing(fs ...Fairing) {
	if fs != nil && len(fs) > 0 {
		f.fairings = append(f.fairings, fs...)
	}
}

func (f *FairingHandler) before(ctx *gin.Context) {
	for _, f := range f.fairings {
		err := f.OnRequest(ctx)
		if err != nil {
			Throw(err.Error(), 400, ctx)
		}
	}
}

func (f *FairingHandler) after(ctx *gin.Context, ret interface{}) interface{} {
	var result = ret
	for _, f := range f.fairings {
		r, err := f.OnResponse(result)
		if err != nil {
			Throw(err.Error(), 400, ctx)
		}
		result = r
	}
	return result
}

func (f *FairingHandler) handlerFairing(responder Responder, ctx *gin.Context) interface{} {
	f.before(ctx)
	return nil
}
