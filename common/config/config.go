package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Read(path string) (*JanctionConf, error) {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config filepath: [%s], error: %s", path, err)
	}
	return &JanctionConf{
		v:        v,
		filepath: path,
	}, nil
}

func ReadFromJson(content string) (*JanctionConf, error) {
	v := viper.New()
	reader := strings.NewReader(content)
	v.SetConfigType("json")
	v.ReadConfig(reader)
	return &JanctionConf{
		v:        v,
		filepath: "",
	}, nil
}

func NewForUnitTest(paths ...string) *JanctionConf {
	var path string
	if len(paths) > 0 {
		path = paths[0]
	}

	v := viper.New()
	v.SetConfigFile(path)
	_ = v.ReadInConfig()
	return &JanctionConf{
		v:        v,
		filepath: path,
	}
}

type JanctionConf struct {
	v        *viper.Viper
	filepath string
}

func (b *JanctionConf) Set(path string, value interface{}) {
	b.v.Set(path, value)
}

func (b *JanctionConf) WatchConfig() {
	b.v.WatchConfig()
}

func (b *JanctionConf) ConfPath() string {
	return b.filepath
}

func (b *JanctionConf) GetInt(path string, def int) int {
	if b.v.IsSet(path) {
		return b.v.GetInt(path)
	}
	return def
}

func (b *JanctionConf) GetBool(path string, def bool) bool {
	if b.v.IsSet(path) {
		return b.v.GetBool(path)
	}
	return def
}

func (b *JanctionConf) GetFloat(path string, def float64) float64 {
	if b.v.IsSet(path) {
		return b.v.GetFloat64(path)
	}
	return def
}

func (b *JanctionConf) GetStrSlice(path string, def []string) []string {
	if b.v.IsSet(path) {
		return b.v.GetStringSlice(path)
	}
	return def
}

func (b *JanctionConf) MustGetStrSlice(path string) []string {
	if !b.v.IsSet(path) {
		panic(any(fmt.Sprintf("cannot get config in %s %s", b.filepath, path)))
	}
	return b.v.GetStringSlice(path)
}

func (b *JanctionConf) GetString(path string, def string) string {
	if b.v.IsSet(path) {
		return b.v.GetString(path)
	}
	return def
}

func (b *JanctionConf) GetIntSlice(path string, def []int) []int {
	if b.v.IsSet(path) {
		return b.v.GetIntSlice(path)
	}
	return def
}

func (b *JanctionConf) Get(path string, def interface{}) interface{} {
	if b.v.IsSet(path) {
		return b.v.Get(path)
	}
	return def
}

func (b *JanctionConf) MustGetInt(path string) int {
	if !b.v.IsSet(path) {
		panic(any(fmt.Sprintf("cannot get config in %s %s", b.filepath, path)))
	}
	return b.v.GetInt(path)
}

func (b *JanctionConf) MustGetFloat(path string) float64 {
	if !b.v.IsSet(path) {
		panic(any(fmt.Sprintf("cannot get config in %s %s", b.filepath, path)))
	}
	return b.v.GetFloat64(path)
}

func (b *JanctionConf) MustGetString(path string) string {
	if !b.v.IsSet(path) {
		panic(any(fmt.Sprintf("cannot get config in %s %s", b.filepath, path)))
	}
	return b.v.GetString(path)
}

func (b *JanctionConf) GetMap(path string) map[string]interface{} {
	if !b.v.IsSet(path) {
		panic(any(fmt.Sprintf("cannot get config in %s %s", b.filepath, path)))
	}
	return b.v.GetStringMap(path)
}
