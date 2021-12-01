package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

type jsonConfig struct {
	Files []string
	json  string
	data  map[string]interface{}
}

func (this *jsonConfig) err(pattern string, params ...interface{}) error {
	return errors.New(fmt.Sprintf("[JsonConfig]: "+pattern, params...))
}

func (this *jsonConfig) fetch(key string) (interface{}, bool, bool) {
	if v, ok := this.data[key]; ok {
		return v, false, true
	}

	if val := gjson.Get(this.json, key); val.Exists() {
		return val, true, true
	}

	return nil, false, false
}

// Load configurations
func (this *jsonConfig) Load() error {
	if this.data == nil {
		this.data = make(map[string]interface{})
	}

	contents := make([]string, 0)
	for _, f := range this.Files {
		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return this.err(err.Error())
		}
		content := string(bytes)
		if !gjson.Valid(content) {
			return errors.New(fmt.Sprintf("%s content is invalid!", f))
		}

		fileName := filepath.Base(f)
		fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

		if len(this.Files) > 1 {
			contents = append(contents, `"`+fileName+`":`+content)
		} else {
			contents = append(contents, content)
		}

	}
	if len(this.Files) > 1 {
		this.json = "{" + strings.Join(contents, ",") + "}"
	} else {
		if !strings.HasPrefix(contents[0], "{") {
			contents[0] = "{" + contents[0] + "}"
		}
		this.json = contents[0]
	}
	return nil
}

// Set configuration
// return error if driver not support set or error happend
func (this *jsonConfig) Set(key string, value interface{}) error {
	this.data[key] = value
	return nil
}

// Get configuration
func (this jsonConfig) Get(key string) interface{} {
	if v, isJSON, exists := this.fetch(key); !exists {
		return nil
	} else if isJSON {
		return v.(gjson.Result).Value()
	} else {
		return v
	}
}

// Exists check if config item exists
func (this jsonConfig) Exists(key string) bool {
	_, _, exists := this.fetch(key)
	return exists
}

// Bool parse dependency as boolean
func (this jsonConfig) Bool(key string, fallback bool) bool {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.False || val.Type == gjson.True {
				return val.Bool()
			}
		} else {
			if val, ok := v.(bool); ok {
				return val
			}
		}
	}
	return fallback
}

// Int parse dependency as int
func (this jsonConfig) Int(key string, fallback int) int {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int(val.Int())
			}
		} else {
			if val, ok := v.(int); ok {
				return val
			}
		}
	}
	return fallback
}

// Int8 parse dependency as int8
func (this jsonConfig) Int8(key string, fallback int8) int8 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int8(val.Int())
			}
		} else {
			if val, ok := v.(int8); ok {
				return val
			}
		}
	}
	return fallback
}

// Int16 parse dependency as int16
func (this jsonConfig) Int16(key string, fallback int16) int16 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int16(val.Int())
			}
		} else {
			if val, ok := v.(int16); ok {
				return val
			}
		}
	}
	return fallback
}

// Int32 parse dependency as int32
func (this jsonConfig) Int32(key string, fallback int32) int32 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int32(val.Int())
			}
		} else {
			if val, ok := v.(int32); ok {
				return val
			}
		}
	}
	return fallback
}

// Int64 parse dependency as int64
func (this jsonConfig) Int64(key string, fallback int64) int64 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return int64(val.Int())
			}
		} else {
			if val, ok := v.(int64); ok {
				return val
			}
		}
	}
	return fallback
}

// UInt parse dependency as uint
func (this jsonConfig) UInt(key string, fallback uint) uint {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint(val.Uint())
			}
		} else {
			if val, ok := v.(uint); ok {
				return val
			}
		}
	}
	return fallback
}

// UInt8 parse dependency as uint8
func (this jsonConfig) UInt8(key string, fallback uint8) uint8 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint8(val.Uint())
			}
		} else {
			if val, ok := v.(uint8); ok {
				return val
			}
		}
	}
	return fallback
}

// UInt16 parse dependency as uint16
func (this jsonConfig) UInt16(key string, fallback uint16) uint16 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint16(val.Uint())
			}
		} else {
			if val, ok := v.(uint16); ok {
				return val
			}
		}
	}
	return fallback
}

// UInt32 parse dependency as uint32
func (this jsonConfig) UInt32(key string, fallback uint32) uint32 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint32(val.Uint())
			}
		} else {
			if val, ok := v.(uint32); ok {
				return val
			}
		}
	}
	return fallback
}

// UInt64 parse dependency as uint64
func (this jsonConfig) UInt64(key string, fallback uint64) uint64 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return uint64(val.Uint())
			}
		} else {
			if val, ok := v.(uint64); ok {
				return val
			}
		}
	}
	return fallback
}

// Float32 parse dependency as float64
func (this jsonConfig) Float32(key string, fallback float32) float32 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return float32(val.Float())
			}
		} else {
			if val, ok := v.(float32); ok {
				return val
			}
		}
	}
	return fallback
}

// Float64 parse dependency as float64
func (this jsonConfig) Float64(key string, fallback float64) float64 {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.Number {
				return val.Float()
			}
		} else {
			if val, ok := v.(float64); ok {
				return val
			}
		}
	}
	return fallback
}

// String parse dependency as string
func (this jsonConfig) String(key string, fallback string) string {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.String {
				return val.String()
			}
		} else {
			if val, ok := v.(string); ok {
				return val
			}
		}
	}
	return fallback
}
