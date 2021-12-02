package config

import "github.com/bopher/utils"

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

func (this memoryConfig) BoolE(key string) (bool, error) {
	if v, err := utils.CastBoolE(this.data[key]); err != nil {
		return false, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) Bool(key string, fallback bool) bool {
	if v, err := this.BoolE(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) IntE(key string) (int, error) {
	if v, err := utils.CastIntE(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) Int(key string, fallback int) int {
	if v, err := this.IntE(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) Int8E(key string) (int8, error) {
	if v, err := utils.CastInt8E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) Int8(key string, fallback int8) int8 {
	if v, err := this.Int8E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) Int16E(key string) (int16, error) {
	if v, err := utils.CastInt16E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) Int16(key string, fallback int16) int16 {
	if v, err := this.Int16E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) Int32E(key string) (int32, error) {
	if v, err := utils.CastInt32E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) Int32(key string, fallback int32) int32 {
	if v, err := this.Int32E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) Int64E(key string) (int64, error) {
	if v, err := utils.CastInt64E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) Int64(key string, fallback int64) int64 {
	if v, err := this.Int64E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) UIntE(key string) (uint, error) {
	if v, err := utils.CastUIntE(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) UInt(key string, fallback uint) uint {
	if v, err := this.UIntE(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) UInt8E(key string) (uint8, error) {
	if v, err := utils.CastUInt8E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) UInt8(key string, fallback uint8) uint8 {
	if v, err := this.UInt8E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) UInt16E(key string) (uint16, error) {
	if v, err := utils.CastUInt16E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) UInt16(key string, fallback uint16) uint16 {
	if v, err := this.UInt16E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) UInt32E(key string) (uint32, error) {
	if v, err := utils.CastUInt32E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) UInt32(key string, fallback uint32) uint32 {
	if v, err := this.UInt32E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) UInt64E(key string) (uint64, error) {
	if v, err := utils.CastUInt64E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) UInt64(key string, fallback uint64) uint64 {
	if v, err := this.UInt64E(key); err == nil {
		return v
	}
	return fallback
}

func (this memoryConfig) Float64E(key string) (float64, error) {
	if v, err := utils.CastFloat64E(this.data[key]); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this memoryConfig) Float64(key string, fallback float64) float64 {
	if v, err := this.Float64E(key); err == nil {
		return v
	}
	return fallback
}
func (this memoryConfig) StringE(key string) (string, error) {
	if v, ok := this.data[key].(string); ok {
		return v, nil
	}
	return "", this.err("failed cast %s as string!", key)
}

func (this memoryConfig) String(key string, fallback string) string {
	if v, err := this.StringE(key); err == nil {
		return v
	}
	return fallback
}
