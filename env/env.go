package env

import (
	_ "embed"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Env interface {
	List() ([]string, error)
}

type envImpl struct {
	typesPaths *[]string
}

// For mocking
var (
	parserParseFile = parser.ParseFile
)

//go:embed types.go
var typesGo string

// List collects string constants from the embedded types.go file, but only those with type Var
func (s *envImpl) List() ([]string, error) {
	var constants []string

	// Parse the Go source file from the embedded content
	fset := token.NewFileSet()
	node, err := parserParseFile(fset, "types.go", typesGo, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	// Inspect the AST and collect constants of type Var
	ast.Inspect(node, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.CONST {
			return true
		}

		for _, spec := range decl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			// Check if the type is explicitly "Var"
			if valueSpec.Type == nil {
				continue
			}

			ident, ok := valueSpec.Type.(*ast.Ident)
			if !ok || ident.Name != "Var" {
				continue
			}

			// Extract the string values
			for _, value := range valueSpec.Values {
				basicLit, ok := value.(*ast.BasicLit)
				if !ok || basicLit.Kind != token.STRING {
					continue
				}
				// Trim the quotes from the string value
				constants = append(constants, strings.Trim(basicLit.Value, "\""))
			}
		}
		return false
	})

	return constants, nil
}
