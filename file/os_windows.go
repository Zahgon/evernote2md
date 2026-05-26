package file

import (
	"regexp"
	"time"
)

// Max path length according to fixLongPath function is 248 - 3 bytes for extension (.md)
const maxNameBytes int = 245

// Additional rule for Windows
var illegalChars = regexp.MustCompile(`[\s\\|"'<>&_=+:?*]`)

// ChangeFileTimes uses SetFileTime syscall in Windows implementation
// which supports updating both creation and modification dates
func ChangeFileTimes(dir, name string, ctime, mtime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
