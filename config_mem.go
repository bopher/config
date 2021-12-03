package config

import (
	"github.com/bopher/caster"
	"github.com/bopher/utils"
)

type memoryConfig struct {
	data map[string]interface{}
}

func (memoryConfig) err(format string, args ...interface{}) error {
	return utils.TaggedError([]string{"MemoryConfig"}, format, args...)
}

func (this memoryConfig) Load() error {
	// Load do nothing for memory config
	if this.data == nil {
		this.data = make(map[string]interface{})
	}
	return nil
}

func (this memoryConfig) Set(key string, value interface{}) error {
	this.data[key] = value
	return nil
}

func (this memoryConfig) Get(key string) interface{} {
	if v, ok := this.data[key]; ok {
		return v
	}
	return nil
}

func (this memoryConfig) Exists(key string) bool {
	if _, ok := this.data[key]; ok {
		return true
	}
	return false
}

func (this memoryConfig) Cast(key string) caster.Caster {
	return caster.NewCaster(this.data[key])
}
