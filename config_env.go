package config

import (
	"fmt"
	"os"

	"github.com/bopher/caster"
	"github.com/bopher/utils"
	"github.com/joho/godotenv"
)

type envConfig struct {
	Files []string
	data  map[string]interface{}
}

func (envConfig) err(format string, args ...interface{}) error {
	return utils.TaggedError([]string{"EnvConfig"}, format, args...)
}

func (this *envConfig) Load() error {
	if this.data == nil {
		this.data = make(map[string]interface{})
	}

	if err := godotenv.Overload(this.Files...); err != nil {
		return this.err(err.Error())
	}

	for k, v := range this.data {
		if err := this.Set(k, v); err != nil {
			return err
		}
	}

	return nil
}

func (this *envConfig) Set(key string, value interface{}) error {
	this.data[key] = value
	if err := os.Setenv(key, fmt.Sprintf("%v", value)); err != nil {
		return this.err(err.Error())
	}
	return nil
}

func (envConfig) Get(key string) interface{} {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return nil
}

func (envConfig) Exists(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

func (this envConfig) Cast(key string) caster.Caster {
	return caster.NewCaster(this.Get(key))
}
