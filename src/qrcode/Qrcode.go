package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"log"
)

type QrCode struct{

}

func NewQrCode() *QrCode {
	return &QrCode{}
}

func (q *QrCode)GenerateQrCode(url string,x,y int)*barcode.Barcode{
	qrCode,err :=qr.Encode(url,qr.M,qr.Auto)
	if err!=nil{
		log.Println("[error] Generate QRCode Error")
	}
	qrCode,err = barcode.Scale(qrCode,200,200)
	if err!=nil{
		log.Println("[error] Scale QRCode Error")
	}
	return &qrCode
}
