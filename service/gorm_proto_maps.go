package service

import (
	"gorm-gen-proto-01/config"
)

type GormProtoDataProvider struct {
	conf *config.SystemConfig
}

func NewGormProtoDataProvider(conf *config.SystemConfig) *GormProtoDataProvider {
	return &GormProtoDataProvider{
		conf: conf,
	}
}

// GetGormProtoMap gorm.Model field type => proto field type
func (o *GormProtoDataProvider) GetGormProtoMap() (map[string]string, error) {
	store, err := NewStore(o.conf.PathMaps + "/type.json")
	if err != nil {
		return nil, err
	}
	return store.Data, nil
}

// GetProtoImportsMap proto field type => proto package
func (o *GormProtoDataProvider) GetProtoImportsMap() (map[string]string, error) {
	store, err := NewStore(o.conf.PathMaps + "/import.json")
	if err != nil {
		return nil, err
	}
	return store.Data, nil
}
