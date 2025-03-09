package main

import (
	"gorm-gen-proto-01/example/models"
	"gorm-gen-proto-01/gen"
)

func GetGormList() []*gen.GormForTmpl {
	return []*gen.GormForTmpl{
		gen.NewGormForTmpl(models.BaseTypes{}, "basepb", "bases"),
		gen.NewGormForTmpl(models.SetTypes{}, "setpb", "sets"),
		gen.NewGormForTmpl(models.StructTypes{}, "structpb", "structs"),
		gen.NewGormForTmpl(models.SpecialTypes{}, "specialpb", "specials"),
	}
}
