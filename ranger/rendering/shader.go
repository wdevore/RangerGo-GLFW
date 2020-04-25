// Package rendering defines features of shaders.
package rendering

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
)

// Shader represents a shader program
type Shader struct {
	vertexCode   string
	fragmentCode string

	vertexSrc   string
	fragmentSrc string

	program uint32 // GLuint
}

// NewShader creates a blank shader. You must call Load before that shader is valid.
func NewShader(vertexSrc, fragmentSrc string) *Shader {
	s := new(Shader)
	s.vertexSrc = vertexSrc
	s.fragmentSrc = fragmentSrc
	return s
}

// Load reads and compiles shader programs
func (s *Shader) Load() error {

	var err error
	s.vertexCode, s.fragmentCode, err = fetch(s.vertexSrc, s.fragmentSrc)
	if err != nil {
		return err
	}

	s.program, err = newProgram(s.vertexCode, s.fragmentCode)
	if err != nil {
		return err
	}

	return nil
}

// Use activates program
func (s *Shader) Use() {
	gl.UseProgram(s.program)
}

func fetch(vertexSrc, fragmentSrc string) (vCode, fCode string, err error) {
	// Vertex source -----------------------------------------------
	filePath := fmt.Sprintf("./assets/%s", vertexSrc)

	var bytes []byte
	bytes, err = ioutil.ReadFile(filePath)

	if err != nil {
		return "", "", err
	}

	vCode = string(bytes)

	// Fragment source -----------------------------------------------
	filePath = fmt.Sprintf("./assets/%s", fragmentSrc)

	bytes, err = ioutil.ReadFile(filePath)

	if err != nil {
		return "", "", err
	}

	fCode = string(bytes)

	return
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compile(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compile(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil

}

func compile(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
