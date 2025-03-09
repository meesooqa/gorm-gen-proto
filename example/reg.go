package example

import (
	"github.com/meesooqa/gorm-gen-proto/example/models"
	"github.com/meesooqa/gorm-gen-proto/gen"
	"github.com/meesooqa/gorm-gen-proto/reg"
)

func init() {
	reg.RegisterGormData([]*gen.GormForTmpl{
		gen.NewGormForTmpl(models.BaseTypes{}, "basepb", "bases"),
		gen.NewGormForTmpl(models.SetTypes{}, "setpb", "sets"),
		gen.NewGormForTmpl(models.StructTypes{}, "structpb", "structs"),
		gen.NewGormForTmpl(models.SpecialTypes{}, "specialpb", "specials"),
	})

	ImportPbPrefix := "github.com/meesooqa/gorm-gen-proto/example/proto"
	importModels := "github.com/meesooqa/gorm-gen-proto/example/models"
	importServices := "github.com/meesooqa/gorm-gen-proto/example/services"
	reg.RegisterSsData([]*gen.SsTmplData{{
		Package:        "basess",
		DbModel:        "models.BaseTypes",
		ImportPb:       ImportPbPrefix + "/basepb",
		ImportServices: importServices,
		ImportModels:   importModels,
	}, {
		Package:        "setss",
		DbModel:        "models.SetTypes",
		ImportPb:       ImportPbPrefix + "/setpb",
		ImportServices: importServices,
		ImportModels:   importModels,
	}, {
		Package:        "structss",
		DbModel:        "models.StructTypes",
		ImportPb:       ImportPbPrefix + "/structpb",
		ImportServices: importServices,
		ImportModels:   importModels,
	}, {
		Package:        "specialss",
		DbModel:        "models.SpecialTypes",
		ImportPb:       ImportPbPrefix + "/specialpb",
		ImportServices: importServices,
		ImportModels:   importModels,
	}})
}
