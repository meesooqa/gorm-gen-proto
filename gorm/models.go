package gorm

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseTypes struct {
	gorm.Model
	UintField     uint
	Uint8Field    uint8
	Uint16Field   uint16
	Uint32Field   uint32
	Uint64Field   uint64
	IntField      int
	Int8Field     int8
	Int16Field    int16
	Int32Field    int32
	Int64Field    int64
	Float32Field  float32
	Float64Field  float64
	BoolField     bool
	StringField   string
	DateTimeField time.Time

	OptionalUintField     *uint
	OptionalUint8Field    *uint8
	OptionalUint16Field   *uint16
	OptionalUint32Field   *uint32
	OptionalUint64Field   *uint64
	OptionalIntField      *int
	OptionalInt8Field     *int8
	OptionalInt16Field    *int16
	OptionalInt32Field    *int32
	OptionalInt64Field    *int64
	OptionalFloat32Field  *float32
	OptionalFloat64Field  *float64
	OptionalBoolField     *bool
	OptionalStringField   *string
	OptionalDateTimeField *time.Time
}

type SetTypes struct {
	gorm.Model
	ByteSliceField   []byte
	StringSliceField []string
	IntSliceField    []int
}

type StructTypes struct {
	gorm.Model
	ItemField          BaseTypes
	OptionalItemField  *BaseTypes
	ItemsField         []BaseTypes
	OptionalItemsField []*BaseTypes
}

type SpecialTypes struct {
	gorm.Model
	SqlNullStringField sql.NullString
	SqlNullInt64Field  sql.NullInt64
	SqlNullTimeField   sql.NullTime
	DeletedAtField     gorm.DeletedAt
	MapField           map[string]interface{} `gorm:"type:jsonb"`
}
