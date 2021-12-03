package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/bopher/caster"
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

func (this jsonConfig) Cast(key string) caster.Caster {
	return caster.NewCaster(this.Get(key))
}
