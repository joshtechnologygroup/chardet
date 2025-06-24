package probe

import (
	"github.com/joshtechnologygroup/chardet/cda"
	"github.com/joshtechnologygroup/chardet/consts"
)

type JOHABProbe struct {
	MultiByteCharSetProbe
}

func NewJOHABProbe() *JOHABProbe {
	return &JOHABProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.Johab,
			consts.Korean,
			consts.UnknownLangFilter,
			cda.NewJOHABDistributionAnalysis(),
			NewCodingStateMachine(JohabSmModel()),
		),
	}
}
