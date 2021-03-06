// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package ansiterm

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

// colorEnabledWriter returns a writer that can handle the ansi color codes
// and true if the writer passed in is a terminal capable of color. If the
// TERM environment variable is set to "dumb", the terminal is not considered
// color capable.
func colorEnabledWriter(w io.Writer) (io.Writer, bool) {
	f, ok := w.(*os.File)
	if !ok {
		return w, false
	}

	// NO_COLOR is a relatively new standard for preventing color enabled
	// writers rather than using the TERM env.
	//
	// "...should check for the presence of a NO_COLOR environment variable
	//  that, when present (regardless of its value), prevents the addition of
	//  ANSI color."
	// See: https://no-color.org/
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return w, false
	}

	// Check the TERM environment variable specifically
	// to check for "dumb" terminals.
	if os.Getenv("TERM") == "dumb" {
		return w, false
	}
	if !isatty.IsTerminal(f.Fd()) {
		return w, false
	}
	return colorable.NewColorable(f), true
}
