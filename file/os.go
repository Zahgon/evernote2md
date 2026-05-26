//go:build !windows

package file

import (
	"regexp"
	"time"
)

// Max path length in bytes, determined empirically.
// P.S. Don't trust Apple documentation
const maxNameBytes int = 704

// Semicolon is not allowed in MacOS and spaces is just my personal preference
var illegalChars = regexp.MustCompile(`[\s:]`)

// ChangeFileTimes matches the file times with the Evernote metadata
//
// Uses touch if available to change both creation  and modification date
// Otherwise it falls back to os.Chtimes to change only the modification date
func ChangeFileTimes(dir, name string, ctime, mtime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// On macOS, first touch set both creation date and modification date
// On Linux, this touch will be ignored by second touch. There is no easy way to setting creation date

// On macOS, second touch updates the modification date and the creation date is preserved
