package sprite

import (
	"errors"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/harrisonthorne/chicky-chicky-go/render"
	"github.com/harrisonthorne/chicky-chicky-go/textures"
	"github.com/harrisonthorne/chicky-chicky-go/utils"

	mgl "github.com/go-gl/mathgl/mgl32"
)

// Sprite is an image that animates.
type Sprite struct {
	texture  uint32
	frames   int
	uvCoords []float32

	currentFrame    float32
	secondsPerFrame float32

	sizeMatrix     mgl.Mat4
	positionMatrix mgl.Mat4
	matrix         mgl.Mat4

    pixelWidth, pixelHeight int
}

var planeVAO, planeVBO uint32

var modelAttrLocation, textureUniform int32
var	program uint32

// InitGL initializes plane vao and vbo
func InitGL() {
	planeVAO, planeVBO = utils.NewTextureVAO(render.TextureProgram(), &planeVertices)

	modelAttrLocation = render.TextureProgram().Locations.ModelMatrixLocation()
	textureUniform = gl.GetUniformLocation(program, gl.Str("tex\x00"))
    program = render.TextureProgram().ID()
}

// New creates a new sprite and returns it
func New(spritePath string, frames int, secondsPerFrame float32) (s *Sprite, err error) {
	s = new(Sprite)

	if frames <= 0 {
		frames = 1
		if secondsPerFrame <= 0 {
			secondsPerFrame = 1
		}
	} else if secondsPerFrame <= 0 {
		return nil, errors.New("secondsPerFrame cannot be less than or equal to 0 if frames is greater than 0")
	}

	// open the sprite file
	spriteFile, err := os.Open(spritePath)
	if err != nil {
		return
	}

	// assign the sprite texture
	s.texture, err = textures.New(spriteFile)
	if err != nil {
		return
	}

	// initialize the rest of the fields
	s.frames = frames
	s.secondsPerFrame = secondsPerFrame

	return
}

// MustNew is like NewSprite, but panics if there's an
// error
func MustNew(spritePath string, frames int, secondsPerFrame float32) *Sprite {
	sprite, err := New(spritePath, frames, secondsPerFrame)

	if err != nil {
		panic(err)
	}

	sprite.positionMatrix = mgl.Ident4()
	sprite.sizeMatrix = mgl.Ident4()
	sprite.updateMatrix()


	return sprite
}

// Animate animates the Sprite.
func (s *Sprite) Animate(delta float32) {
	// if one frame or less, animation doesn't matte,
	if s.frames <= 1 {
		return
	}
	s.currentFrame += delta / s.secondsPerFrame
	for s.currentFrame >= float32(s.frames) {
		s.currentFrame -= float32(s.frames)
	}
}

// SetSize sets the size of the sprite.
func (s *Sprite) SetSize(width, height float32) {
	s.sizeMatrix = mgl.Scale3D(width, height, 1)
	s.updateMatrix()
}

func (s *Sprite) PixelWidth() int {
    return s.pixelWidth
}

func (s *Sprite) PixelHeight() int {
    return s.pixelHeight
}

// SetPosition sets the position of the sprite.
func (s *Sprite) SetPosition(x, y, z float32) {
	s.positionMatrix = mgl.Translate3D(x, y, z)
	s.updateMatrix()
}

func (s *Sprite) updateMatrix() {
	s.matrix = s.positionMatrix.Mul4(s.sizeMatrix)
}

// Render renders the sprite onto the screen.
func (s *Sprite) Render(c *render.Camera) {
	program := render.TextureProgram().ID()
	gl.UseProgram(program)

	c.SetProgramAttributes(render.TextureProgram())

	modelAttrLocation := render.TextureProgram().Locations.ModelMatrixLocation()
	gl.UniformMatrix4fv(modelAttrLocation, 1, false, &s.matrix[0])

	gl.Uniform1i(render.TextureProgram().Locations.SpriteFramesLocation(), int32(s.frames))
	gl.Uniform1i(render.TextureProgram().Locations.SpriteCurrentFrameLocation(), int32(s.currentFrame)) // number bound here must match the active texture

	textureUniform := render.TextureProgram().Locations.TextureLocation()
	gl.Uniform1i(textureUniform, 0) // number bound here must match the active texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, s.texture)

	gl.BindVertexArray(planeVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 2*3)
}

var planeVertices = []float32{
	// first triangle
	-0.5, 0.5, 0, 0, 0,
	-0.5, -0.5, 0, 0, 1,
	0.5, -0.5, 0, 1, 1,

	// second triangle
	-0.5, 0.5, 0, 0, 0,
	0.5, -0.5, 0, 1, 1,
	0.5, 0.5, 0, 1, 0,
}
