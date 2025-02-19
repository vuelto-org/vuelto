package processing

import (
	"github.com/disintegration/imaging"

	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
)

func ConvertNRGBAtoRGBA(nrgba *image.NRGBA) *image.RGBA {
	rgba := image.NewRGBA(nrgba.Bounds())
	draw.Draw(rgba, rgba.Bounds(), nrgba, nrgba.Bounds().Min, draw.Src)
	return rgba
}

func Blur(image *image.RGBA, ratio float64) *image.RGBA {
	return ConvertNRGBAtoRGBA(imaging.Blur(image, ratio))
}

func Contrast(image *image.RGBA, ratio float64) *image.RGBA {
	return ConvertNRGBAtoRGBA(imaging.AdjustContrast(image, ratio))
}

func Sharpen(image *image.RGBA, ratio float64) *image.RGBA {
	return ConvertNRGBAtoRGBA(imaging.Sharpen(image, ratio))
}

func Invert(image *image.RGBA) *image.RGBA {
	return ConvertNRGBAtoRGBA(imaging.Invert(image))
}
