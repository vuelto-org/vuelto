//go:build windows || linux || darwin || freebsd || netbsd

/*
 * Copyright (C) 2026 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

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
