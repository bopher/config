# Config

Configuration manager with default `Env`, `Json` and `Memory` driver.

## Create New Config Driver

Config library contains three different driver by default.

### Env Driver

Env driver use environment file (.env) for managing configuration.

```go
import "github.com/bopher/config"
envConf, err := config.NewEnvConfig("app.env", "db.env", ".env")
```

### JSON Driver

JSON driver use json file for managing configuration.

**Caution:** When you pass multiple file, for accessing config you must pass file name as first part of config path!

```go
import "github.com/bopher/config"
jsonConf, err := config.NewJSONConfig("app.json", "db.json", "global.json")
```

### Memory Driver

Use in-memory array for keeping and managing configuration.

```go
import "github.com/bopher/config"
memConf, err := config.NewMemoryConfig(map[string]interface{}{
    "name": "My First App",
    "key": "My Secret Key",
})
```

## Usage

Config interface contains following methods:

### Load

Load/Reload configurations.

```go
// Signature:
Load() error

// Example
err := envConf.Load()
```

### Set

Set configuration item. this function override preloaded config.

```go
// Signature:
Set(key string, value interface{}) error

// Example
err := memConf.Set("name", "My App")
err = envConf.Set("APP_NAME", "My App")
err = jsonConf.Set("app_name", "My App")
```

**Cation:** For setting/overriding config item in `JSON` driver with multiple files pass filename as first part of config path.

```go
import "github.com/bopher/config"
jsonConf, err := config.NewJSONConfig("file1.json", "file2.json")
err = jsonConf.Set("file1.app.title", "Some")
```

### Get

Get configuration. Get function return config item as `interface{}`. if you need get config with type use helper get functions described later.

```go
// Signature:
Get(key string) interface{}
```

**Caution:** For `JSON` driver with multiple file you must pass filename as first part of config path!

```go
item := jsonConf.Get("file1.app.title")
```

### Exists

Check if config item exists.

```go
// Signature:
Exists(key string) bool
```

### Getters

Getters function allow you to cast config item directly as type. Getters item return error when item not exists or type cast failed!

```go
// BoolE parse item as boolean or return error on fail
BoolE(key string) (bool, error)

// IntE parse item as int or return error on fail
IntE(key string) (int, error)

// Int8E parse item as int8 or return error on fail
Int8E(key string) (int8, error)

// Int16E parse item as int16 or return error on fail
Int16E(key string) (int16, error)

// Int32E parse item as int32 or return error on fail
Int32E(key string) (int32, error)

// Int64E parse item as int64 or return error on fail
Int64E(key string) (int64, error)

// UIntE parse item as uint or return error on fail
UIntE(key string) (uint, error)

// UInt8E parse item as uint8 or return error on fail
UInt8E(key string) (uint8, error)

// UInt16E parse item as uint16 or return error on fail
UInt16E(key string) (uint16, error)

// UInt32E parse item as uint32 or return error on fail
UInt32E(key string) (uint32, error)

// UInt64E parse item as uint64 or return error on fail
UInt64E(key string) (uint64, error)

// Float64E parse item as float64 or return error on fail
Float64E(key string) (float64, error)

// StringE parse item as string or return error on fail
StringE(key string) (string, error)
```

### Error Safe Getters

You can use safe getters to cast config item and pass fallback value in case of item casting failed!

```go
// Bool parse item as boolean or return fallback
Bool(key string, fallback bool) bool

// Int parse item as int or return fallback
Int(key string, fallback int) int

// Int8 parse item as int8 or return fallback
Int8(key string, fallback int8) int8

// Int16 parse item as int16  or return fallback
Int16(key string, fallback int16) int16

// Int32 parse item as int32 or return fallback
Int32(key string, fallback int32) int32

// Int64 parse item as int64 or return fallback
Int64(key string, fallback int64) int64

// UInt parse item as uint or return fallback
UInt(key string, fallback uint) uint

// UInt8 parse item as uint8 or return fallback
UInt8(key string, fallback uint8) uint8

// UInt16 parse item as uint16 or return fallback
UInt16(key string, fallback uint16) uint16

// UInt32 parse item as uint32 or return fallback
UInt32(key string, fallback uint32) uint32

// UInt64 parse item as uint64 or return fallback
UInt64(key string, fallback uint64) uint64

// Float64 parse item as float64 or return fallback
Float64(key string, fallback float64) float64

// String parse item as string or return fallback
String(key string, fallback string) string
```
