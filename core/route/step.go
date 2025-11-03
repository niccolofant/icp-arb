package route

import (
	"fmt"

	"github.com/niccolofant/ic-arb/core/icp"
)

type Step struct {
	id        string
	fromToken icp.Token
	toToken   icp.Token
	dex       icp.Dex
}

func NewStep(
	fromToken icp.Token,
	toToken icp.Token,
	dex icp.Dex,
) (*Step, error) {
	if fromToken.Equal(toToken) {
		return nil, fmt.Errorf("fromToken cannot be the same as toToken")
	}

	dexNotAggregated, ok := dex.(icp.DexNotAggregated)
	if ok {
		if !dexNotAggregated.SupportToken(fromToken) {
			return nil, fmt.Errorf("fromToken not supported")
		}

		if !dexNotAggregated.SupportToken(toToken) {
			return nil, fmt.Errorf("toToken not supported")
		}
	}

	id := fmt.Sprintf(
		"[%s:%s|(%s:%s>%s:%s)]",
		dex.Type(),
		dex.CanisterID(),
		fromToken,
		fromToken.CanisterID(),
		toToken,
		toToken.CanisterID(),
	)

	return &Step{
		id:        id,
		fromToken: fromToken,
		toToken:   toToken,
		dex:       dex,
	}, nil
}

func (s Step) ID() string {
	return s.id
}

func (s Step) FromToken() icp.Token {
	return s.fromToken
}

func (s Step) ToToken() icp.Token {
	return s.toToken
}

func (s Step) Dex() icp.Dex {
	return s.dex
}

func (s Step) String() string {
	return fmt.Sprintf(
		"[%s|(%s>%s)]",
		s.dex.Type(),
		s.fromToken,
		s.toToken,
	)
}
