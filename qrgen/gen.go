package qrgen

import (
	"bytes"
	"image"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// 画像データを受け取り（byteデータ）、Goで画像データとして操作可能なImageオブジェクトに変換
func CreateImage(b []byte) (image.Image, error) {
	// pngやfaviconなどの画像データはバイト形式で表現されるデータ
	// Goのパッケージに読ませて、加工や出力等できる状態にする
	return png.Decode(bytes.NewReader(b))
}

func GenQRCode(url string, width, height int) (barcode.Barcode, error) {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	return barcode.Scale(qrCode, width, height)
}
