//Copyright: 2020 Kacper Kazimierak

package sluggo

import (
	"regexp"
	"strings"
)

type SlugOptions struct {
	TrimSpaces bool
	// If true spaces will be trimmed from both ends of the input text, default true.
	// If PreserveLength set to true - spaces will not be trimmed

	Case int
	// Case conversion: sluggo.DefaultCase = 0, sluggo.Lowercase = 1, sluggo.Uppercase = 2, default Lowercase

	SpaceSymbol string
	// Symbol for replacing space. URL safe symbols are; a-z A-Z 0-9 "." "-" "_" "~", default _

	CharSymbol string
	// Symbol for replacing unwanted chars.  URL safe symbols are; a-z A-Z 0-9 "." "-" "_" "~", default -

	MaxLength int
	// Trim text to this length, if set to 0 text will not be trimmed, default 32 chars

	PreserveLength bool
	// By default neighbouring chars to replace will be replaced with 1 CharSymbol or SpaceSymbol,
	//for example: "foo  b@@@r" => "foo_b-r". If true input text and output text will be the same length,
	//for example (with default SpaceSymbol and CharSymbol): "foo  bar" => "foo__bar" instead of "foo_bar",
	//"foo b@@r" => "foo_b--r" instead of "foo_b-r", default false
}

const (
	DefaultCase = 0
	Lowercase   = 1
	Uppercase   = 2
	DefaultMaxLength = 32
)

var (
	allowedCharsRegexp = regexp.MustCompile("[^a-z0-9-_~.]+")
	whitespaceRegexp   = regexp.MustCompile("\\s+")

	preserveAllowedCharsRegexp = regexp.MustCompile("[^a-z0-9-_~.]")
	preserveWhitespaceRegexp   = regexp.MustCompile("\\s")
)

func NewDefaultOptions() *SlugOptions {
	return &SlugOptions{
		TrimSpaces:     true,
		MaxLength:      DefaultMaxLength,
		Case:           Lowercase,
		SpaceSymbol:    "_",
		CharSymbol:     "-",
		PreserveLength: false,
	}
}

// Returns slug from supplied text using default options:
// - TrimSpaces:     true,
// - MaxLength:      0,
// - Case:           Lowercase,
// - SpaceSymbol:    "_",
// - CharSymbol:     "-",
// - PreserveLength: false,
func GetSlug(text string) string {
	// Run GetSlugWithOpts with default options set
	return GetSlugWithOpts(text, NewDefaultOptions())
}

// Return slug created using parameters provided in *SlugOptions
func GetSlugWithOpts(text string, opts *SlugOptions) string {

	var spaceRegexp, charsRegexp *regexp.Regexp

	// Trim trailing and leading spaces
	if opts.TrimSpaces && !opts.PreserveLength {
		text = strings.TrimSpace(text)
	}

	switch opts.Case {
	case Lowercase:
		text = strings.ToLower(text)
	case Uppercase:
		text = strings.ToUpper(text)
	}

	if opts.PreserveLength {
		charsRegexp = preserveAllowedCharsRegexp
		spaceRegexp = preserveWhitespaceRegexp
	} else {
		charsRegexp = allowedCharsRegexp
		spaceRegexp = whitespaceRegexp
	}

	//	Remove spaces
	text = spaceRegexp.ReplaceAllString(text, opts.SpaceSymbol)

	//	Remove unwanted characters, leaving only url safe symbols
	text = charsRegexp.ReplaceAllString(text, opts.CharSymbol)

	//Truncate slug to required MaxLength, if MaxLength is 0 text will not be truncated
	if len(text) > opts.MaxLength && opts.MaxLength > 0 {
		text = text[:opts.MaxLength]
	}

	return text
}
