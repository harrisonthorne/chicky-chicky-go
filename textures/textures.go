package textures

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png" // imported for png support for textures
	"io"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// LazyTexture will dynamically load a texture the instant
// its ID is attempted to be accessed.
type LazyTexture struct {
	path string
	id   uint32
}

func NewLazyTexture(imagePath string) *LazyTexture {
	return &LazyTexture{path: imagePath}
}

// ID returns the texture ID, creating the texture if
// necessary
func (l *LazyTexture) ID() uint32 {
	if l.id == 0 {
		imageFile, err := os.Open(l.path)
		if err != nil {
			return 0
		}

		l.id, err = New(imageFile)
		if err != nil {
			panic(err)
		}
	}

	return l.id
}

// New creates a new texture with the image data from
// the reader.
func New(imageReader io.Reader) (uint32, error) {
	img, _, err := image.Decode(imageReader)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Over)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}

// Bind binds the provided texture for use with OpenGL.
func Bind(texture uint32) {
	gl.BindTexture(gl.TEXTURE_2D, texture)
}
