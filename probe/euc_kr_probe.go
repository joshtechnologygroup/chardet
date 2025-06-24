package probe

import (
	"github.com/joshtechnologygroup/chardet/cda"
	"github.com/joshtechnologygroup/chardet/consts"
)

type EUCKRProbe struct {
	MultiByteCharSetProbe
}

func NewEUCKRProbe() *EUCKRProbe {
	return &EUCKRProbe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.EucKr,
			consts.Korean,
			consts.UnknownLangFilter,
			cda.NewEUCKRDistributionAnalysis(),
			NewCodingStateMachine(EucKrSmModel()),
		),
	}
}
