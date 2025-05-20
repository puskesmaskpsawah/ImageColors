package services

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"log"
	"strconv"
)

func GetPixelColor(base64Image string, pixelX int, pixelY int) string {
	imgBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		log.Println("Base64 decode error:", err)
		return ""
	}

	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Println("PNG decode error:", err)
		return ""
	}
	bounds := img.Bounds()
	log.Println("Image size:", bounds.Dx(), "x", bounds.Dy())

	x, y := pixelX, pixelY
	if x < bounds.Dx() && y < bounds.Dy() {

		r, g, b, a := img.At(x, y).RGBA()
		log.Printf("Pixel (%d,%d) color - R:%d G:%d B:%d A:%d\n", x, y, r>>8, g>>8, b>>8, a>>8)

		rStr := strconv.Itoa(int(r >> 8))
		gStr := strconv.Itoa(int(g >> 8))
		bStr := strconv.Itoa(int(b >> 8))
		aStr := strconv.Itoa(int(a >> 8))
		return "R:" + rStr + " G:" + gStr + " B:" + bStr + " A:" + aStr
	} else {
		log.Println("Pixel out of bounds!")
		return ""
	}
}
