package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	c, err := Load("testdata/config.yml")
	require.NoError(t, err)
	assert.IsType(t, &SystemConfig{}, c.System)
	assert.Equal(t, "var/data/maps", c.System.PathMaps)
	assert.Equal(t, "var/data/templates", c.System.PathTmpl)
	assert.Equal(t, "pb/proto", c.System.ProtoRoot)
}

func TestLoadConfigNotFoundFile(t *testing.T) {
	r, err := Load("/tmp/ded02288-9203-403e-86d8-ada2f937a37b.txt")
	assert.Nil(t, r)
	assert.EqualError(t, err, "open /tmp/ded02288-9203-403e-86d8-ada2f937a37b.txt: no such file or directory")
}

func TestLoadConfigInvalidYaml(t *testing.T) {
	r, err := Load("testdata/file.txt")
	assert.Nil(t, r)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `Not Yaml` into config.Conf")
}
