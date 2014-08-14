package completecommand

import (
	"flag"
	"fmt"
	"os"
)

const completeFlagName = "__complete__"

var completeFlag = flag.Bool(completeFlagName, false, "return autocomplete details")

// Complete checks whether the __complete__ flag is set, and if so returns a list of all the
// defined flags which corresponds to the format expected by the _arguments function in zsh, and
// then immediately exits the executable. If the __complete__ flag is not set, then it just returns
// immediately.
func Complete() {
	flag.Parse()
	if !*completeFlag {
		return
	}

	// TODO: Generate more friendly candiates for Bash completion. See
	// http://fahdshariff.blogspot.co.uk/2011/04/writing-your-own-bash-completion.html

	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == completeFlagName {
			return
		}

		// TODO: Argument type:
		// fmt.Printf("-%s=-[%s]:::_files\n", f.Name, f.Usage)

		fmt.Printf("-%s=-[%s]\n", f.Name, f.Usage)
	})

	// TODO: Multiple choice:
	// fmt.Printf(":action:(aaa bbb ccc)\n")

	// TODO: File names:
	// fmt.Printf("*::file:_files\n")

	os.Exit(0)
}
