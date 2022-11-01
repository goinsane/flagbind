// Package flagbind provides utilities to bind flags of GoLang's flag package to struct.

package flagbind

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Bind binds variables of the given flag.FlagSet to the target that is struct pointer.
// Fields of target can have tags such as name, default, usage.
// If name tag of struct field is not given, flag name will be generated by struct field name (ex: BoolFlag is -bool-flag).
// If name tag is "-" or struct field isn't exported or anonymous, struct field will be ignored.
// It will panic when target isn't struct pointer or nil, or any struct field has unknown type.
func Bind(fs *flag.FlagSet, target interface{}) {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Pointer || val.Elem().Kind() != reflect.Struct {
		panic("target must be struct pointer")
	}
	if val.IsNil() {
		panic("target nil pointer")
	}
	val = val.Elem()
	typ := val.Type()
	for i, j := 0, typ.NumField(); i < j; i++ {
		sVal := val.Field(i)
		sField := typ.Field(i)
		if !sField.IsExported() || sField.Anonymous || sField.Tag.Get("name") == "-" {
			continue
		}
		parser := &_Parser{
			Target:    sVal,
			Name:      toArgName(sField.Name),
			Default:   "",
			DefaultOK: false,
			Usage:     sField.Tag.Get("usage"),
		}
		if v, ok := sField.Tag.Lookup("name"); ok {
			parser.Name = v
		}
		if v, ok := sField.Tag.Lookup("default"); ok {
			parser.Default = v
			parser.DefaultOK = true
		}
		if sVal.Kind() == reflect.Bool {
			fs.BoolVar(sVal.Addr().Interface().(*bool), parser.Name, false, parser.Usage)
			continue
		}
		parser.Reset()
		parser.SetDefault()
		fs.Func(parser.Name, parser.Usage, parser.Set)
	}
}

type _Parser struct {
	Target    reflect.Value
	Name      string
	Default   string
	DefaultOK bool
	Usage     string
}

func (p *_Parser) Set(value string) (err error) {
	ifc, _, kind := p.Target.Interface(), p.Target.Type(), p.Target.Kind()
	switch ifc.(type) {
	case bool, *bool:
		var x bool
		x, err = strconv.ParseBool(value)
		if err != nil {
			return errParse
		}
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(x))
		} else {
			p.Target.Set(reflect.ValueOf(&x))
		}
	case int, *int:
		var x int64
		x, err = strconv.ParseInt(value, 10, 0)
		if err != nil {
			return numError(err)
		}
		y := int(x)
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(y))
		} else {
			p.Target.Set(reflect.ValueOf(&y))
		}
	case uint, *uint:
		var x uint64
		x, err = strconv.ParseUint(value, 10, 0)
		if err != nil {
			return numError(err)
		}
		y := uint(x)
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(y))
		} else {
			p.Target.Set(reflect.ValueOf(&y))
		}
	case int64, *int64:
		var x int64
		x, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return numError(err)
		}
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(x))
		} else {
			p.Target.Set(reflect.ValueOf(&x))
		}
	case uint64, *uint64:
		var x uint64
		x, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			return numError(err)
		}
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(x))
		} else {
			p.Target.Set(reflect.ValueOf(&x))
		}
	case int32, *int32:
		var x int64
		x, err = strconv.ParseInt(value, 10, 32)
		if err != nil {
			return numError(err)
		}
		y := int32(x)
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(y))
		} else {
			p.Target.Set(reflect.ValueOf(&y))
		}
	case uint32, *uint32:
		var x uint64
		x, err = strconv.ParseUint(value, 10, 32)
		if err != nil {
			return numError(err)
		}
		y := uint32(x)
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(y))
		} else {
			p.Target.Set(reflect.ValueOf(&y))
		}
	case string, *string:
		x := value
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(x))
		} else {
			p.Target.Set(reflect.ValueOf(&x))
		}
	case float64, *float64:
		var x float64
		x, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return numError(err)
		}
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(x))
		} else {
			p.Target.Set(reflect.ValueOf(&x))
		}
	case float32, *float32:
		var x float64
		x, err = strconv.ParseFloat(value, 32)
		if err != nil {
			return numError(err)
		}
		y := float32(x)
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(y))
		} else {
			p.Target.Set(reflect.ValueOf(&y))
		}
	case time.Duration, *time.Duration:
		var x time.Duration
		x, err = time.ParseDuration(value)
		if err != nil {
			return errParse
		}
		if kind != reflect.Pointer {
			p.Target.Set(reflect.ValueOf(x))
		} else {
			p.Target.Set(reflect.ValueOf(&x))
		}
	case flag.Value:
		err = ifc.(flag.Value).Set(value)
		if err != nil {
			return err
		}
	case func(string) error:
		err = ifc.(func(string) error)(value)
		if err != nil {
			return err
		}
	case func(string, string) error:
		err = ifc.(func(string, string) error)(p.Name, value)
		if err != nil {
			return err
		}
	default:
		panic(fmt.Errorf("unknown type for flag -%s", p.Name))
	}
	return nil
}

func (p *_Parser) Reset() {
	ifc, typ, kind := p.Target.Interface(), p.Target.Type(), p.Target.Kind()
	switch ifc.(type) {
	case bool, *bool:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case int, *int:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case uint, *uint:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case int64, *int64:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case uint64, *uint64:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case int32, *int32:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case uint32, *uint32:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case string, *string:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case float64, *float64:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case float32, *float32:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case time.Duration, *time.Duration:
		if kind != reflect.Pointer {
		} else {
			p.Target.Set(reflect.Zero(typ))
		}
	case flag.Value:
	case func(string) error:
	case func(string, string) error:
	default:
		panic(fmt.Errorf("unknown type for flag -%s", p.Name))
	}
}

func (p *_Parser) SetDefault() {
	if p.DefaultOK {
		if e := p.Set(p.Default); e != nil {
			panic(fmt.Errorf("unable to set default value for flag -%s: %w", p.Name, e))
		}
		return
	}
}
