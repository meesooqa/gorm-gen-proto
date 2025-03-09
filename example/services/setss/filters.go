package setss

import (
	"gorm.io/gorm"

	"github.com/meesooqa/gorm-gen-proto/example/services"
)

func ExampleFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
