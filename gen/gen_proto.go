package gen

import (
	"log/slog"
	"reflect"
	"unicode"
	"unicode/utf8"

	"gorm-gen-proto-01/config"
)

type GormForTmpl struct {
	model    any
	Package  string
	Endpoint string
}

func NewGormForTmpl(model any, pkg, ep string) *GormForTmpl {
	return &GormForTmpl{
		model:    model,
		Package:  pkg,
		Endpoint: ep,
	}
}

type Proto3TmplDataBuilder struct {
	logger        *slog.Logger
	data          *Proto3TmplData
	dataProvider  *GormProtoDataProvider
	gormModel     any
	minModelField int
}

type Proto3TmplData struct {
	ApiVersion   string
	Package      string
	Imports      []string
	ModelFields  []ProtoField
	FilterFields []ProtoField
	Endpoint     string
}

type ProtoField struct {
	Name    string
	Type    string
	Index   int
	Options string
}

func NewProto3TmplDataBuilder(logger *slog.Logger, conf *config.Conf, gormModelForBuilder *GormForTmpl) *Proto3TmplDataBuilder {
	return &Proto3TmplDataBuilder{
		logger:        logger,
		dataProvider:  NewGormProtoDataProvider(logger, conf.System),
		gormModel:     gormModelForBuilder.model,
		minModelField: 1,
		data: &Proto3TmplData{
			ApiVersion: "v1",
			Package:    gormModelForBuilder.Package,
			Endpoint:   gormModelForBuilder.Endpoint,
		},
	}
}

func (o *Proto3TmplDataBuilder) ProvideData() (*Proto3TmplData, error) {
	var err error
	err = o.fillModelFields(o.gormModel)
	if err != nil {
		return nil, err
	}
	err = o.fillImports()
	if err != nil {
		return nil, err
	}
	err = o.fillGetListRequestFilters()
	if err != nil {
		return nil, err
	}
	return o.data, nil
}

func (o *Proto3TmplDataBuilder) fillGetListRequestFilters() error {
	if len(o.data.ModelFields) > 0 {
		// gt, lt
		//for _, field := range o.ModelFields {
		//}
	}
	return nil
}

func (o *Proto3TmplDataBuilder) fillImports() error {
	importsMap, err := o.dataProvider.GetProtoImportsMap()
	if err != nil {
		return err
	}
	o.data.Imports = append(o.data.Imports, "google/api/annotations.proto")
	if len(o.data.ModelFields) > 0 {
		for _, field := range o.data.ModelFields {
			pkg, exists := importsMap[field.Type]
			if exists {
				o.data.Imports = append(o.data.Imports, pkg)
			}
		}
	}
	return nil
}

func (o *Proto3TmplDataBuilder) fillModelFields(gormModel interface{}) error {
	typeMap, err := o.dataProvider.GetGormProtoMap()
	if err != nil {
		return err
	}
	t := reflect.TypeOf(gormModel)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		pbType, exists := typeMap[field.Type.String()]
		if exists {
			if pbType != "" {
				// optional repeated map
				// optional repeated: repeated field is inherently optional
				// TODO Options
				o.data.ModelFields = append(o.data.ModelFields, ProtoField{
					Name:    o.toLowerFirstLetter(field.Name),
					Type:    pbType,
					Index:   i + o.minModelField,
					Options: "",
				})
			} else {
				o.logger.Error("empty type", slog.String("type", field.Type.String()), slog.String("name", field.Name))
			}
		} else {
			o.logger.Error("non-existed type", slog.String("type", field.Type.String()), slog.String("name", field.Name))
		}
	}
	return nil
}

func (o *Proto3TmplDataBuilder) toLowerFirstLetter(s string) string {
	if s == "" {
		return s
	}
	firstRune, size := utf8.DecodeRuneInString(s)
	loweredRune := unicode.ToLower(firstRune)
	return string(loweredRune) + s[size:]
}
