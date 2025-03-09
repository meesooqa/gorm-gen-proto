package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Conf from config yml
type Conf struct {
	System *SystemConfig `yaml:"system"`
}

// SystemConfig is the configuration for App
type SystemConfig struct {
	PathMaps  string `yaml:"path_maps"`
	PathTmpl  string `yaml:"path_tmpl"`
	ProtoRoot string `yaml:"proto_root"`
}

// Load config from file
func Load(fname string) (res *Conf, err error) {
	res = &Conf{}
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}
