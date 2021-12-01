package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type envConfig struct {
	Files []string
	data  map[string]interface{}
}

func (this *envConfig) err(pattern string, params ...interface{}) error {
	return errors.New(fmt.Sprintf("[EnvConfig]: "+pattern, params...))
}

// Load configurations
func (this *envConfig) Load() error {
	if err := godotenv.Overload(this.Files...); err != nil {
		return this.err(err.Error())
	}
	if this.data == nil {
		this.data = make(map[string]interface{})
	} else {
		for k, v := range this.data {
			os.Setenv(k, fmt.Sprintf("%v", v))
		}
	}
	return nil
}

// Set configuration
// return error if driver not support set or error happend
func (this *envConfig) Set(key string, value interface{}) error {
	this.data[key] = value
	if err := os.Setenv(key, fmt.Sprintf("%v", value)); err != nil {
		return this.err(err.Error())
	}
	return nil
}

// Get configuration
func (envConfig) Get(key string) interface{} {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return nil
}

// Exists check if config item exists
func (envConfig) Exists(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

// Bool parse dependency as boolean
func (envConfig) Bool(key string, fallback bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseBool(val); err == nil {
			return res
		}
	}
	return fallback
}

// Int parse dependency as int
func (envConfig) Int(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int(res)
		}
	}
	return fallback
}

// Int8 parse dependency as int8
func (envConfig) Int8(key string, fallback int8) int8 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int8(res)
		}
	}
	return fallback
}

// Int16 parse dependency as int16
func (envConfig) Int16(key string, fallback int16) int16 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int16(res)
		}
	}
	return fallback
}

// Int32 parse dependency as int32
func (envConfig) Int32(key string, fallback int32) int32 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int32(res)
		}
	}
	return fallback
}

// Int64 parse dependency as int64
func (envConfig) Int64(key string, fallback int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int64(res)
		}
	}
	return fallback
}

// UInt parse dependency as uint
func (envConfig) UInt(key string, fallback uint) uint {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint(res)
		}
	}
	return fallback
}

// UInt8 parse dependency as uint8
func (envConfig) UInt8(key string, fallback uint8) uint8 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint8(res)
		}
	}
	return fallback
}

// UInt16 parse dependency as uint16
func (envConfig) UInt16(key string, fallback uint16) uint16 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint16(res)
		}
	}
	return fallback
}

// UInt32 parse dependency as uint32
func (envConfig) UInt32(key string, fallback uint32) uint32 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint32(res)
		}
	}
	return fallback
}

// UInt64 parse dependency as uint64
func (envConfig) UInt64(key string, fallback uint64) uint64 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseUint(val, 10, 64); err == nil {
			return uint64(res)
		}
	}
	return fallback
}

// Float32 parse dependency as float64
func (envConfig) Float32(key string, fallback float32) float32 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseFloat(val, 64); err == nil {
			return float32(res)
		}
	}
	return fallback
}

// Float64 parse dependency as float64
func (envConfig) Float64(key string, fallback float64) float64 {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.ParseFloat(val, 64); err == nil {
			return res
		}
	}
	return fallback
}

// String parse dependency as string
func (envConfig) String(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
