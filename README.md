# Config

Configuration manager with default `Env`, `Json` and `Memory` driver.

## Create New Config Driver

Config library contains three different driver by default.

### Env Driver

Env driver use environment file (.env) for managing configuration.

```go
import "github.com/bopher/config"
envConf, ok := config.NewEnvConfig("app.env", "db.env", ".env")
```

### JSON Driver

JSON driver use json file for managing configuration.

**Caution:** When you pass multiple file, for accessing config you must pass file name as first part of config path!

```go
import "github.com/bopher/config"
jsonConf, ok := config.NewJSONConfig("app.json", "db.json", "global.json")
```

### Memory Driver

Use in-memory array for keeping and managing configuration.

```go
import "github.com/bopher/config"
memConf, ok := config.NewMemoryConfig(map[string]interface{}{
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
Load() bool

// Example
ok := envConf.Load()
```

### Set

Set configuration item. this function override in-file config in `Env` and `JSON` driver.

```go
// Signature:
Set(key string, value interface{}) bool

// Example
memConf.Set("name", "My App")
envConf.Set("APP_NAME", "My App")
jsonConf.Set("app.name", "My App")
```

**Cation:** For setting/overriding config item in `JSON` driver with multiple files pass filename as first part of config path.

```go
import "github.com/bopher/config"
jsonConf, ok := config.NewJSONConfig("file1.json", "file2.json")
jsonConf.Set("file1.app.title", "Some")
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

### Get By Type Methods

```go
// Bool parse dependency as boolean
// return fallback, false if value not exists
Bool(key string, fallback bool) bool
// Int parse dependency as int
// return fallback, false if value not exists
Int(key string, fallback int) int
// Int8 parse dependency as int8
// return fallback, false if value not exists
Int8(key string, fallback int8) int8
// Int16 parse dependency as int16
// return fallback, false if value not exists
Int16(key string, fallback int16) int16
// Int32 parse dependency as int32
// return fallback, false if value not exists
Int32(key string, fallback int32) int32
// Int64 parse dependency as int64
// return fallback, false if value not exists
Int64(key string, fallback int64) int64
// UInt parse dependency as uint
// return fallback, false if value not exists
UInt(key string, fallback uint) uint
// UInt8 parse dependency as uint8
// return fallback, false if value not exists
UInt8(key string, fallback uint8) uint8
// UInt16 parse dependency as uint16
// return fallback, false if value not exists
UInt16(key string, fallback uint16) uint16
// UInt32 parse dependency as uint32
// return fallback, false if value not exists
UInt32(key string, fallback uint32) uint32
// UInt64 parse dependency as uint64
// return fallback, false if value not exists
UInt64(key string, fallback uint64) uint64
// Float32 parse dependency as float64
// return fallback, false if value not exists
Float32(key string, fallback float32) float32
// Float64 parse dependency as float64
// return fallback, false if value not exists
Float64(key string, fallback float64) float64
// String parse dependency as string
// return fallback, false if value not exists
String(key string, fallback string) string
```
