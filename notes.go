package main

import (
	"github.com/wormi4ok/evernote2md/encoding/markdown"
)

// noteFilesDir saves markdown notes in a directory on the filesystem
type noteFilesDir struct {
	path string

	// flags modifying the logic for saving notes
	flagFolders    bool
	flagTimestamps bool

	// A map to keep track of what notes are already created
	names map[string]int
}

func newNoteFilesDir(output string, folders, timestamps bool) *noteFilesDir {
	_ = "STUB: not implemented"
	return nil
}

// SaveNote along with media resources
func (d *noteFilesDir) SaveNote(title string, md *markdown.Note) error {
	_ = "STUB: not implemented"
	return nil
}

// Continue processing on error

func (d *noteFilesDir) Path() string {
	_ = "STUB: not implemented"

	// uniqueName returns a unique note name
	return ""
}

func (d *noteFilesDir) uniqueName(title string) string { _ = "STUB: not implemented"; return "" }
