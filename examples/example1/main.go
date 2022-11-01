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
		FuncFlag    func(string) error
		DefFlag     int `default:"35"`
		Def2Flag    int
	}
	st.FuncFlag = func(value string) error {
		fmt.Println(value)
		return nil
	}
	st.DefFlag = 45
	st.Def2Flag = 55
	flagbind.Bind(fs, &st)
	args := []string{"-bool-flag", "-bool-flag2", "true", "-int-flag", "10", "-str", "def", "-cust", "10.6", "-func-flag", "aaa"}
	_ = fs.Parse(args)
	fmt.Printf("%+v\n", st)
}
