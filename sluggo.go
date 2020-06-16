package sluggo

import (
	"regexp"
	"strings"
)

type SlugOptions struct {
	TrimSpaces       bool
	Case             int
	WhiteSpaceSymbol string
	CharSymbol       string
	MaxLength        int
}

const (
	DefaultCase = 0
	Lowercase = 1
	Uppercase = 2
)

var (
	whitelistCharsRegexp = regexp.MustCompile("[^a-z0-9-_~.]+")
	whiteSpaceRegexp = regexp.MustCompile("[ ]+")
)

func NewDefaultOptions() *SlugOptions {
	return &SlugOptions{
		TrimSpaces:       true,
		MaxLength:        0,
		Case:             Lowercase,
		WhiteSpaceSymbol: "_",
		CharSymbol:       "-",
	}
}

func GetSlug(text string) string {
	return GetSlugWithOpts(text, NewDefaultOptions())
}

func GetSlugWithOpts(text string, opts *SlugOptions) string {

	// Trim trailing and leading spaces
	if opts.TrimSpaces {
		text = strings.TrimSpace(text)
	}

	switch opts.Case {
	case Lowercase:
		text = strings.ToLower(text)
	case Uppercase:
		text = strings.ToUpper(text)
	}

	//	Remove spaces
	text = whiteSpaceRegexp.ReplaceAllString(text, opts.WhiteSpaceSymbol)

	//	Remove unwanted characters, leaving only url safe symbols
	text = whitelistCharsRegexp.ReplaceAllString(text, opts.CharSymbol)

	//Truncate slug to required MaxLength, if MaxLength is 0 text will not be truncated
	if len(text) > opts.MaxLength && opts.MaxLength > 0 {
		text = text[:opts.MaxLength-1]
	}

	return text
}
