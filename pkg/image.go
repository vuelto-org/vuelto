package pkg

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"os"

	"github.com/vuelto-org/vuelto/internal/gl"
)

type Image struct {
	texture       uint32
	x, y          float32
	width, height float32
}

var ImageArray []uint32

// Loads a new image and returns a Image struct. Can be later drawn using the Draw() method
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

	var textureID uint32
	gl.GenTextures(1, textureID)

	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, rgba.Rect.Size().X, rgba.Rect.Size().Y, 0, gl.RGBA, gl.UNSIGNED_BYTE, rgba.Pix)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	ImageArray = append(ImageArray, textureID)

	return Image{
		texture: textureID,
		x:       x,
		y:       y,
		width:   width,
		height:  height,
	}
}

// Draws the image that's loaded before.
func (img Image) Draw() {
	gl.BindTexture(gl.TEXTURE_2D, img.texture)
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

func cleanTex() {
	for _, i := range ImageArray {
		gl.DeleteTextures(1, &i)
	}
}
