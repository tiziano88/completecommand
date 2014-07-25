package completecommand

import (
	"flag"
	"fmt"
	"os"
)

const completeFlagName = "__complete__"

var completeFlag = flag.Bool(completeFlagName, false, "return autocomplete details")

// Complete checks whether the __complete__ flag is set, and if so returns a space-separated list
// of all the defined flags, and then immediately exits the executable. If the __complete__ flag is
// not set, then it just returns immediately.
func Complete() {
	flag.Parse()
	if !*completeFlag {
		return
	}
	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == completeFlagName {
			return
		}
		fmt.Printf("%s ", f.Name)
	})
	os.Exit(0) // TODO: Maybe use magic number.
}
