package main

import (
	"github.com/meesooqa/gorm-gen-proto/example/models"
	"github.com/meesooqa/gorm-gen-proto/gen"
)

func GetGormList() []*gen.GormForTmpl {
	return []*gen.GormForTmpl{
		gen.NewGormForTmpl(models.BaseTypes{}, "basepb", "bases"),
		gen.NewGormForTmpl(models.SetTypes{}, "setpb", "sets"),
		gen.NewGormForTmpl(models.StructTypes{}, "structpb", "structs"),
		gen.NewGormForTmpl(models.SpecialTypes{}, "specialpb", "specials"),
	}
}
