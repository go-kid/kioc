package scan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var testFile = `package test

// Starter @Component
type Starter struct {
}

// New @Component
func New() *Starter {
	return &Starter{}
}

func main() {
	err:=errors.New()
}
`

func Test_analyseFile(t *testing.T) {
	fset := token.NewFileSet()
	var f *ast.File
	f, err := parser.ParseFile(fset, "", testFile, parser.ParseComments)
	assert.NoError(t, err)
	registers, err := analyseToken(f)
	assert.NoError(t, err)
	fmt.Println(registers[1])
	assert.Equal(t, "Starter", registers[0].Name)
	assert.Equal(t, "Component", registers[0].Group)
	assert.Equal(t, "type", registers[0].Kind)
	assert.Equal(t, "test", registers[0].Pkg)

	assert.Equal(t, "New", registers[1].Name)
	assert.Equal(t, "Component", registers[1].Group)
	assert.Equal(t, "func", registers[1].Kind)
	assert.Equal(t, "test", registers[1].Pkg)
}
