package configs

import (
	"github.com/WangYiwei-oss/jdframe/src/configinjector"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"log"
	"sync"
)

func init() {
	configinjector.RegisterBean("QRCODE", NewQrCode)
}

var doOnceQR sync.Once

type QrCode struct {
}

var QrC *QrCode

func NewQrCode() *QrCode {
	doOnceQR.Do(func() {
		QrC = &QrCode{}
	})
	return QrC
}

func (q *QrCode) GenerateQrCode(url string, x, y int) *barcode.Barcode {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		log.Println("[error] Generate QRCode Error")
	}
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		log.Println("[error] Scale QRCode Error")
	}
	return &qrCode
}
