package probe

import (
	"github.com/joshtechnologygroup/chardet/consts"
)

type Probe interface {
	Feed([]byte) consts.ProbingState
	Reset()

	GetConfidence() float64

	CharSetName() string
	Language() string
	SetActive(bool)
	IsActive() bool
}
