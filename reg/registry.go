package reg

import "github.com/meesooqa/gorm-gen-proto/gen"

var registry []*gen.GormForTmpl

func Register(models []*gen.GormForTmpl) {
	registry = append(registry, models...)
}

func GetRegistry() []*gen.GormForTmpl {
	return registry
}
