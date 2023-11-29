package primitive

import "regexp"

var EmailPattern, _ = regexp.Compile(`.+@.+\..+`)
