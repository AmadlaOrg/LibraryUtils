package env

import (
	_ "embed"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type IEnv interface {
	List() ([]string, error)
}

type SEnv struct {
	typesPaths *[]string
}

// For mocking
var (
	parserParseFile = parser.ParseFile
)

//go:embed types.go
var typesGo string

// List collects string constants from the embedded types.go file, but only those with type Var
func (s *SEnv) List() ([]string, error) {
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

/*
// List collects string constants from the embedded types.go file
func List() ([]string, error) {
	var constants []string

	// Parse the Go source file from the embedded content
	fset := token.NewFileSet()
	node, err := parserParseFile(fset, "types.go", typesGo, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	// Inspect the AST and collect string constants
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

/*import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

var filepathAbs = filepath.Abs

// List collects string constants from the types.go file in the env package
func List() ([]string, error) {
	// Get the absolute path of the current directory
	dir, err := filepathAbs(filepath.Dir("."))
	if err != nil {
		return nil, fmt.Errorf("failed to get the absolute path of the current directory: %v", err)
	}

	// Construct the full path to types.go
	filename := filepath.Join(dir, "env", "types.go")
	var constants []string

	// Parse the Go source file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %v", filename, err)
	}

	// Inspect the AST and collect string constants
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
}*/
