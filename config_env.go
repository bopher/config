package config

import (
	"fmt"
	"os"

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
	if err := godotenv.Overload(this.Files...); err != nil {
		return this.err(err.Error())
	}
	if this.data == nil {
		this.data = make(map[string]interface{})
	} else {
		for k, v := range this.data {
			if err := this.Set(k, v); err != nil {
				return err
			}
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

func (this envConfig) BoolE(key string) (bool, error) {
	if v, err := utils.CastBoolE(this.Get(key)); err != nil {
		return false, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) Bool(key string, fallback bool) bool {
	if v, err := this.BoolE(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) IntE(key string) (int, error) {
	if v, err := utils.CastIntE(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) Int(key string, fallback int) int {
	if v, err := this.IntE(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) Int8E(key string) (int8, error) {
	if v, err := utils.CastInt8E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) Int8(key string, fallback int8) int8 {
	if v, err := this.Int8E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) Int16E(key string) (int16, error) {
	if v, err := utils.CastInt16E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) Int16(key string, fallback int16) int16 {
	if v, err := this.Int16E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) Int32E(key string) (int32, error) {
	if v, err := utils.CastInt32E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) Int32(key string, fallback int32) int32 {
	if v, err := this.Int32E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) Int64E(key string) (int64, error) {
	if v, err := utils.CastInt64E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) Int64(key string, fallback int64) int64 {
	if v, err := this.Int64E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) UIntE(key string) (uint, error) {
	if v, err := utils.CastUIntE(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) UInt(key string, fallback uint) uint {
	if v, err := this.UIntE(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) UInt8E(key string) (uint8, error) {
	if v, err := utils.CastUInt8E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) UInt8(key string, fallback uint8) uint8 {
	if v, err := this.UInt8E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) UInt16E(key string) (uint16, error) {
	if v, err := utils.CastUInt16E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) UInt16(key string, fallback uint16) uint16 {
	if v, err := this.UInt16E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) UInt32E(key string) (uint32, error) {
	if v, err := utils.CastUInt32E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) UInt32(key string, fallback uint32) uint32 {
	if v, err := this.UInt32E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) UInt64E(key string) (uint64, error) {
	if v, err := utils.CastUInt64E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) UInt64(key string, fallback uint64) uint64 {
	if v, err := this.UInt64E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) Float64E(key string) (float64, error) {
	if v, err := utils.CastFloat64E(this.Get(key)); err != nil {
		return 0, this.err(err.Error())
	} else {
		return v, nil
	}
}

func (this envConfig) Float64(key string, fallback float64) float64 {
	if v, err := this.Float64E(key); err == nil {
		return v
	}
	return fallback
}

func (this envConfig) StringE(key string) (string, error) {
	if val, ok := os.LookupEnv(key); ok {
		return val, nil
	}
	return "", this.err("failed cast %s as string!", key)
}

func (this envConfig) String(key string, fallback string) string {
	if v, err := this.StringE(key); err == nil {
		return v
	}
	return fallback
}
