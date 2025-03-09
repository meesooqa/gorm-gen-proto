package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"reflect"
	"text/template"

	"gorm-gen-proto-01/config"
	"gorm-gen-proto-01/example/models"
	"gorm-gen-proto-01/gen"
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

	gg := []*gen.GormForTmpl{
		gen.NewGormForTmpl(models.BaseTypes{}, "basepb", "bases"),
		gen.NewGormForTmpl(models.SetTypes{}, "setpb", "sets"),
		gen.NewGormForTmpl(models.StructTypes{}, "structpb", "structs"),
		gen.NewGormForTmpl(models.SpecialTypes{}, "specialpb", "specials"),
	}

	// generate proto files
	pg := gen.NewProto3Generator(conf.System, templates)
	for _, g := range gg {
		err := pg.Run(logger, g)
		if err != nil {
			log.Fatal(err)
		}
	}
	// generate go files by `protoc`
	pe := gen.NewProtocExecutor()
	for _, g := range gg {
		protoFilePath := g.GetProtoFilePath(conf.System)
		err := pe.Run(conf.System.ProtoRoot, protoFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}

	logger.Info("end")
}
