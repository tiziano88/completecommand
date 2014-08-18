package completecommand

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const completeFlagName = "__complete__"

var completeFlag = flag.String(completeFlagName, "", "return autocomplete details")

// Complete checks whether the __complete__ flag is set, and if so returns a list of all the
// defined flags which corresponds to the format expected by the _arguments function in zsh, and
// then immediately exits the executable. If the __complete__ flag is not set, then it just returns
// immediately.
func Complete() {
	flag.Parse()

	switch *completeFlag {
	case "":
		return
	case "zsh":
		completeZSH()
	case "bash":
		completeBash()
	}

	os.Exit(0)
}

func completeZSH() {
	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == completeFlagName {
			return
		}

		// TODO: Argument type:
		// fmt.Printf("-%s=-[%s]:::_files\n", f.Name, f.Usage)

		fmt.Printf("-%s=-[%s (%q)]\n", f.Name, f.Usage, f.Value.String())
	})

	// Allow completing unlimited file names in the general case.
	fmt.Printf("*:file:_files\n")
}

func completeBash() {
	// https://www.gnu.org/software/bash/manual/html_node/Programmable-Completion-Builtins.html

	names := make([]string, 0)
	flag.VisitAll(func(f *flag.Flag) {
		if f.Name == completeFlagName {
			return
		}
		names = append(names, "-"+f.Name)
	})
	fmt.Printf("-W '%s' ", strings.Join(names, " "))

	// Allow completing unlimited file names in the general case.
	fmt.Printf("-f ")
}
