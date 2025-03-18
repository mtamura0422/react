//go:generate go run .
//go:generate gofmt -w ../

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

// 構造体定義
type StructDef struct {
	Name   string
	Fields []FieldDef
}

// 構造体フィールド定義
type FieldDef struct {
	Name      string
	Type      string
	ArrayType string
}

// 配列のフィールド保存用
type FieldArray struct {
	Name string
	Type string
}

// ファイルをパースしてASTから必要な構造体情報を返却
func parseFirstStruct(fpath string) ([]StructDef, []FieldArray, error) {

	// ファイルの中身をast構造にする
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fpath, nil, parser.Mode(0))
	if err != nil {
		return nil, nil, err
	}

	list := []StructDef{}
	farray := []FieldArray{}

	ast.Inspect(f, func(n ast.Node) bool {

		switch x := n.(type) {
		case *ast.TypeSpec:

			sdef := StructDef{}

			// 対象が構造体
			st, ok := x.Type.(*ast.StructType)
			if !ok {
				return true
			}
			sdef.Name = x.Name.Name

			for _, fld := range st.Fields.List {
				// bun.BaseModel `bun:"table:recipes"` などはNamesがない
				if fld.Names == nil {
					continue
				}

				arrayType := ""
				switch fldType := fld.Type.(type) {
				case *ast.ArrayType:

					var fldTypeBuf bytes.Buffer
					err := printer.Fprint(&fldTypeBuf, fset, fldType.Elt)
					if err != nil {
						log.Fatalf("failed printing %s", err)
					}

					arrayType = fldTypeBuf.String()
					farray = append(farray, FieldArray{Name: fld.Names[0].Name, Type: arrayType})

				}

				var typeNameBuf bytes.Buffer
				err := printer.Fprint(&typeNameBuf, fset, fld.Type)
				if err != nil {
					log.Fatalf("failed printing %s", err)
				}

				sdef.Fields = append(
					sdef.Fields,
					FieldDef{Name: fld.Names[0].Name,
						Type:      typeNameBuf.String(),
						ArrayType: arrayType,
					})
			}
			list = append(list, sdef)
		}
		return true

	})

	return list, farray, nil
}

// マッパーファイルを生成
func createMapperFile(outputFilePath string, def StructDef, fa []FieldArray) error {
	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.Must(template.ParseFiles("./mappter.tmpl"))
	data := map[string]interface{}{
		"ModelName":   def.Name, // "Comment"
		"Fields":      def.Fields,
		"ArrayFields": fa,
	}

	if err := t.Execute(file, data); err != nil {
		return err
	}
	fmt.Println(outputFilePath + " is generated.")
	return nil
}

// エントリーポイント
func main() {

	inputFilePaths := []string{"../recipe_model.go", "../recipe_material_model.go"}

	for _, filepath := range inputFilePaths {
		list, farray, err := parseFirstStruct(filepath)
		if err != nil || len(list) == 0 {
			fmt.Fprintf(os.Stderr, "model parse faild.\n: %s", err)
			os.Exit(1)
		}

		outputFilePath := "../" + strings.Split(strings.Split(filepath, "/")[1], ".")[0] + "_mapper.gen.go"
		err = createMapperFile(outputFilePath, list[0], farray)
		if err != nil {
			fmt.Fprintf(os.Stderr, "code generate failed.\n: %s", err)
			os.Exit(1)
		}
	}
}
