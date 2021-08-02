package config

// NewEnvConfig create a new env file configuration manager instance
func NewEnvConfig(filenames ...string) (Config, bool) {
	ec := &envConfig{Files: filenames}
	if ec.Load() {
		return ec, true
	}
	return nil, false
}

// NewJSONConfig create a new json file configuration manager instance
func NewJSONConfig(filenames ...string) (Config, bool) {
	jc := &jsonConfig{Files: filenames}
	if jc.Load() {
		return jc, true
	}
	return nil, false
}

// NewMemoryConfig create a new in-memory configuration manager instance
func NewMemoryConfig(config map[string]interface{}) (Config, bool) {
	mc := new(memoryConfig)
	mc.Load()
	if config != nil {
		mc.data = config
	}
	return mc, true
}
