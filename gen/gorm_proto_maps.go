package gen

import (
	"log/slog"

	"gorm-gen-proto-01/config"
)

type GormProtoDataProvider struct {
	logger *slog.Logger
	conf   *config.SystemConfig
}

func NewGormProtoDataProvider(logger *slog.Logger, conf *config.SystemConfig) *GormProtoDataProvider {
	return &GormProtoDataProvider{
		logger: logger,
		conf:   conf,
	}
}

// GetGormProtoMap gorm.Model field type => proto field type
func (o *GormProtoDataProvider) GetGormProtoMap() (map[string]string, error) {
	store, err := NewStore(o.conf.PathMaps + "/type.json")
	if err != nil {
		return nil, err
	}
	return store.data, nil
}

// GetProtoImportsMap proto field type => proto package
func (o *GormProtoDataProvider) GetProtoImportsMap() (map[string]string, error) {
	store, err := NewStore(o.conf.PathMaps + "/import.json")
	if err != nil {
		return nil, err
	}
	return store.data, nil
}

// load datamap from file
//func (o *GormProtoDataProvider) load(fname string) (map[string]string, error) {
//	res := make(map[string]string)
//	data, err := os.ReadFile(fname)
//	if err != nil {
//		return nil, err
//	}
//	if err := yaml.Unmarshal(data, &res); err != nil {
//		return nil, err
//	}
//	return res, nil
//}
