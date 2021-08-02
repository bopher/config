package config

type memoryConfig struct {
	data map[string]interface{}
}

// Load configurations
func (c *memoryConfig) Load() bool {
	// Load do nothing for memory config
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	return true
}

// Set configuration
// return false if driver not support set or error happend
func (c *memoryConfig) Set(key string, value interface{}) bool {
	c.data[key] = value
	return true
}

// Get configuration
func (c *memoryConfig) Get(key string) interface{} {
	if v, ok := c.data[key]; ok {
		return v
	}
	return nil
}

// Exists check if config item exists
func (c *memoryConfig) Exists(key string) bool {
	if _, ok := c.data[key]; ok {
		return true
	}
	return false
}

// Bool parse dependency as boolean
func (c *memoryConfig) Bool(key string, fallback bool) bool {
	if v, ok := c.data[key].(bool); ok {
		return v
	}
	return fallback
}

// Int parse dependency as int
func (c *memoryConfig) Int(key string, fallback int) int {
	if v, ok := c.data[key].(int); ok {
		return v
	}
	return fallback
}

// Int8 parse dependency as int8
func (c *memoryConfig) Int8(key string, fallback int8) int8 {
	if v, ok := c.data[key].(int8); ok {
		return v
	}
	return fallback
}

// Int16 parse dependency as int16
func (c *memoryConfig) Int16(key string, fallback int16) int16 {
	if v, ok := c.data[key].(int16); ok {
		return v
	}
	return fallback
}

// Int32 parse dependency as int32
func (c *memoryConfig) Int32(key string, fallback int32) int32 {
	if v, ok := c.data[key].(int32); ok {
		return v
	}
	return fallback
}

// Int64 parse dependency as int64
func (c *memoryConfig) Int64(key string, fallback int64) int64 {
	if v, ok := c.data[key].(int64); ok {
		return v
	}
	return fallback
}

// UInt parse dependency as uint
func (c *memoryConfig) UInt(key string, fallback uint) uint {
	if v, ok := c.data[key].(uint); ok {
		return v
	}
	return fallback
}

// UInt8 parse dependency as uint8
func (c *memoryConfig) UInt8(key string, fallback uint8) uint8 {
	if v, ok := c.data[key].(uint8); ok {
		return v
	}
	return fallback
}

// UInt16 parse dependency as uint16
func (c *memoryConfig) UInt16(key string, fallback uint16) uint16 {
	if v, ok := c.data[key].(uint16); ok {
		return v
	}
	return fallback
}

// UInt32 parse dependency as uint32
func (c *memoryConfig) UInt32(key string, fallback uint32) uint32 {
	if v, ok := c.data[key].(uint32); ok {
		return v
	}
	return fallback
}

// UInt64 parse dependency as uint64
func (c *memoryConfig) UInt64(key string, fallback uint64) uint64 {
	if v, ok := c.data[key].(uint64); ok {
		return v
	}
	return fallback
}

// Float32 parse dependency as float64
func (c *memoryConfig) Float32(key string, fallback float32) float32 {
	if v, ok := c.data[key].(float32); ok {
		return v
	}
	return fallback
}

// Float64 parse dependency as float64
func (c *memoryConfig) Float64(key string, fallback float64) float64 {
	if v, ok := c.data[key].(float64); ok {
		return v
	}
	return fallback
}

// String parse dependency as string
func (c *memoryConfig) String(key string, fallback string) string {
	if v, ok := c.data[key].(string); ok {
		return v
	}
	return fallback
}
