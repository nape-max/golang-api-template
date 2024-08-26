package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	conf "maxnap/platform/internal/config"

	"github.com/BurntSushi/toml"
)

var handlerTemplate = template.Must(template.New("").Parse(`package handler

import (
	"context"

	"maxnap/platform/internal/generated/schema"
)

func (s Server) {{ .HandlerName }}(
	ctx context.Context,
	request schema.{{ .RequestObjectType }},
) (schema.{{ .ReturnObjectType }}, error) {
	return schema.{{ .HandlerName }}200JSONResponse{
	}, nil
}
`))

type Param struct {
	HandlerName       string
	RequestObjectType string
	ReturnObjectType  string
}

func main() {
	_, err := os.Stat("config.toml")
	if err != nil {
		panic(fmt.Errorf("cannot receive stat of config file: %w", err))
	}

	var cfg conf.ConfigGenerator
	_, err = toml.DecodeFile("config.toml", &cfg)
	if err != nil {
		panic(fmt.Errorf("cannot decode config to struct: %w", err))
	}

	// Цель генерации передаётся переменной окружения
	pathToGeneratedServer := cfg.Generator.PathToGeneratedServer
	if pathToGeneratedServer == "" {
		log.Fatal("GOFILE must be set")
	}

	// Разбираем целевой файл в AST
	astInFile, err := parser.ParseFile(
		token.NewFileSet(),
		pathToGeneratedServer,
		nil,
		// Нас интересуют комментарии
		parser.ParseComments,
	)
	if err != nil {
		panic(fmt.Errorf("parse file error: %w", err))
	}

	generators := make([]Param, 0)

	// traverse all tokens
	ast.Inspect(astInFile, func(n ast.Node) bool {
		switch t := n.(type) {
		// find variable declarations
		case *ast.TypeSpec:
			if t.Name.Name != "StrictServerInterface" {
				return false
			}

			for _, m := range t.Type.(*ast.InterfaceType).Methods.List {
				name := m.Names[0].Name
				param := Param{
					HandlerName: name,
				}

				for _, p := range m.Type.(*ast.FuncType).Params.List {
					if p.Names[0].Name != "ctx" {
						param.RequestObjectType = fmt.Sprintf("%s", p.Type)
					}
				}

				for _, r := range m.Type.(*ast.FuncType).Results.List {
					if fmt.Sprintf("%s", r.Type) != "error" {
						param.ReturnObjectType = fmt.Sprintf("%s", r.Type)
					}
				}

				generators = append(generators, param)
			}
		}

		return true
	})

	for _, g := range generators {
		handlerFileName := ToSnakeCase(g.HandlerName) + ".go"
		handlersDir := cfg.Generator.PathToHandlers

		entries, err := os.ReadDir(handlersDir)
		if err != nil {
			panic(fmt.Errorf("cannot read dir: %w", err))
		}

		var found bool
		for _, e := range entries {
			if e.Name() == handlerFileName {
				found = true
				break
			}
		}

		if found {
			fmt.Printf("\t%-14s %s\n", "Already Exist:", handlerFileName)
			continue
		}

		var buf bytes.Buffer
		err = handlerTemplate.Execute(&buf, g)
		if err != nil {
			panic(fmt.Errorf("execute template: %v", err))
		}

		err = os.WriteFile(handlersDir+handlerFileName, buf.Bytes(), 0o644)
		if err != nil {
			fmt.Printf("\t%-14s %s: %s\n", "Error:", handlerFileName, err.Error())
			continue
		}

		fmt.Printf("\t%-14s %s\n", "Generated:", handlerFileName)
	}
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
