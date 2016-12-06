package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := []byte(`package main

  import "fmt"
  import "koding/multiconfig"

  func main() {
    var sum int = 3 + 2
    fmt.Println(sum)
  }`)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "main", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	found := false
	// =======================================================
	visitorFunc := func(n ast.Node) bool {
		binaryExpr, ok := n.(*ast.ImportSpec)
		if !ok {
			return true
		}
		if binaryExpr.Path.Value == `"koding/multiconfig"` {
			fmt.Printf("Found binary expression at: %d:%d\n",
				fset.Position(binaryExpr.Pos()).Line,
				fset.Position(binaryExpr.Pos()).Column)
			fmt.Println(binaryExpr.Path.Value)
			found = true
			return true
		}
		return false
	}
	ast.Inspect(node, visitorFunc)
	if found {
		visitorFunc := func(n ast.Node) bool {
			gg, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			fmt.Println(gg.Fun)
			return false
		}
		ast.Inspect(node, visitorFunc)
	}
	// =======================================================

	// =======================================================
	// for _, decl := range node.Decls {
	// 	// search for main function
	// 	fn, ok := decl.(*ast.FuncDecl)
	// 	if !ok {
	// 		continue
	// 	}

	// 	// inside the main function body
	// 	for _, stmt := range fn.Body.List {
	// 		// search through statements for a declaration ...
	// 		declStmt, ok := stmt.(*ast.DeclStmt)
	// 		if !ok {
	// 			continue
	// 		}
	// 		// continue with declStmt.Decl
	// 		genDecl, ok := declStmt.Decl.(*ast.GenDecl)
	// 		if !ok {
	// 			continue
	// 		}

	// 		// declarations can have multiple specs,
	// 		// search for a valuespec
	// 		for _, spec := range genDecl.Specs {
	// 			valueSpec, ok := spec.(*ast.ValueSpec)
	// 			if !ok {
	// 				continue
	// 			}
	// 			for _, expr := range valueSpec.Values {
	// 				// search for a binary expr
	// 				binaryExpr, ok := expr.(*ast.BinaryExpr)
	// 				if !ok {
	// 					continue
	// 				}
	// 				// found it!
	// 				fmt.Printf("Found binary expression at: %d:%d\n",
	// 					fset.Position(binaryExpr.Pos()).Line,
	// 					fset.Position(binaryExpr.Pos()).Column,
	// 				)
	// 			}
	// 		}
	// 	}
	// }
	// =======================================================
}
