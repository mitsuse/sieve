package classifier

import (
	"github.com/mitsuse/matrix-go/dense"
)

const (
	AlreadyInitializedError  = "AlreadyInitializedError"
	IncompatibleVersionError = "IncompatibleVersion"
)

type classifierJson struct {
	Version int           `json:"version"`
	Weights *dense.Matrix `json:"weights"`
}
