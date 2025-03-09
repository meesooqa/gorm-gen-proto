package example

import (
	"github.com/meesooqa/gorm-gen-proto/example/models"
	"github.com/meesooqa/gorm-gen-proto/gen"
	"github.com/meesooqa/gorm-gen-proto/reg"
)

func init() {
	reg.Register([]*gen.GormForTmpl{
		gen.NewGormForTmpl(models.BaseTypes{}, "basepb", "bases"),
		gen.NewGormForTmpl(models.SetTypes{}, "setpb", "sets"),
		gen.NewGormForTmpl(models.StructTypes{}, "structpb", "structs"),
		gen.NewGormForTmpl(models.SpecialTypes{}, "specialpb", "specials"),
	})
}
