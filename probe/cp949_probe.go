package probe

import (
	"github.com/joshtechnologygroup/chardet/cda"
	"github.com/joshtechnologygroup/chardet/consts"
)

type CP949Probe struct {
	MultiByteCharSetProbe
}

func NewCP949Probe() *CP949Probe {
	return &CP949Probe{
		MultiByteCharSetProbe: NewMultiByteCharSetProbe(
			consts.CP949,
			consts.Korean,
			consts.UnknownLangFilter,
			cda.NewEUCKRDistributionAnalysis(),
			NewCodingStateMachine(CP949SmModel()),
		),
	}
}
