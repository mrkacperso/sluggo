# sluggo
Lightweight go library to create url-friendly text slugs used im my other projects.
Creates url-friendly slugs from strings.

## Usage 
### Importing
```go
import 	"github.com/mrkacperso/sluggo"
```
### With default options
Example:
```go
slug := sluggo.GetSlug("Foo, bar And other")
```
will return
`foo-_bar_and_other`

Default options are:
 - TrimSpaces bool, default true -
 If true spaces will be trimmed from both ends of the input text. If PreserveLength set to true - spaces will not be trimmed.

 - Case int, default sluggo.Lowercase - Case conversion, options: 
 sluggo.DefaultCase = 0, sluggo.Lowercase = 1, sluggo.Uppercase = 2, default Lowercase
 if none of above options is provided by default sluggo.DefaultCase = 0 will be used.

 - SpaceSymbol string, default "_" - Symbol for replacing space. URL safe symbols are; a-z A-Z 0-9 "." "-" "_" "~".

 - CharSymbol string, default "-" - Symbol for replacing unwanted chars.  URL safe symbols are; a-z A-Z 0-9 "." "-" "_" "~".

 - MaxLength int, default 32 - Trim text to this length, if set to 0 text will not be trimmed, default 32 chars

 - PreserveLength bool, default false - By default neighbouring chars to replace will be replaced with 1 CharSymbol or SpaceSymbol,
	//for example: "foo  b@@@r" => "foo_b-r". If true input text and output text will be the same length,
	//for example (with default SpaceSymbol and CharSymbol): "foo  bar" => "foo__bar" instead of "foo_bar",
	//"foo b@@r" => "foo_b--r" instead of "foo_b-r", default false

### With custom options
Example:
```go
opts := &sluggo.SlugOptions{
               			TrimSpaces:     true,
               			MaxLength:      15,
               			Case:           sluggo.Uppercase,
               			SpaceSymbol:    "-",
               			CharSymbol:     "~",
               			PreserveLength: true,
               		}

slug := sluggo.GetSlugWithOpts("This will, be new slug$ ", opts)
```

## License
The source files of this project are distributed under the Mozilla Public License, version 2.0, unless otherwise noted. Please read the FAQ if you have further questions regarding the license.