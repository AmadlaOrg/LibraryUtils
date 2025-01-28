package file

import "errors"

var (
	ErrorFailedToStatFile = errors.New("failed to stat file")
	ErrorNotAFile         = errors.New("not a file")
	ErrorIsDir            = errors.New("is a directory")
	ErrorMagicIsNotAMatch = errors.New("magic is not a match")
)
