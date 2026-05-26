// Evernote2md is a cli tool to convert Evernote notes exported in *.enex format
// to a directory with markdown files.
//
// Usage:
//
//	evernote2md <file> [-o <outputDir>]
//
// If outputDir is not specified, current workdir is used.
package main

import (
	"path/filepath"

	"github.com/briandowns/spinner"
	"github.com/integrii/flaggy"

	"github.com/wormi4ok/evernote2md/internal"
)

var version = "dev"

func init() {
	flaggy.SetName("evernote2md")
	flaggy.SetDescription(" Convert Evernote notes exported in *.enex format to markdown files")
	flaggy.SetVersion(version)

	flaggy.DefaultParser.ShowCompletion = false
	flaggy.DefaultParser.ShowHelpOnUnexpected = false
	flaggy.DefaultParser.AdditionalHelpPrepend = "http://github.com/wormi4ok/evernote2md"
}

func main() {
	var input, outputOverride string
	var outputDir = filepath.FromSlash("./notes")
	var tagTemplate = internal.DefaultTagTemplate
	var folders, noHighlights, escapeSpecialChars, resetTimestamps, addFrontMatter, debug bool

	flaggy.AddPositionalValue(&input, "input", 1, true, "Evernote export file, directory or a glob pattern")
	flaggy.AddPositionalValue(&outputDir, "output", 2, false, "Output directory")

	flaggy.String(&tagTemplate, "t", "tagTemplate", "Define how Evernote tags are formatted")
	flaggy.String(&outputOverride, "o", "outputDir", "Override the directory where markdown files will be created")

	flaggy.Bool(&folders, "", "folders", "Put every note in a separate folder")
	flaggy.Bool(&noHighlights, "", "noHighlights", "Disable converting Evernote highlights to inline HTML tags")
	flaggy.Bool(&escapeSpecialChars, "", "escape-special-chars", "Escape special characters to ensure correct rendering of the converted files")
	flaggy.Bool(&resetTimestamps, "", "resetTimestamps", "Create files ignoring timestamps in the note attributes")
	flaggy.Bool(&addFrontMatter, "", "addFrontMatter", "Prepend FrontMatter to markdown files")
	flaggy.Bool(&debug, "v", "debug", "Show debug output")

	flaggy.Parse()

	if len(outputOverride) > 0 {
		outputDir = outputOverride
	}

	files, err := matchInput(input)
	failWhen(err)
	output := newNoteFilesDir(outputDir, folders, !resetTimestamps)
	converter, err := internal.NewConverter(tagTemplate, addFrontMatter, !noHighlights, escapeSpecialChars)
	failWhen(err)

	setLogLevel(debug)
	run(files, output, newSpinner(debug), converter)
}

func newSpinner(disabled bool) *spinner.Spinner { _ = "STUB: not implemented"; return nil }

func run(files []string, output *noteFilesDir, sp *spinner.Spinner, c *internal.Converter) {
	_ = "STUB: not implemented"
	return
}

func progressError(err error, name string, text string) bool {
	_ = "STUB: not implemented"
	return false
}

// Erase current spinner

// matchInput finds all files matching input pattern
// If input is a path to a directory, it will search for *.enex files inside the directory
func matchInput(input string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// If input is a directory, find all *.enex files and return

// User glob patterns may include directories that we filter out

func setLogLevel(debug bool) { _ = "STUB: not implemented"; return }

func failWhen(err error) { _ = "STUB: not implemented"; return }
