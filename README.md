# flagbind

[![Go Reference](https://pkg.go.dev/badge/github.com/goinsane/flagbind.svg)](https://pkg.go.dev/github.com/goinsane/flagbind)

Package flagbind provides utilities to bind flags of GoLang's flag package to struct.

flagbind supports these types:

- bool, *bool
- int, *int
- uint, *uint
- int64, *int64
- uint64, *uint64
- int32, *int32
- uint32, *uint32
- string, *string
- float64, *float64
- float32, *float32
- time.Duration, *time.Duration
- flag.Value
- func(value string) error
- func(name string, value string) error
