package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// Observe how the b's and the d's, despite appearing in the
	// second cell of each line, belong to different columns.
	w := tabwriter.NewWriter(os.Stdout, 4, 4,  1,'.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()

	fmt.Println("-------------")

	w = new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()
}