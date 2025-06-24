package probe

import (
	"math"

	"github.com/joshtechnologygroup/chardet/cda"
	"github.com/joshtechnologygroup/chardet/consts"
)

type EUCJPProbe struct {
	MultiByteCharSetProbe
	contextAnalyzer cda.Analyzer
}

func NewEUCJPProbe() *EUCJPProbe {
	ep := &EUCJPProbe{
		contextAnalyzer: cda.NewEUCJPContextAnalysis(),
	}
	ep.MultiByteCharSetProbe = NewMultiByteCharSetProbe(
		consts.EucJp,
		consts.Japanese,
		consts.UnknownLangFilter,
		cda.NewEUCJPDistributionAnalysis(),
		NewCodingStateMachine(EucJpSmModel()),
	)
	return ep
}

func (e *EUCJPProbe) Reset() {
	e.MultiByteCharSetProbe.Reset()
	e.contextAnalyzer.Reset()
}

func (e *EUCJPProbe) Feed(buf []byte) consts.ProbingState {
loop:
	for i, b := range buf {
		// PY3K: byte_str is a byte array, so byte is an int, not a byte
		codingState := e.codingSM.NextState(b)
		switch codingState {
		case consts.ErrorMachineState:
			e.state = consts.NotMeProbingState
			break loop
		case consts.ItsMeMachineState:
			e.state = consts.FoundItProbingState
			break loop
		case consts.StartMachineState:
			charLen := e.codingSM.CurrentCharLength()
			if i == 0 {
				e.lastChar[1] = b
				e.contextAnalyzer.Feed(e.lastChar[:], charLen)
				e.distributionAnalyzer.Feed(e.lastChar[:], charLen)
			} else {
				e.contextAnalyzer.Feed(buf[i-1:i+1], charLen)
				e.distributionAnalyzer.Feed(buf[i-1:i+1], charLen)
			}
		default:
		}
	}

	e.lastChar[0] = buf[len(buf)-1]
	if e.state == consts.DetectingProbingState &&
		e.contextAnalyzer.GotEnoughData() &&
		(e.GetConfidence() > e.ShortcutThreshold) {
		e.state = consts.FoundItProbingState
	}
	return e.state
}

func (e *EUCJPProbe) GetConfidence() float64 {
	return math.Max(e.contextAnalyzer.GetConfidence(), e.distributionAnalyzer.GetConfidence())
}
