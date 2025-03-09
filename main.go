package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	"gorm-gen-proto-01/config"
	"gorm-gen-proto-01/gen"
	"gorm-gen-proto-01/gorm"
)

var conf *config.Conf
var templates *template.Template

func init() {
	c, err := config.Load("var/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	conf = c

	funcMap := template.FuncMap{
		"notEmpty": func(arr interface{}) bool {
			v := reflect.ValueOf(arr)
			return v.Kind() == reflect.Slice && v.Len() > 0
		},
	}
	templates = template.Must(
		template.New("").
			Funcs(funcMap).
			ParseGlob(fmt.Sprintf("%s/*.tmpl", conf.System.PathTmpl)),
	)
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Info("begin")

	// generate proto files
	gg := []*gen.GormForTmpl{
		gen.NewGormForTmpl(gorm.BaseTypes{}, "basepb", "bases"),
		gen.NewGormForTmpl(gorm.SetTypes{}, "setpb", "sets"),
		gen.NewGormForTmpl(gorm.StructTypes{}, "structpb", "structs"),
		gen.NewGormForTmpl(gorm.SpecialTypes{}, "specialpb", "specials"),
	}
	for _, g := range gg {
		generateProto(logger, g)
	}
	// generate go files by `protoc`
	pe := gen.NewProtocExecutor()
	for _, g := range gg {
		pe.Run(g.Package, g.Package)
	}

	logger.Info("end")
}

func generateProto(logger *slog.Logger, gm *gen.GormForTmpl) {
	// template data
	tmplData := gen.NewProto3TmplDataBuilder(logger, conf, gm)
	data, err := tmplData.ProvideData()
	if err != nil {
		log.Fatalf("tmpl data obtainig: %s", err.Error())
	}
	// create file
	filePath := fmt.Sprintf("pb/proto/%s/%s.proto", gm.Package, gm.Package)
	dir := filepath.Dir(filePath)
	if err = os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("dir creating: %s", err.Error())
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("file creating: %s", err.Error())
	}
	// write
	err = templates.ExecuteTemplate(file, "proto3.tmpl", data)
	if err != nil {
		log.Fatalf("tmpl executing: %s", err.Error())
	}
}
