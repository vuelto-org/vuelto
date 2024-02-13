package src

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"os"

	"github.com/go-gl/gl/v2.1/gl"
	cgl "github.com/vuelto-org/vuelto/internal/gl"
)

type Image struct {
	texture       uint
	x, y          float32
	width, height float32
}

var ImageArray []Image

func (r *Renderer2D) LoadImage(imagePath string, x, y, width, height float32) Image {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalln("Failed to open image: ", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln("Failed to decode image: ", err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Over)

	var vueltoImage Image
	cgl.GenTextures(1, vueltoImage.texture)
	cgl.BindTexture(gl.TEXTURE_2D, vueltoImage.texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	vueltoImage.x = x
	vueltoImage.y = y
	vueltoImage.width = width
	vueltoImage.height = height

	ImageArray = append(ImageArray, vueltoImage)

	return vueltoImage
}

func (img Image) Draw() {
	cgl.BindTexture(gl.TEXTURE_2D, img.texture)
	defer gl.BindTexture(gl.TEXTURE_2D, 0)

	gl.Begin(gl.QUADS)
	defer gl.End()

	gl.TexCoord2f(0.0, 0.0)
	gl.Vertex2f(img.x, img.y)
	gl.TexCoord2f(1.0, 0.0)
	gl.Vertex2f(img.x+img.width, img.y)
	gl.TexCoord2f(1.0, 1.0)
	gl.Vertex2f(img.x+img.width, img.y+img.height)
	gl.TexCoord2f(0.0, 1.0)
	gl.Vertex2f(img.x, img.y+img.height)
}
