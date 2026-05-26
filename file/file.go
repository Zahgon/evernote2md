package file

import (
	"io"
	"regexp"
)

const (
	// Mon Jan 2 15:04:05 -0700 MST 2006 represented as yyyyMMddhhmm
	touchTimeFormat = "200601021504"

	// OS allow 255 character for filenames = 252 + 3 (.md)
	maxNameChars = 252
)

var (
	baseNameSeparators = regexp.MustCompile(`[./]`)

	dashes = regexp.MustCompile(`[\-_]{2,}`)
)

// Save a new file in a given dir with the following content.
// Creates a directory if necessary.
func Save(dir, name string, content io.Reader) error { _ = "STUB: not implemented"; return nil }

// BaseName normalizes a given string to use it as a safe filename
func BaseName(s string) string {
	_ = "STUB: not implemented"
	// Replace separator characters with a dash
	return ""
}

// Remove any trailing space to avoid ending on -

// Replace inappropriate characters with an underscore

// Remove any multiple dashes caused by replacements above

// Check file name length in bytes

// Trim filename to the max allowed number of bytes
