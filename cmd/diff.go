package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/pmezard/go-difflib/difflib"
)

// diff print the diff between original dataset accesses and the ones from the configuration file
func diff(meta AccessList, metaToUpdate AccessList) bool {

	sort.Sort(meta)
	sort.Sort(metaToUpdate)

	if !reflect.DeepEqual(meta, metaToUpdate) {
		text, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
			A: meta.Entries(),
			B: metaToUpdate.Entries(),
		})

		for _, line := range strings.Split(text, "\n") {
			switch {
			case strings.HasPrefix(line, "+"):
				fmt.Println(ansi.Color(fmt.Sprintf("  %s", line), "green"))
			case strings.HasPrefix(line, "-"):
				fmt.Println(ansi.Color(fmt.Sprintf("  %s", line), "red"))
			case strings.HasPrefix(line, "@"):
				continue
			default:
				fmt.Println(line)
			}
		}
		return true
	}
	fmt.Println("  Already up-to-date")
	return false
}
