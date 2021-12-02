package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/bopher/utils"
	"github.com/tidwall/gjson"
)

type jsonConfig struct {
	Files []string
	json  string
	data  map[string]interface{}
}

func (jsonConfig) err(format string, args ...interface{}) error {
	return utils.TaggedError([]string{"JSONConfig"}, format, args...)
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
			return this.err(fmt.Sprintf("%s content is invalid!", f))
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

func (this *jsonConfig) Set(key string, value interface{}) error {
	this.data[key] = value
	return nil
}

func (this jsonConfig) Get(key string) interface{} {
	if v, isJSON, exists := this.fetch(key); !exists {
		return nil
	} else if isJSON {
		return v.(gjson.Result).Value()
	} else {
		return v
	}
}

func (this jsonConfig) Exists(key string) bool {
	_, _, exists := this.fetch(key)
	return exists
}

func (this jsonConfig) BoolE(key string) (bool, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastBoolE(v); err != nil {
			return false, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return false, this.err("failed to get %s as bool.", key)
}

func (this jsonConfig) Bool(key string, fallback bool) bool {
	if v, err := this.BoolE(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) IntE(key string) (int, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastIntE(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as int.", key)
}

func (this jsonConfig) Int(key string, fallback int) int {
	if v, err := this.IntE(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) Int8E(key string) (int8, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastInt8E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as int8.", key)
}

func (this jsonConfig) Int8(key string, fallback int8) int8 {
	if v, err := this.Int8E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) Int16E(key string) (int16, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastInt16E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as int16.", key)
}

func (this jsonConfig) Int16(key string, fallback int16) int16 {
	if v, err := this.Int16E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) Int32E(key string) (int32, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastInt32E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as int32.", key)
}

func (this jsonConfig) Int32(key string, fallback int32) int32 {
	if v, err := this.Int32E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) Int64E(key string) (int64, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastInt64E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as int64.", key)
}

func (this jsonConfig) Int64(key string, fallback int64) int64 {
	if v, err := this.Int64E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) UIntE(key string) (uint, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastUIntE(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as uint.", key)
}

func (this jsonConfig) UInt(key string, fallback uint) uint {
	if v, err := this.UIntE(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) UInt8E(key string) (uint8, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastUInt8E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as uint8.", key)
}

func (this jsonConfig) UInt8(key string, fallback uint8) uint8 {
	if v, err := this.UInt8E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) UInt16E(key string) (uint16, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastUInt16E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as uint16.", key)
}

func (this jsonConfig) UInt16(key string, fallback uint16) uint16 {
	if v, err := this.UInt16E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) UInt32E(key string) (uint32, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastUInt32E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as uint32.", key)
}

func (this jsonConfig) UInt32(key string, fallback uint32) uint32 {
	if v, err := this.UInt32E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) UInt64E(key string) (uint64, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastUInt64E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as uint64.", key)
}

func (this jsonConfig) UInt64(key string, fallback uint64) uint64 {
	if v, err := this.UInt64E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) Float64E(key string) (float64, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			v = v.(gjson.Result).Raw
		}

		if v, err := utils.CastFloat64E(v); err != nil {
			return 0, this.err(err.Error())
		} else {
			return v, nil
		}
	}
	return 0, this.err("failed to get %s as float64.", key)
}

func (this jsonConfig) Float64(key string, fallback float64) float64 {
	if v, err := this.Float64E(key); err == nil {
		return v
	}
	return fallback
}

func (this jsonConfig) StringE(key string) (string, error) {
	if v, isJSON, exists := this.fetch(key); exists {
		if isJSON {
			val := v.(gjson.Result)
			if val.Type == gjson.String {
				return val.String(), nil
			}
		} else {
			if val, ok := v.(string); ok {
				return val, nil
			}
		}
	}
	return "", this.err("failed cast %s as string!", key)
}

func (this jsonConfig) String(key string, fallback string) string {
	if v, err := this.StringE(key); err == nil {
		return v
	}
	return fallback
}
