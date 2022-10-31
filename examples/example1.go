//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"

	"github.com/goinsane/flagbind"
)

func main() {
	fs := flag.NewFlagSet("test", flag.PanicOnError)
	var st struct {
		BoolFlag    bool
		BoolFlag2   *bool
		IntFlag     int `default:"25"`
		IntFlag2    *int
		StrFlag     string  `name:"str" default:"abc" usage:"str usage"`
		CustomFlag  float64 `name:"cust" usage:"custom flag usage"`
		IgnoredFlag int64   `name:"-"`
	}
	flagbind.Bind(fs, &st)
	args := []string{"-bool-flag", "-bool-flag2", "true", "-int-flag", "10", "-str", "def", "-cust", "10.6"}
	_ = fs.Parse(args)
	fmt.Printf("%+v\n", st)
}
