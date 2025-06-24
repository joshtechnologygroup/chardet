package probe

import (
	"github.com/joshtechnologygroup/chardet/cda"
	"github.com/joshtechnologygroup/chardet/consts"
)

type EUCTWProbe struct {
	MultiByteCharSetProbe
}

func NewEUCTWProbe() *EUCTWProbe {
	return &EUCTWProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.EucTw,
			consts.Chinese,
			consts.UnknownLangFilter,
			cda.NewEUCTWDistributionAnalysis(),
			NewCodingStateMachine(EucTwSmModel()),
		),
	}
}
